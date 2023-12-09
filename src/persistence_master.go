package dash

import (
	"context"
	"time"

	qm "github.com/chk-n/dash/src/sql/query_master"
)

type dbMaster interface {
	CreateClientAuthToken(ctx context.Context, arg qm.CreateClientAuthTokenParams) error
	CreateNode(ctx context.Context, arg qm.CreateNodeParams) error
	CreateService(ctx context.Context, arg qm.CreateServiceParams) error
	CreateServiceAccessRule(ctx context.Context, arg qm.CreateServiceAccessRuleParams) error
	CreateServiceInstance(ctx context.Context, arg qm.CreateServiceInstanceParams) error
	CreateServiceLoad(ctx context.Context, arg qm.CreateServiceLoadParams) error
	GetAvailableNodes(ctx context.Context, arg qm.GetAvailableNodesParams) ([]qm.Node, error)
	GetNodePublicKey(ctx context.Context, nodeID string) ([]string, error)
	GetNodeIdsByService(ctx context.Context, serviceID string) ([]string, error)
	GetServiceLoadSince(ctx context.Context, arg qm.GetServiceLoadSinceParams) ([]qm.ServiceLoad, error)
	UpdateServiceConfig(ctx context.Context, arg qm.UpdateServiceConfigParams) error
	UpdateServiceWasm(ctx context.Context, arg qm.UpdateServiceWasmParams) error
	UpdateNodeAvailableCapacity(ctx context.Context, arg qm.UpdateNodeAvailableCapactiyParams) error
}

type PersistenceMaster struct {
	db dbMaster
}

func NewPersistenceMaster(db dbMaster) *PersistenceMaster {
	return &PersistenceMaster{
		db: db,
	}
}

type CreateClientAuthTokenParams struct {
	Token     string
	IpAddress string
	Expiry    time.Time
}

func (p *PersistenceMaster) CreateClientAuthToken(ctx context.Context, arg CreateClientAuthTokenParams) error {
	return p.db.CreateClientAuthToken(ctx, qm.CreateClientAuthTokenParams{
		Token:     arg.Token,
		IpAddress: arg.IpAddress,
		Expiry:    arg.Expiry,
	})
}

type CreateNodeParams struct {
	NodeID       string
	PublicKey    string
	Region       string
	MaxCpu       int64
	MaxRam       int64
	AvailableCpu int64
	AvailableRam int64
	IpAddress    string
	Port         string
}

func (p *PersistenceMaster) CreateNode(ctx context.Context, arg CreateNodeParams) error {
	return p.db.CreateNode(ctx, qm.CreateNodeParams{
		NodeID:       arg.NodeID,
		PublicKey:    arg.PublicKey,
		Region:       arg.Region,
		MaxCpu:       arg.MaxCpu,
		MaxRam:       arg.MaxRam,
		AvailableCpu: arg.AvailableCpu,
		AvailableRam: arg.AvailableRam,
		IpAddress:    arg.IpAddress,
		Port:         arg.Port,
	})
}

type CreateServiceParams struct {
	ServiceID      string
	MinInstances   int64
	MaxInstances   int64
	CpuPerInstance int64
	RamPerInstance int64
	Wasm           []byte
}

func (p *PersistenceMaster) CreateService(ctx context.Context, arg CreateServiceParams) error {
	return p.db.CreateService(ctx, qm.CreateServiceParams{
		ServiceID:      arg.ServiceID,
		MinInstances:   arg.MinInstances,
		MaxInstances:   arg.MaxInstances,
		CpuPerInstance: arg.CpuPerInstance,
		RamPerInstance: arg.RamPerInstance,
		Wasm:           arg.Wasm,
	})
}

type CreateServiceAccessRuleParams struct {
	ServiceIDSource      string
	ServiceIDDestination string
}

func (p *PersistenceMaster) CreateServiceAccessRule(ctx context.Context, arg CreateServiceAccessRuleParams) error {
	return p.db.CreateServiceAccessRule(ctx, qm.CreateServiceAccessRuleParams{
		ServiceIDSource:      arg.ServiceIDSource,
		ServiceIDDestination: arg.ServiceIDDestination,
	})
}

type CreateServiceInstanceParams struct {
	ServiceID         string
	ServiceInstanceID string
	NodeID            string
}

func (p *PersistenceMaster) CreateServiceInstance(ctx context.Context, arg CreateServiceInstanceParams) error {
	return p.db.CreateServiceInstance(ctx, qm.CreateServiceInstanceParams{
		ServiceID:         arg.ServiceID,
		ServiceInstanceID: arg.ServiceInstanceID,
		NodeID:            arg.NodeID,
	})
}

type CreateServiceLoadParams struct {
	ServiceInstanceID string
	Cpu               int64
	Ram               int64
	ClientCreatedAt   time.Time
}

