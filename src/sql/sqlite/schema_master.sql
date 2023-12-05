-- TODO create resource type enum

-- * TABLES

CREATE TABLE node (
    node_id TEXT NOT NULL,
    public_key TEXT NOT NULL,
    region TEXT NOT NULL,
    max_cpu INTEGER NOT NULL,
    max_ram INTEGER NOT NULL,
    available_cpu INTEGER NOT NULL,
    available_ram INTEGER NOT NULL,
    ip_address TEXT NOT NULL,
    port TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL, 
    modified_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);

-- stores 1:1 relationship between nodes and service instances
CREATE TABLE service_instance (
    service_instance_id TEXT NOT NULL,
    node_id TEXT NOT NULL,
    service_id TEXT NOT NULL,
    -- TODO: add status e.g. initialising, running, stopped
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE TABLE service_load (
    service_id TEXT NOT NULL,
    service_instance_id TEXT NOT NULL,
    cpu INTEGER NOT NULL, -- how much CPU to allocate
    ram INTEGER NOT NULL, 
    client_created_at TIMESTAMP NOT NULL, 
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE TABLE service (
    service_id TEXT NOT NULL,
    registry_url TEXT NOT NULL,  -- e.g. gcr.io...
    image_path TEXT NOT NULL, -- e.g. /company/backend-service
    min_instances INTEGER NOT NULL,
    max_instances INTEGER NOT NULL,
    cpu_per_instance INTEGER NOT NULL,
    ram_per_instance INTEGER NOT NULL,
    tags TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL, 
    modified_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE TABLE service_access_rule (   
    service_id_source TEXT NOT NULL,
    service_id_destination TEXT NOT NULL,
    CONSTRAINT unique_acr UNIQUE (service_id_source, service_id_destination)
);

CREATE TABLE service_image (
    service_id TEXT NOT NULL,
    image BLOB NOT NULL,
    modified_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    CONSTRAINT unique_service_image UNIQUE (service_id)    
);

CREATE TABLE service_scale_event (
    service_id TEXT NOT NULL,
    instance_count INTEGER NOT NULL, -- new abosulte number of instances
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);

-- maps services to gateways (n:m)
CREATE TABLE service_gateway (
    gateway_id TEXT NOT NULL,
    service_id TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);

-- TODO: maybe store if a node can act as a gateway directly with node
CREATE TABLE gateway (
    gateway_id TEXT NOT NULL,
    node_id TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE TABLE client_auth_token (
    token TEXT NOT NULL, -- hashed token
    ip_address TEXT NOT NULL, -- ip address allowed to use token
    expiry TIMESTAMP NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL, 
    CONSTRAINT unique_token UNIQUE (token)
);

CREATE TABLE client_auth_token_logs (
    token TEXT NOT NULL, -- hashed token
    ip_address TEXT NOT NULL, -- ip address used with token
    user_agent TEXT NOT NULL,   --ua user with token
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE TABLE resource_scale_event (
    resource_id TEXT NOT NULL,
    resource_count INTEGER NOT NULL, -- new absolute resource count after making scale decision
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE TABLE registry_credential (
    registry_url TEXT NOT NULL,
    username TEXT NOT NULL,
    password TEXT NOT NULL, --hashed
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    modified_at TIMESTAMP WITH TIME ZONE NOT NULL, -- for auditing purposes
    modified_by TEXT NOT NULL  -- TODO requires users
);

--* Triggers

CREATE TRIGGER prevent_updates_client_auth_token
BEFORE UPDATE ON client_auth_token
BEGIN
    SELECT RAISE(FAIL, 'Updates are not allowed on this table.');
END;

--* Indexes

CREATE INDEX idx_node_id ON node(node_id);
CREATE INDEX idx_service_id_instance ON service_instance(service_instance_id);
CREATE INDEX idx_service_id_load_instance ON service_load(service_instance_id);
CREATE INDEX idx_service_id ON service(service_id);
CREATE INDEX idx_service_id_source ON service_access_rule(service_id_source);
CREATE INDEX idx_service_id_destination ON service_access_rule(service_id_destination);
CREATE INDEX idx_service_id_image ON service_image(service_id);
CREATE INDEX idx_token_auth ON client_auth_token(token, ip_address);
CREATE INDEX idx_token_logs ON client_auth_token_logs(token, ip_address);
CREATE INDEX idx_registry_credential_url ON registry_credential(registry_url);

