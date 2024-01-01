package rosda

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/chk-n/gomemq"
	"github.com/chk-n/rosda/pkg/scheduler"
	"github.com/google/uuid"
)

const (
	maxSubscriberAckTime  = 10 * time.Second
	maxSubscriberDoneTime = 5 * time.Minute
)

// Interfaces
type serviceManagerMQTopic[T any] interface {
	Subscribe(handler ServiceManagerHandler[T])
	Publish(msg []byte)
	PublishBatchDone(msg [][]byte) *gomemq.Context
}

type serviceManagerStore interface {
	CreateService(ctx context.Context, arg CreateServiceParams) error
	UpdateBudget(ctx context.Context, arg UpdateBudgetParams) error
	GetService(ctx context.Context, id string) (*Service, error)
	GetInstances(ctx context.Context, serviceId string) ([]*ServiceInstance, error)
}

type serviceManagerScheduler interface {
	Schedule(container scheduler.Containers, nodes scheduler.Nodes) ([]scheduler.ContainerPlacement, int)
}

type serviceDistributedRWMutex interface {
	Lock(ctx context.Context, id string) (*RWMutex, error)
}

type ServiceManager struct {
	ctx       context.Context
	store     serviceManagerStore
	nodeLock  sync.RWMutex
	scheduler serviceManagerScheduler
	dmu       serviceDistributedRWMutex
	//tracer serviceManagerTracer
}

const (
	serviceCreateTopic            = "service/create"
	serviceDeleteTopic            = "service/delete"
	manageServiceCreateCloudTopic = "manage/service/create/cloud"
	manageServiceDeleteCloudTopic = "manage/service/delete/cloud"
)

type ServiceManagerHandler[T any] func(b T) error

// nolint:errcheck
func NewServiceManager() *ServiceManager {
	s := &ServiceManager{}

	// define general topic structure
	cfg := gomemq.ConfigTopic{}

	// Register subscribers
	{
		cfg.Name = serviceCreateTopic
		t, _ := gomemq.NewTopic[serviceCreate](cfg)
		t.Subscribe(s.createHandler)
	}
	{
		cfg.Name = serviceDeleteTopic
		t, _ := gomemq.NewTopic[serviceDelete](cfg)
		t.Subscribe(s.deleteHandler)
	}
	{
		cfg.Name = manageServiceCreateCloudTopic
		t, _ := gomemq.NewTopic[manageServiceCreateCloud](cfg)
		t.Subscribe(s.manageServiceCreateCloud)
	}

	{
		cfg.Name = manageServiceDeleteCloudTopic
		t, _ := gomemq.NewTopic[manageServiceDeleteCloud](cfg)
		t.Subscribe(s.manageServiceDeleteCloud)
	}

	return s
}

type serviceCreate struct {
	HostUrl        string
	ServiceUrl     string
	ServiceVersion string
	Config         serviceConfig
	// TODO service type (web, job, worker)
}

// Used for creating/adding new services in a region
func (s *ServiceManager) createHandler(msg serviceCreate) error {
	// TODO: start asychronous image pull

	if err := s.store.CreateService(s.ctx, CreateServiceParams{
		// TODO:
	}); err != nil {
		return err
	}

	// update datacenter budget
	if err := s.updateDatacenterBudget(msg); err != nil {
		return err
	}

	// TODO prepare cloud provider environment (if required: create lb, set up network, push container to cloud provider)

	if err := s.notifyServiceCreateCloud(msg); err != nil {
		// TODO: clean up inconsistencies (maybe publish to clean up topic)
		return fmt.Errorf("notifying service creation failed: %w", err)
	}

	return nil

}

type serviceDelete struct {
	ServiceId string
}

// Handles deleting service(s) in a region
func (s *ServiceManager) deleteHandler(msg serviceDelete) error {

	// aquire distributed lock
	mu, err := s.dmu.Lock(s.ctx, msg.ServiceId)
	if err != nil {
		return err
	}
	defer mu.Unlock(s.ctx)

	insts, err := s.store.GetInstances(s.ctx, msg.ServiceId)
	if err != nil {
		// TODO: clean up inconsistencies (maybe publish to clean up topic)
		return err
	}

	return s.notifyServiceDeleteCloud(insts)
}

type manageServiceCreateCloud struct {
	ServiceId  string
	InstanceId string
	Service    serviceCreate
}

func (s *ServiceManager) manageServiceCreateCloud(msg manageServiceCreateCloud) error {

	// TODO spin up vm instance

	// TODO configure firewall for instance (inbound, outbound)

	// TODO add instance to lb (if external)

	// TODO monitor everything done
}

