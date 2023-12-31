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

## Why this project?

As a developer, I sought a tool that combined the ease of GCP, AWS and DO with the affordability of other cloud alternatives (Hetzer, Infomaniak). Hence, I created this project; a straightforward, scalable solution that 'just works'. 

## Who is this not for?

* Large organizations that already operate Kubernetes clusters (or similar), managed by dedicated teams
* Companies operating at "web-scale", with advanced and complex infrastructure requirements

## Components: master
- `admin_api`: http endpoint for CLIs and web pages to interact with
- `master_api`: dRPC endpoint for workers to interact with (e.g. ping)
- `service_manager`: handles CRUD operations for services
- `sentry_receiver`: aggregates service metrics and logs
- `service_scaler`: scales services based on load
- `node_bouncer`: decides whether node can join cluster
- `persistence`: sql data layer
- `sysnet`: local network configurations (e.g. firewalls)
- simple leader election protocol for masters
- `config_parser`: reads and validates service configuration files

and other miscelaneous utilities

## Components: worker
- `worker_api`: dRPC endpoint for masters to interact with (e.g. CRUD operations for services)
- `persistence`: sqlite
- `sentry_collector`: monitors service state (logs, metrics)
- `proxy`: traefik proxy service routing
- `sysnet`: local network configurations (e.g. firewalls)
- `container_runtime`: runs containers

## Third party dependencies
- [chi](https://github.com/go-chi/chi) (TODO remove and replace with built-in router)
- [crun](https://github.com/containers/crun)
- [drpc](https://github.com/storj/drpc)
- [gomemq](https://github.com/chk-n/gomemq)*
- oauth lib
- [retry](https://github.com/chk-n/retry)* 
- [pgx](https://github.com/jackc/pgx)
- [postgresql](https://www.postgresql.org)
- [protobuf](https://github.com/golang/protobuf)
- proxy/lb
- [uuid](https://github.com/google/uuid)

for dev/tests
- [testify](https://github.com/stretchr/testify)
- [oapi-codegen](https://github.com/deepmap/oapi-codegen)
- [sqlc](https://github.com/sqlc-dev/sqlc)

* Written by author
## V2 Goals/Ideas
- Web portal
- Internal services to service communication
- VPN for master-to-master communication  (across datacenter / region boundaries)
- Scheduled services
- Federated cluster
- Gateways ("managed" proxy/lb)
- Stateful services
