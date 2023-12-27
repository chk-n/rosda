# Dash
Lightweight distributed container orchestrator (to be) written in go

## Goals

- Simple to use and manage
- Lightweight (in regards to memory, cpu and LoC)
- Highly available
- Self healing
- Autoscales services
- Minimal third party dependencies
- CLI for monitoring and management of cluster

## Components: master
- `admin_api`: http endpoint for CLIs and web pages to interact with
- `master_api`: dRPC endpoint for slaves to interact with (e.g. ping)
- `service_manager`: handles CRUD operations for services
- `service_state`: aggregates service metrics and logs
- `service_scaler`: scales services based on load
- `node_bouncer`: decides whether node can join cluster
- `message_queue`: in memory message queue
- `persistence`: sqlite
- `sysnet`: local network configurations (e.g. firewalls)
- `godaft`: simple leader election protocol for masters
- `config_parser`: reads and validates service configuration files

and other miscelaneous utilities

## Components: worker
- `worker_api`: dRPC endpoint for masters to interact with (e.g. CRUD operations for services)
- `persistence`: sqlite
- `sentry`: monitors service state (logs, metrics)
- `proxy`: traefik proxy service routing
- `sysnet`: local network configurations (e.g. firewalls)
- `container_runtime`: runs containers

## Third party dependencies
- [chi](https://github.com/go-chi/chi) (TODO remove and replace with built-in router)
- [crun](https://github.com/containers/crun)
- [drpc](https://github.com/storj/drpc)
- oauth lib
- [pgx](https://github.com/jackc/pgx)
- [postgresql](https://www.postgresql.org)
- [protobuf](https://github.com/golang/protobuf)
- [uuid](https://github.com/google/uuid)

for dev/tests
- [testify](https://github.com/stretchr/testify)
- [oapi-codegen](https://github.com/deepmap/oapi-codegen)
- [sqlc](https://github.com/sqlc-dev/sqlc)

## V2 Goals/Ideas
- Web portal
- Internal services to service communication
- VPN communication between master and worker (across datacenter / region boundaries)
- Scheduled services
- Gateways ("managed" proxy/lb)