type manageServiceDeleteCloud struct {
	InstanceId string
}

func (s *ServiceManager) manageServiceDeleteCloud(msg manageServiceDeleteCloud) error {

	// TODO update load balancer routing

	// TODO inform them to delete vm instance

	// TODO monitor everything done
}

// // Performs rolling update for a service image
// func (s *ServiceManager) updateHandler(msg UpdateServiceImage) error {
// 	// aquire distributed lock
// 	dmu, err := s.dmu.Lock(s.ctx, msg.ServiceId)
// 	if err != nil {
// 		return err
// 	}
// 	defer dmu.Unlock(s.ctx)

// 	// TODO determine what needs to be updated
// 	// - service image changes
// 	// - region chaned
// 	// - service count changes
// 	// - resource changes

// 	// TODO: store update in db
// 	// TODO: find relevant workers
// 	// TODO: publish to monitr update
// 	// TODO: inform workers (here update type comes into play)
// 	// TODO: wait for done from all update monitor
// 	// TODO: confirm with worker to show uodate

// 	return nil
// }

// func (s *ServiceManager) updateConfigHandler(b []byte) error {

// 	// TODO: AQUIRE DISTRIBUTED LOCK

// 	// TODO: get previous config
// 	// TODO: store update in db
// 	// TODO: check if region added/removed or service count change
// 	// TODO: if remove or service count down => publish to scale-odwn
// 	// TODO: if region added or service count up => scale-up

// 	return nil
// }

// Handles instance load (cpu, ram)
// func (s *ServiceManager) loadHandler(msg ServiceLoad) error {
// 	arg := CreateServiceLoadParams{
// 		ServiceInstanceID: msg.ServiceId,
// 		Cpu:               msg.Cpu,
// 		Ram:               msg.Ram,
// 		ClientCreatedAt:   msg.ClientCreatedAt,
// 	}
// 	if err := s.store.CreateServiceLoad(s.ctx, arg); err != nil {
// 		return err
// 	}

// 	return nil

// 	// These are tasks of autoscaller
// 	// TODO: if scale down then dereserve capacity from nodes and inform them (first gateway needs to be updated then node informed then local query executed)
// 	// TODO: search for available nodes and directly reserve
// }

// // Manages updating a service running on a given worker node
// func (s *ServiceManager) manageServiceUpdateOnNode(b []byte) error {
// 	// TODO: generate credentials for worker to access update
// 	// TODO: share credentials with workers and allow them to pull
// 	// TODO: fetch status from workers
// 	// TODO: if fail retry an update w/ attempt count +1

// 	return nil
// }

// func (s *ServiceManager) manageConfigUpdateOnNode(b []byte) error {
// 	// TODO: inform worker they need to update
// 	// TODO: monitor that services were updated

// 	return nil
// }

// Ensures worker stops service and capacity reclaimed
// func (s *ServiceManager) manageServiceScaleDownOnNode(b []byte) error {
// 	// TODO: inform load balancer to remove worker node for service
// 	// TODO: inform selected worker
// 	// TODO: reclaim capacity for node in db

// 	return nil
// }

// func (s *ServiceManager) manageServiceScaleUpOnNode(b []byte) error {
// 	// TODO: register undo action; dereserve capacity
// 	// TODO: inform selected worker
// 	// TODO: let them download image
// 	// TODO: ensure the images are running
// 	// TODO: register load balancer undo action
// 	// TODO: ping undoer so action does happen

// 	return nil
// }

// ---------------- //
// Helper functions //
// ---------------- //

func (s *ServiceManager) updateDatacenterBudget(msg serviceCreate) error {
	return s.store.UpdateBudget(s.ctx, UpdateBudgetParams{
		Datacenters: msg.Config.Regions,
		Cpu:         msg.Config.CpuPerInstance * msg.Config.MinInstances,
		Ram:         msg.Config.RamPerInstance * msg.Config.MinInstances,
	})
}

func (s *ServiceManager) notifyServiceCreateCloud(msg serviceCreate) error {
	// Generate new service id
	sId, err := genUid()
	if err != nil {
		return err
	}

	outs := []manageServiceCreateCloud{}
	for i := 0; i < int(msg.Config.MinInstances); i++ {
		// instance id
		iId, err := genUid()
		if err != nil {
			return err
		}
		outs = append(outs, manageServiceCreateCloud{
			ServiceId:  sId,
			InstanceId: iId,
			Service:    msg,
		})
	}

	ctx, _ := gomemq.PublishBatchDone[manageServiceCreateCloud](manageServiceCreateCloudTopic, outs)

	// wait for done or return timeout error
	select {
	case <-ctx.Done():
		return nil
	case <-ctx.WithAckTimeout(maxSubscriberAckTime):
		// cancel any future work, blocks until current work finishes or errors
		ctx.Cancel()
	case <-ctx.WithDoneTimeout(maxSubscriberDoneTime):
		ctx.Cancel()
		return errors.New("") // TODO
	}

	return nil
}

