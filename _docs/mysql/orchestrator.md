---
title:
---

Orchestrator
===

[orchestrator](https://github.com/openark/orchestrator) is a tool for MySQL HA and replication management.

Getting started
---

This is a simple set up document to set up an orchestrator.

I just followed these documents at first.
- https://github.com/wagnerjfr/orchestrator-mysql-replication-docker
- https://github.com/openark/orchestrator/blob/master/docs/install.md

### Build a docker image

At first, build a docker image.
I couldn't find which docker image in DockerHub can be used in a docker-compose.

```
git clone https://github.com/openark/orchestrator
cd orchestrator
docker build . -t docker/Dockerfile -t orchestrator:latest
```

### Set up an MySQL server for orchestrator

Set up a MySQL server for Orchestrator and create a user and a database.

```
CREATE DATABASE IF NOT EXISTS orchestrator;
CREATE USER 'orchestrator'@'%' IDENTIFIED BY 'password';
GRANT ALL PRIVILEGES ON `orchestrator`.* TO 'orchestrator'@'%';
```

### Set up an user and permissions on each MySQL server in a MySQL cluster

In each MySQL server of a topology, create an orchestrator user to manage by Orchestrator.

```
CREATE USER 'orchestrator'@'%' IDENTIFIED BY 'password';
GRANT SUPER, PROCESS, REPLICATION SLAVE, RELOAD ON *.* TO 'orchestrator'@'%';
GRANT SELECT ON mysql.slave_master_info TO 'orchestrator'@'%';
```

### Set up an user for orchestrator

Configure MySQL settings for orchestrator MySQL server and also managed MySQL topology servers in `/etc/orchestrator.conf.json`.

```
{
    "MySQLOrchestratorHost": "orchestrator_mysql_db",
    "MySQLOrchestratorPort": 3306,
    "MySQLOrchestratorDatabase": "orchestrator",
    "MySQLOrchestratorUser": "orchestrator",
    "MySQLOrchestratorPassword": "orch_backend_password",

    "MySQLTopologyUser": "orchestrator",
    "MySQLTopologyPassword": "orch_topology_password"
}
```

Then start the Orchestrator with a few Docker containers.


### Out of scopes
* How to set a main DB of a topology by the configuration
* How to solve "NoFailoverSupportStructureWarning"
