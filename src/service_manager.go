package dash

import (
	"context"
	"fmt"
	"sync"

	"google.golang.org/protobuf/proto"
)

type serviceManagerMQ interface {
	Join(topic string) (serviceManagerMQTopic, error)
}

type serviceManagerMQTopic interface {
	Subscribe(handler ServiceManagerHandler)
}

type serviceManagerStore interface {
}

type ServiceManager struct {
	ctx      context.Context
	mq       serviceManagerMQ
	store    serviceManagerStore
	nodeLock sync.RWMutex
	//tracer serviceManagerTracer
}

type ServiceManagerHandler func(b []byte) error

func NewServiceManager() (*ServiceManager, error) {
	s := &ServiceManager{}

	// register handlers
	hs := map[string]ServiceManagerHandler{
		"service/create":            s.createHandler,
		"service/update":            s.updateContainerHandler,
		"service/event/load":        s.loadHandler,
		"config/update":             s.updateConfigHandler,
		"manage/service/create":     s.manageServiceCreationOnNode,
		"manage/service/update":     s.manageServiceUpdateOnNode,
		"manage/service/scale-up":   s.manageServiceScaleUpOnNode,
		"manage/service/scale-down": s.manageServiceScaleDownOnNode,
	}
	for k := range hs {
		if err := s.register(k, hs[k]); err != nil {
			return nil, fmt.Errorf("unable to register handler to topic %s: %v", k, err.Error())
		}
	}

	return s, nil
}

func (s *ServiceManager) register(topic string, handler ServiceManagerHandler) error {
	t, err := s.mq.Join(topic)
	if err != nil {
		return err
	}
	t.Subscribe(handler)
	return nil
}

// Used for creating new services
func (s *ServiceManager) createHandler(b []byte) error {
	//msg :=
	if err := proto.Unmarshal(b, msg); err != nil {
		return err
	}

	// TODO: create service in db
	// TODO: find available workers and directly reserve capacity (in one query)
	// TODO: publish to manage creation
	return nil
}

func (s *ServiceManager) updateContainerHandler(b []byte) error {
	//msg :=
	if err := proto.Unmarshal(b, msg); err != nil {
		return err
	}

	// TODO: store update in db
	// TODO: find relevant workers

	return nil
}

func (s *ServiceManager) updateConfigHandler(b []byte) error {
	//msg :=
	if err := proto.Unmarshal(b, msg); err != nil {
		return err
	}

	// TODO: store update in db
	// TODO: check if region update, node count
	// TODO: fetch nodes that need updating
	// TODO: inform peers they need to update
	// TODO: monitor that services were updated
	return nil
}

// Handles instance load (cpu, ram)
func (s *ServiceManager) loadHandler(b []byte) error {
	//msg :=
	if err := proto.Unmarshal(b, msg); err != nil {
		return err
	}
	arg := CreateServiceLoadParams{
		ServiceInstanceID: msg.ServiceId,
		Cpu:               msg.Cpu,
		Ram:               msg.Ram,
		ClientCreatedAt:   msg.CreatedAt,
	}
	if err := s.p.CreateServiceLoad(s.ctx, arg); err != nil {
		return err
	}

	return nil

	// These are tasks of autoscaller
	// TODO: if scale down then dereserve capacity from nodes and inform them (first gateway needs to be updated then node informed then local query executed)
	// TODO: search for available nodes and directly reserve
}

func (s *ServiceManager) manageServiceCreationOnNode(b []byte) error {
	// TODO: inform selected workers
	// TODO: if worker(s) down unreserve capacity and try again
	// TODO: if timeout reached send cancel request to worker(s) and reassign
	// TODO: stores service instance in db
	return nil
}

// Manages updating a service running on a given worker node
func (s *ServiceManager) manageServiceUpdateOnNode(b []byte) error {
	//msg :=
	if err := proto.Unmarshal(b, msg); err != nil {
		return err
	}
	// TODO: generate credentials for worker to access update
	// TODO: share credentials with workers and allow them to pull
	// TODO: fetch status from workers
	// TODO: if fail retry an update w/ attempt count +1

	return nil
}

// Ensures worker stops service and capacity reclaimed
func (s *ServiceManager) manageServiceScaleDownOnNode(b []byte) error {
	// TODO: inform load balancer to remove worker node for service
	// TODO: inform selected worker
	// TODO: reclaim capacity for node in db

	return nil
}

func (s *ServiceManager) manageServiceScaleUpOnNode(b []byte) error {
	// TODO: register undo action; dereserve capacity
	// TODO: inform selected worker
	// TODO: let them download image
	// TODO: ensure the images are running
	// TODO: register load balancer undo action
	// TODO: ping undoer so action does happen

	return nil
}

// Calculates by how much a service should be scaled
func (s *ServiceManager) calculateScaleFactor(serviceId string) (int32, error) {
	panic("implement me")
}
