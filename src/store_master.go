package dash

import (
	"errors"
	"fmt"

	"github.com/coreos/etcd/clientv3"
	"golang.org/x/net/context"
)

var (
	errNotEnoughBudget = errors.New("store_master: not enough budget left")
)

const (
	storeService          = "/service"
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

type CreateServiceParams struct {
	ServiceId     string
	ImageUrl      string // NOTE idk if we will keep this field
	ImageVersion  string
	MinInstances  int64
	MaxInstances  int64
	Cpu           int64
	Ram           int64
	Datacenter    string
	CloudProvider string
}

func (s *StoreMaster) CreateService(ctx context.Context, arg CreateServiceParams) error {
	val, err := structToStr(arg)
	if err != nil {
		return err
	}
	if _, err := s.etcd.Put(ctx, storeService+arg.ServiceId, *val); err != nil {
		return err
	}
	return nil
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