func (s *ServiceManager) notifyServiceDeleteCloud(insts []*ServiceInstance) error {
	outs := []manageServiceDeleteCloud{}
	for _, i := range insts {
		outs = append(outs, manageServiceDeleteCloud{InstanceId: i.InstanceId})
	}

	ctx, _ := gomemq.PublishBatchDone[manageServiceDeleteCloud](manageServiceDeleteCloudTopic, outs)

	// wait for done or return timeout error
	select {
	case <-ctx.Done():
		return nil
	case <-ctx.WithAckTimeout(maxSubscriberAckTime):
		// cancel any future work, blocks until current work finishes or errors
		ctx.Cancel()
	case <-ctx.WithDoneTimeout(maxSubscriberDoneTime):
		ctx.Cancel()
		return errors.New("") // TODO
	}

	return nil
}

func genUid() (string, error) {
	uuid, err := uuid.NewV7()
	if err != nil {
		return "", err
	}
	return uuid.String(), nil
}

// UNUSED onprem code (might be required later when worker imlemented)

// func (s *ServiceManager) manageServiceCreateOnprem(b []byte) error {
// 	// TODO: inform worker
// 	// TODO: if worker down return err
// 	// TODO: ping worker once done
// 	// TODO: stores service instance in db
// 	return nil
// }

// func (s *ServiceManager) updateNodeCapacity(placements []scheduler.ContainerPlacement) error {
// 	// reserve capacity here to reduce delay between scheduler deciding on worker and capacity being reserved
// 	tx, err := s.db.Begin()
// 	if err != nil {
// 		return err
// 	}
// 	store := s.store.WithTx(tx)
// 	for _, p := range placements {
// 		cpu, ram := scheduler.CountResources(p.Containers)
// 		if err := store.UpdateNodeAvailableCapacity(s.ctx, UpdateNodeAvailableCapactiyParams{
// 			NodeId:       p.NodeId,
// 			AvailableCpu: cpu,
// 			AvailableRam: ram,
// 		}); err != nil {
// 			if err := tx.Rollback(); err != nil {
// 				return fmt.Errorf("unable to rollback %w", err)
// 			}
// 			return err
// 		}
// 	}

// 	return tx.Commit()
// }
// func (s *ServiceManager) notifyServiceCreatenOnprem(placements []scheduler.ContainerPlacement, msg CreateService) error {
// 	outs := []ManageCreateService{}
// 	for _, p := range placements {
// 		uid, _ := uuid.NewV7()
// 		outs = append(outs, ManageCreateService{
// 			WorkerId:          p.NodeId,
// 			ServiceInstanceId: uid.String(),
// 			Service:           msg,
// 		})
// 	}

// 	ctx, _ := gomemq.PublishBatchDone[ManageCreateService](manageServiceCreateTopic, outs)

// 	// wait for done or return timeout error
// 	select {
// 	case <-ctx.Done():
// 		return nil
// 	case <-ctx.WithAckTimeout(maxSubscriberAckTime):
// 		// cancel any future work, blocks until current work finishes or errors
// 		ctx.Cancel()
// 	case <-ctx.WithDoneTimeout(maxSubscriberDoneTime):
// 		ctx.Cancel()
// 		// TODO: clean up inconsistencies (maybe publish to clean up topic)
// 		return errors.New("") // TODO
// 	}

// 	return nil
// }

// func (s *ServiceManager) scheduleService(msg createService) ([]scheduler.ContainerPlacement, error) {
// 	avlbWorkers, err := s.store.GetAvailableNodes(s.ctx, GetAvailableNodesParams{
// 		// TODO:
// 	})
// 	if err != nil {
// 		return nil, err
// 	}

// 	// TODO: query scheduler
// 	containers := internal.ServiceToSchedulerContainer(msg)
// 	nodes := internal.NodesToSchedulerNodes(avlbWorkers)
// 	placements, count := s.scheduler.Schedule(containers, nodes)
// 	if int(msg.Config.MinInstances) != count {
// 		return nil, fmt.Errorf("unable to create service: requested service count could not be scheduled")
// 	}

// 	return placements, nil
// }
