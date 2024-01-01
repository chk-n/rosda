package rosda

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/chk-n/gomemq"
	"github.com/chk-n/rosda/pkg/scheduler"
	"github.com/chk-n/rosda/src/internal"
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
	serviceCreateTopic       = "service/create"
	serviceDeleteTopic       = "service/delete"
	manageServiceCreateTopic = "manage/service/create"
	manageServiceDeleteTopic = "manage/service/delete"
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
		t, _ := gomemq.NewTopic[ServiceCreate](cfg)
		t.Subscribe(s.createHandler)
	}
	{
		cfg.Name = serviceDeleteTopic
		t, _ := gomemq.NewTopic[ServiceDelete](cfg)
		t.Subscribe(s.deleteHandler)
	}
	{
		cfg.Name = manageServiceCreateTopic
		t, _ := gomemq.NewTopic[ManageServiceCreate](cfg)
		t.Subscribe(s.manageServiceCreateCloud)
	}

	{
		cfg.Name = manageServiceDeleteTopic
		t, _ := gomemq.NewTopic[ManageServiceDelete](cfg)
		t.Subscribe(s.manageServiceDeleteCloud)
	}

	return s
}

// Used for creating/adding new services in a region
func (s *ServiceManager) createHandler(msg CreateService) error {
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

	// TODO prepare cloud provider environment

	resp, err := s.notifyServiceCreateCloud(msg)
	if err != nil {
		return fmt.Errorf("notifying service creation failed: %w", err)
	}

	// TODO inform load balancer of change

	return nil

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

// Handles deleting service(s) in a region
func (s *ServiceManager) deleteHandler(imageUrl string) error {

	// TODO AQUIRE DISTRIBUTED LOCK

	// TODO find nodes where services to be deleted

	// TODO update load balancer routing

	// TODO inform them to delete service

	// TODO publish to monitor delete

	return nil
}

func (s *ServiceManager) manageServiceCreateCloud(msg ServiceCreateCloud) error {
	// TODO: send request to cloud provider
	//
}

func (s *ServiceManager) manageServiceDeleteCloud() error {

}

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

func (s *ServiceManager) updateDatacenterBudget(msg CreateService) error {
	return s.store.UpdateBudget(s.ctx, UpdateBudgetParams{
		Datacenters: msg.Config.Regions,
		Cpu:         msg.Config.CpuPerInstance * msg.Config.MinInstances,
		Ram:         msg.Config.RamPerInstance * msg.Config.MinInstances,
	})
}

func (s *ServiceManager) notifyServiceCreateCloud(msg CreateService) error {

	// Generate new service id
	sId, err := genUid()
	if err != nil {
		return err
	}

	outs := []ManageServiceCreateCloud{}
	for i := 0; i < int(msg.Config.MinInstances); i++ {
		// instance id
		iId, err := genUid()
		if err != nil {
			return err
		}
		outs = append(outs, ManageServiceCreateCloud{})
	}

	ctx, _ := gomemq.PublishBatchDone[ManageServiceCreateCloud](manageServiceCreateCloudTopic, outs)

	// wait for done or return timeout error
	select {
	case <-ctx.Done():
		return nil
	case <-ctx.WithAckTimeout(maxSubscriberAckTime):
		// cancel any future work, blocks until current work finishes or errors
		ctx.Cancel()
	case <-ctx.WithDoneTimeout(maxSubscriberDoneTime):
		ctx.Cancel()
		// TODO: clean up inconsistencies (maybe publish to clean up topic)
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
