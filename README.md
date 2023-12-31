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

As a developer, I sought a tool that combined the ease of GCP, AWS and DO with the affordability of other cloud alternatives (Hetzner, Infomaniak). Hence, I created this project; a straightforward, scalable solution that 'just works'. 

## Who is this not for?

* Large organizations that already operate Kubernetes clusters (or similar), managed by dedicated teams
* Companies operating at "web-scale", with advanced and complex infrastructure requirements

## Third party dependencies 
For accountablility it is listed here
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

\* Written by author
## V2 Goals/Ideas
- Web portal
- Internal services to service communication
- VPN for master-to-master communication  (across datacenter / region boundaries)
- Scheduled services
- Federated cluster
- Gateways ("managed" proxy/lb)
- Stateful services
