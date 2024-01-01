package rosda

import (
	"context"
	"errors"
	"fmt"

	"github.com/coreos/etcd/clientv3"
)

var (
	errNotEnoughBudget = errors.New("store_master: not enough budget left")
	errEmptyResponse   = errors.New("store_master: empty response received")
	errFailedDelete    = errors.New("store_master: failed delete")
)

const (
	storeService          = "/service"
	storeInstance         = "/instance"
	storeDatacenterBudget = "/datacenter/budget"
)

type StoreMaster struct {
	etcd *clientv3.Client
}

func NewStoreMaster(addr string) (*StoreMaster, error) {
	etcd, err := clientv3.New(clientv3.Config{Endpoints: []string{addr}})
	if err != nil {
		return nil, err
	}

	return &StoreMaster{
		etcd: etcd,
	}, nil
}

func (s *StoreMaster) CreateService(ctx context.Context, svc Service) error {
	return s.put(ctx, storeService+svc.ServiceId, svc)
}

func (s *StoreMaster) GetService(ctx context.Context, id string) (*Service, error) {
	resp, err := s.etcd.Get(ctx, storeService+id)
	if err != nil {
		return nil, err
	}
	if len(resp.Kvs) < 1 || resp.Kvs[0] == nil {
		return nil, errEmptyResponse
	}
	v := resp.Kvs[0].Value
	var svc Service
	if err := strToStruct(string(v), &svc); err != nil {
		return nil, err
	}
	return &svc, nil
}

func (s *StoreMaster) CreateInstance(ctx context.Context, inst ServiceInstance) error {
	return s.put(ctx, storeInstance+inst.InstanceId, inst)
}

func (s *StoreMaster) DeleteInstance(ctx context.Context, id string) error {
	resp, err := s.etcd.Delete(ctx, storeInstance+id)
	if err != nil {
		return err
	}
	if resp.Deleted != 1 {
		return errFailedDelete
	}
	return nil
}

func (s *StoreMaster) GetInstances(ctx context.Context, serviceId string) ([]*ServiceInstance, error) {
	vals, err := s.multiGet(ctx, storeInstance, []string{serviceId})
	if err != nil {
		return nil, err
	}

	insts := make([]*ServiceInstance, 0, len(vals))

	// TODO test how slow this decoding is
	for _, v := range vals {
		var inst ServiceInstance
		if err := strToStruct(v, &inst); err != nil {
			return nil, err
		}
		insts = append(insts, &inst)
	}
	return insts, nil
}

type UpdateBudgetParams struct {
	Datacenters []string
	Cpu         int64
	Ram         int64
}

// Update the budget for a specific region. NOTE: DO NOT call without global lock on service!
func (s *StoreMaster) UpdateBudget(ctx context.Context, arg UpdateBudgetParams) error {
	vals, err := s.multiGet(ctx, storeDatacenterBudget, arg.Datacenters)
	if err != nil {
		return err
	}

	newVals := make([]string, 0, len(arg.Datacenters))
	for _, v := range vals {
		var vOld UpdateBudget
		if err := strToStruct(string(v), &vOld); err != nil {
			return err
		}

		// Ensure enough budget left
		if vOld.Cpu < arg.Cpu && vOld.Ram < arg.Ram {
			return errors.Join(errNotEnoughBudget, fmt.Errorf("have %s, need %s"))
		}

		// Update budget
		vOld.Cpu -= arg.Cpu
		vOld.Ram -= arg.Ram
		vNewStr, err := structToStr(vOld)
		if err != nil {
			return err
		}
		newVals = append(newVals, *vNewStr)
	}

	return s.multiPutAtomic(ctx, storeDatacenterBudget, arg.Datacenters, newVals, vals)
}

func (s *StoreMaster) put(ctx context.Context, key string, v any) error {
	val, err := structToStr(v)
	if err != nil {
		return err
	}
	if _, err := s.etcd.Put(ctx, key, *val); err != nil {
		return err
	}
	return nil
}

func (s *StoreMaster) multiGet(ctx context.Context, prefix string, keys []string) ([]string, error) {
	txn := s.etcd.Txn(ctx)
	// Prepare update operations
	for _, k := range keys {
		txn = txn.
			// ensure key exists before hand
			If(clientv3.Compare(clientv3.Version(prefix+k), ">", 0)).
			Then(clientv3.OpGet(prefix + k))
	}
	resp, err := txn.Commit()
	if err != nil {
		return nil, err
	}

	if !resp.Succeeded {
		return nil, errors.New("unable to get keys")
	}

	vals := make([]string, 0, len(keys))
	for _, resp := range resp.Responses {
		for _, kv := range resp.GetResponseRange().Kvs {
			vals = append(vals, string(kv.Value))
		}
	}
	return vals, nil
}

func (s *StoreMaster) multiPutAtomic(ctx context.Context, prefix string, keys []string, newVals []string, oldVals []string) error {
	if len(keys) != len(newVals) {
		return errors.New("keys len does not match vals len")
	}
	if len(newVals) != len(oldVals) {
		return errors.New("newVals len does not match oldVals len")
	}

	txn := s.etcd.Txn(ctx)
	for i := 0; i < len(keys); i++ {
		k := prefix + keys[i]
		txn = txn.
			// ensure value hasnt change since previous get
			If(clientv3.Compare(clientv3.Value(k), "=", oldVals[i])).
			// ensure key exists before hand
			If(clientv3.Compare(clientv3.Version(k), ">", 0)).
			Then(clientv3.OpPut(k, newVals[i]))
	}
	resp, err := txn.Commit()
	if err != nil {
		return err
	}

	if !resp.Succeeded {
		return errors.New("unable to store multiple kv pairs")
	}
	return nil
}