func (p *PersistenceMaster) CreateServiceLoad(ctx context.Context, arg CreateServiceLoadParams) error {
	return p.db.CreateServiceLoad(ctx, qm.CreateServiceLoadParams{
		ServiceInstanceID: arg.ServiceInstanceID,
		Cpu:               arg.Cpu,
		Ram:               arg.Ram,
		ClientCreatedAt:   arg.ClientCreatedAt,
	})
}

type GetAvailableNodesParams struct {
	Region       string
	AvailableCpu int64
	AvailableRam int64
}

func (p *PersistenceMaster) GetAvailableNodes(ctx context.Context, arg GetAvailableNodesParams) ([]Node, error) {
	nodes, err := p.db.GetAvailableNodes(ctx, qm.GetAvailableNodesParams{
		Region:       arg.Region,
		AvailableCpu: arg.AvailableCpu,
		AvailableRam: arg.AvailableRam,
	})
	if err != nil {
		return nil, err
	}
	// Convert internal representation to external representation (assuming Node struct exists)
	var externalNodes []Node
	for _, node := range nodes {
		externalNodes = append(externalNodes, Node{
			NodeId:       node.NodeID,
			PublicKey:    node.PublicKey,
			Region:       node.Region,
			MaxCpu:       node.MaxCpu,
			MaxRam:       node.MaxRam,
			AvailableCpu: node.AvailableCpu,
			AvailableRam: node.AvailableRam,
			IpAddress:    node.IpAddress,
			Port:         node.Port,
		})
	}
	return externalNodes, nil
}

func (p *PersistenceMaster) GetNodePublicKey(ctx context.Context, nodeID string) ([]string, error) {
	return p.db.GetNodePublicKey(ctx, nodeID)
}

type GetNodesByServiceRow struct {
	NodeID  string
	NodeUrl string
}

func (p *PersistenceMaster) GetNodeIdsByService(ctx context.Context, serviceID string) ([]string, error) {
	ids, err := p.db.GetNodeIdsByService(ctx, serviceID)
	if err != nil {
		return nil, err
	}
	// Convert internal representation to external representation (assuming Node struct exists)
	var nodeIds []string
	for _, id := range ids {
		nodeIds = append(nodeIds, id)
	}
	return nodeIds, nil
}

type GetServiceLoadSinceParams struct {
	ServiceID string
	CreatedAt time.Time
}

func (p *PersistenceMaster) GetServiceLoadSince(ctx context.Context, arg GetServiceLoadSinceParams) ([]ServiceLoad, error) {
	loads, err := p.db.GetServiceLoadSince(ctx, qm.GetServiceLoadSinceParams{
		ServiceID: arg.ServiceID,
	})
	if err != nil {
		return nil, err
	}
	// Convert internal representation to external representation (assuming ServiceLoad struct exists)
	var externalLoads []ServiceLoad
	for _, load := range loads {
		externalLoads = append(externalLoads, ServiceLoad{
			ServiceId:         load.ServiceID,
			ServiceInstanceId: load.ServiceInstanceID,
			Cpu:               load.Cpu,
			Ram:               load.Ram,
			ClientCreatedAt:   load.ClientCreatedAt,
		})
	}
	return externalLoads, nil
}

type UpdateServiceConfigParams struct {
	ServiceId      string
	MinInstances   int64
	MaxInstances   int64
	CpuPerInstance int64
	RamPerInstance int64
}

func (p *PersistenceMaster) UpdateServiceConfig(ctx context.Context, arg UpdateServiceConfigParams) error {
	return p.db.UpdateServiceConfig(ctx, qm.UpdateServiceConfigParams{
		ServiceID:      arg.ServiceId,
		MinInstances:   arg.MinInstances,
		MaxInstances:   arg.MaxInstances,
		CpuPerInstance: arg.CpuPerInstance,
		RamPerInstance: arg.RamPerInstance,
	})
}

type UpdateServiceWasmParams struct {
	Wasm      []byte
	ServiceId string
}

func (p *PersistenceMaster) UpdateServiceWasm(ctx context.Context, arg UpdateServiceWasmParams) error {
	return p.db.UpdateServiceWasm(ctx, qm.UpdateServiceWasmParams{
		ServiceID: arg.ServiceId,
		Wasm:      arg.Wasm,
	})
}

type UpdateNodeAvailableCapactiyParams struct {
	AvailableCpu int64
	AvailableRam int64
	NodeId       string
}

func (p *PersistenceMaster) UpdateNodeAvailableCapacity(ctx context.Context, arg UpdateNodeAvailableCapactiyParams) error {
	return p.db.UpdateNodeAvailableCapacity(ctx, qm.UpdateNodeAvailableCapactiyParams{
		AvailableCpu: arg.AvailableCpu,
		AvailableRam: arg.AvailableRam,
		NodeID:       arg.NodeId,
	})
}
