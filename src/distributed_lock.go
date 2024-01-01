package rosda

import (
	"context"

	"github.com/coreos/etcd/clientv3"
	"go.etcd.io/etcd/clientv3/concurrency"
)

type DistributedLock struct {
	cli *clientv3.Client
}

func NewDistributedLock(etcdClient *clientv3.Client) *DistributedLock {
	return &DistributedLock{
		cli: etcdClient,
	}
}

type RWMutex struct {
	sesh *concurrency.Session
	l    *concurrency.Mutex
}

func (d *DistributedLock) Lock(ctx context.Context, id string) (*RWMutex, error) {
	// create a sessions to aqcuire a lock
	s, err := concurrency.NewSession(d.cli)
	if err != nil {
		return nil, err
	}
	l := concurrency.NewMutex(s, id)
	// acquire lock (or wait to have it)
	if err := l.Lock(ctx); err != nil {
		return nil, err
	}
	return &RWMutex{
		sesh: s,
	}, nil
}

func (mu *RWMutex) Unlock(ctx context.Context) error {
	defer mu.sesh.Close()
	return mu.l.Unlock(ctx)
}
