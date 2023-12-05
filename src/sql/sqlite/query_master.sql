-- name: CreateNode :exec
INSERT INTO node (node_id, public_key, region, max_cpu, max_ram, available_cpu, available_ram, ip_address, port)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?);

-- name: CreateServiceLoad :exec
INSERT INTO service_load (service_instance_id, cpu, ram, client_created_at)
VALUES (?, ?, ?, ?);

-- name: CreateService :exec
INSERT INTO service(service_id, registry_url, image_path, min_instances, max_instances, cpu_per_instance, ram_per_instance, tags)
VALUES (?, ?, ?, ?, ?, ?, ?, ?);

-- name: CreateServiceInstance :exec
INSERT INTO service_instance(service_id, service_instance_id, node_id)  
VALUES (?, ?, ?);

-- name: CreateServiceAccessRule :exec
INSERT INTO service_access_rule (service_id_source, service_id_destination)
VALUES (?, ?);

-- name: CreateServiceImage :exec
REPLACE INTO service_image(service_id, image)
VALUES (?, ?);

-- name: CreateClientAuthToken :exec
INSERT INTO client_auth_token(token, ip_address, expiry)
VALUES (?, ?, ?);


-- name: UpdateImageUrl :exec
UPDATE service SET image_path = ? 
WHERE service_id = ?;


-- name: GetNodeIdsByService :many
SELECT node_id FROM service_instance
WHERE service_id = ?;

-- name: GetServiceLoadSince :many
SELECT * FROM service_load 
WHERE service_id = ? AND created_at >= ?;

-- name: GetAvailableNodes :many
SELECT * FROM node
WHERE region = ? AND available_cpu > ? AND available_ram > ?;


-- name: GetNodePublicKey :many
SELECT public_key FROM node
WHERE node_id = ?;
