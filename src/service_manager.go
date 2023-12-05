package dash

import (
	"context"
	"sync"

	"google.golang.org/protobuf/proto"
)

type ServiceManager struct {
	ctx      context.Context
	mq       messageQueue
	cm       containerManager
	p        persistenceMaster
	nodeLock sync.RWMutex
	//tracer serviceManagerTracer
}

type ServiceManagerHandler func(b []byte) error

func NewServiceManager() *ServiceManager {

	// define handlers
	s := &ServiceManager{}

	// register handlers
	s.register("", s.createHandler)
	s.register("", s.updateHandler)
	s.register("", s.loadHandler)
	s.register("", s.gatewayHandler)
	return s
}

func (s *ServiceManager) register(topic string, handler ServiceManagerHandler) {
	s.mq.Subscribe(s.ctx, topic, handler)
}

// Used for creating new services
func (s *ServiceManager) createHandler(b []byte) error {
	//msg :=
	if err := proto.Unmarshal(b, msg); err != nil {
		return err
	}
	if err := s.cm.Pull(msg.ImageRef); err != nil {
		return err
	}
	// TODO: find available workers and directly reserve capacity (in one query)
	// TODO: inform selected workers
	// TODO: monitor if service is up (workers first need to download)
	// TODO: if timeout reached send cancel request to worker(s) and reassign
	// TODO: stores instances in db
	// TODO: publish to updte lb or gateway topic
	return nil
}

func (s *ServiceManager) updateHandler(b []byte) error {

	//msg :=
	if err := proto.Unmarshal(b, msg); err != nil {
		return err
	}
	if err := s.cm.Pull(msg.ImageRef); err != nil {
		return err
	}
	// TODO: fetch nodes that need updating
	// TODO: inform peers they need to update
	// TODO: update image_path of service
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
		ServiceInstanceID: msg.InstanceId,
		Cpu:               msg.Cpu,
		Ram:               msg.Ram,
		ClientCreatedAt:   msg.CreatedAt,
	}
	if err := s.p.CreateServiceLoad(s.ctx, arg); err != nil {
		return err
	}

	scaleFactor, err := s.calculateScaleFactor(msg.ServiceId)
	if err != nil {
		return err
	}

	// TODO: if scale down then dereserve capacity from nodes and inform them (first gateway needs to be updated then node informed then local query executed)
	// TODO: search for available nodes and directly reserve
	// TODO: inform selected workers
	// TODO: let them download image
	// TODO: ensure the images are running
	// TODO: if anything fails dereserve capacity
	// TODO: store scale event
	// TODO: publish to update lb or gateway topic
	// TODO: get config
	// TODO: based on current load and config make decision
	return nil
}

func (s *ServiceManager) gatewayHandler(b []byte) error {
	// TODO: read message using protbuf
	// TODO: get gateways available based on region
	// TODO: send create or update requests to those gateways
	// TODO: store all updates
	// TODO: ensure it all worked
	return nil
}

// Calculates by how much a service should be scaled
func (s *ServiceManager) calculateScaleFactor(serviceId string) (int32, error) {
	panic("implement me")
}
