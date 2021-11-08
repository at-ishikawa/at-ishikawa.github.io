MySQL cluster example
===

These are summary of the architecture of `docker-compose.yml`.

The list of containers
---

### MySQL 5.7 Cluster
* main: Main DB
* read-replica-1: Read replica DB 1
* read-replica-2: Read replica DB 2

### [Orchestrator](https://github.com/openark/orchestrator)
* orchestrator: Application container for orchestrator
* orchestrator-mysql: MySQL for orchestrator container

### Consul
* consul-server1:
* consul-server2:
* consul-client:

### Utilities
* dockerize:
* curl:

## The services
- `main:3306`: MySQL main DB
- `orchestrator:3000`: The Web UI of the Orchestrator
- `consul-server1:8500`: The HTTP server of the Consul
- `consul-server2:8600`: The DNS server of the Consul


Getting Started
----

How to start containers.

```
> make up
```

Image
---
Each image of MySQL contains following library
1. [gh-ost](https://github.com/github/gh-ost)


More details
---

This docker-compose is used to describe a few articles.
- [MySQL replication](https://at-ishikawa.github.io/docs/mysql/replication/)
- [gh-ost tutorial](https://at-ishikawa.github.io/docs/mysql/gh-ost/)
- [MySQL Orchestrator](https://at-ishikawa.github.io/docs/mysql/orchestrator/)
