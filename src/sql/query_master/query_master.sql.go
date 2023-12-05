// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: query_master.sql

package query_master

import (
	"context"
	"time"
)

const createClientAuthToken = `-- name: CreateClientAuthToken :exec
INSERT INTO client_auth_token(token, ip_address, expiry)
VALUES (?, ?, ?)
`

type CreateClientAuthTokenParams struct {
	Token     string
	IpAddress string
	Expiry    time.Time
}

func (q *Queries) CreateClientAuthToken(ctx context.Context, arg CreateClientAuthTokenParams) error {
	_, err := q.db.ExecContext(ctx, createClientAuthToken, arg.Token, arg.IpAddress, arg.Expiry)
	return err
}

const createNode = `-- name: CreateNode :exec
INSERT INTO node (node_id, public_key, region, max_cpu, max_ram, available_cpu, available_ram, ip_address, port)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
`

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

func (q *Queries) CreateNode(ctx context.Context, arg CreateNodeParams) error {
	_, err := q.db.ExecContext(ctx, createNode,
		arg.NodeID,
		arg.PublicKey,
		arg.Region,
		arg.MaxCpu,
		arg.MaxRam,
		arg.AvailableCpu,
		arg.AvailableRam,
		arg.IpAddress,
		arg.Port,
	)
	return err
}

const createService = `-- name: CreateService :exec
INSERT INTO service(service_id, registry_url, image_path, min_instances, max_instances, cpu_per_instance, ram_per_instance, tags)
VALUES (?, ?, ?, ?, ?, ?, ?, ?)
`

type CreateServiceParams struct {
	ServiceID      string
	RegistryUrl    string
	ImagePath      string
	MinInstances   int64
	MaxInstances   int64
	CpuPerInstance int64
	RamPerInstance int64
	Tags           string
}

func (q *Queries) CreateService(ctx context.Context, arg CreateServiceParams) error {
	_, err := q.db.ExecContext(ctx, createService,
		arg.ServiceID,
		arg.RegistryUrl,
		arg.ImagePath,
		arg.MinInstances,
		arg.MaxInstances,
		arg.CpuPerInstance,
		arg.RamPerInstance,
		arg.Tags,
	)
	return err
}

const createServiceAccessRule = `-- name: CreateServiceAccessRule :exec
INSERT INTO service_access_rule (service_id_source, service_id_destination)
VALUES (?, ?)
`

type CreateServiceAccessRuleParams struct {
	ServiceIDSource      string
	ServiceIDDestination string
}

func (q *Queries) CreateServiceAccessRule(ctx context.Context, arg CreateServiceAccessRuleParams) error {
	_, err := q.db.ExecContext(ctx, createServiceAccessRule, arg.ServiceIDSource, arg.ServiceIDDestination)
	return err
}

const createServiceImage = `-- name: CreateServiceImage :exec
REPLACE INTO service_image(service_id, image)
VALUES (?, ?)
`

type CreateServiceImageParams struct {
	ServiceID string
	Image     []byte
}

func (q *Queries) CreateServiceImage(ctx context.Context, arg CreateServiceImageParams) error {
	_, err := q.db.ExecContext(ctx, createServiceImage, arg.ServiceID, arg.Image)
	return err
}

const createServiceInstance = `-- name: CreateServiceInstance :exec
INSERT INTO service_instance(service_id, service_instance_id, node_id)  
VALUES (?, ?, ?)
`

type CreateServiceInstanceParams struct {
	ServiceID         string
	ServiceInstanceID string
	NodeID            string
}

func (q *Queries) CreateServiceInstance(ctx context.Context, arg CreateServiceInstanceParams) error {
	_, err := q.db.ExecContext(ctx, createServiceInstance, arg.ServiceID, arg.ServiceInstanceID, arg.NodeID)
	return err
}

const createServiceLoad = `-- name: CreateServiceLoad :exec
INSERT INTO service_load (service_instance_id, cpu, ram, client_created_at)
VALUES (?, ?, ?, ?)
`

type CreateServiceLoadParams struct {
	ServiceInstanceID string
	Cpu               int64
	Ram               int64
	ClientCreatedAt   time.Time
}

func (q *Queries) CreateServiceLoad(ctx context.Context, arg CreateServiceLoadParams) error {
	_, err := q.db.ExecContext(ctx, createServiceLoad,
		arg.ServiceInstanceID,
		arg.Cpu,
		arg.Ram,
		arg.ClientCreatedAt,
	)
	return err
}

const getAvailableNodes = `-- name: GetAvailableNodes :many
SELECT node_id, public_key, region, max_cpu, max_ram, available_cpu, available_ram, ip_address, port, created_at, modified_at FROM node
WHERE region = ? AND available_cpu > ? AND available_ram > ?
`

type GetAvailableNodesParams struct {
	Region       string
	AvailableCpu int64
	AvailableRam int64
}

func (q *Queries) GetAvailableNodes(ctx context.Context, arg GetAvailableNodesParams) ([]Node, error) {
	rows, err := q.db.QueryContext(ctx, getAvailableNodes, arg.Region, arg.AvailableCpu, arg.AvailableRam)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Node
	for rows.Next() {
		var i Node
		if err := rows.Scan(
			&i.NodeID,
			&i.PublicKey,
			&i.Region,
			&i.MaxCpu,
			&i.MaxRam,
			&i.AvailableCpu,
			&i.AvailableRam,
			&i.IpAddress,
			&i.Port,
			&i.CreatedAt,
			&i.ModifiedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getNodeIdsByService = `-- name: GetNodeIdsByService :many
SELECT node_id FROM service_instance
WHERE service_id = ?
`

func (q *Queries) GetNodeIdsByService(ctx context.Context, serviceID string) ([]string, error) {
	rows, err := q.db.QueryContext(ctx, getNodeIdsByService, serviceID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var node_id string
		if err := rows.Scan(&node_id); err != nil {
			return nil, err
		}
		items = append(items, node_id)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getNodePublicKey = `-- name: GetNodePublicKey :many
SELECT public_key FROM node
WHERE node_id = ?
`

func (q *Queries) GetNodePublicKey(ctx context.Context, nodeID string) ([]string, error) {
	rows, err := q.db.QueryContext(ctx, getNodePublicKey, nodeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var public_key string
		if err := rows.Scan(&public_key); err != nil {
			return nil, err
		}
		items = append(items, public_key)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getServiceLoadSince = `-- name: GetServiceLoadSince :many
SELECT service_id, service_instance_id, cpu, ram, client_created_at, created_at FROM service_load 
WHERE service_id = ? AND created_at >= ?
`

type GetServiceLoadSinceParams struct {
	ServiceID string
	CreatedAt time.Time
}

func (q *Queries) GetServiceLoadSince(ctx context.Context, arg GetServiceLoadSinceParams) ([]ServiceLoad, error) {
	rows, err := q.db.QueryContext(ctx, getServiceLoadSince, arg.ServiceID, arg.CreatedAt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ServiceLoad
	for rows.Next() {
		var i ServiceLoad
		if err := rows.Scan(
			&i.ServiceID,
			&i.ServiceInstanceID,
			&i.Cpu,
			&i.Ram,
			&i.ClientCreatedAt,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateImageUrl = `-- name: UpdateImageUrl :exec
UPDATE service SET image_path = ? 
WHERE service_id = ?
`

type UpdateImageUrlParams struct {
	ImagePath string
	ServiceID string
}

func (q *Queries) UpdateImageUrl(ctx context.Context, arg UpdateImageUrlParams) error {
	_, err := q.db.ExecContext(ctx, updateImageUrl, arg.ImagePath, arg.ServiceID)
	return err
}