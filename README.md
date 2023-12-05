# Dash
Lightweight distributed wasm orchestrator (to be) written in go

## Goals

- Simple to use and manage
- Lightweight (in regards to memory, cpu and LoC)
- High availability
- Autoscale services
- Minimal third party dependencies
- CLI (monitor and manage cluster)

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
- `proxy`: traefik proxy for request to service mapping
- `sysnet`: local network configurations (e.g. firewalls)
- `wasm_runtime`: runs wasm services

## V2 Goals/Ideas
- Web portal
- Internal services to service communication
- Scheduled services
- Gateways ("managed" proxy/lb)
