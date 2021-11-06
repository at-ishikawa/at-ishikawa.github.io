---
title: MySQL Orchestrator
---

[Orchestrator](https://github.com/openark/orchestrator) is a tool for MySQL HA and replication management.


Getting started
---

This is a simple set up document to set up an orchestrator.

I just followed these documents at first.
- https://github.com/wagnerjfr/orchestrator-mysql-replication-docker
- https://github.com/openark/orchestrator/blob/master/docs/install.md

### Prerequisite

- In order to promote a read replica to a main DB, it looks GTID based replication needs to be enabled.

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


### Support a failover

There is an [MySQL configuration about failover](https://github.com/openark/orchestrator/blob/master/docs/configuration-recovery.md#mysql-configuration).
Also, all MySQL servers have to turn on `log_slave_updates` option.


orchestrator-cli
---

In order to list up clusters
```
bash-5.1# orchestrator-client  -c clusters
main:3306
```

In order to check a topology of each cluster
```
bash-5.1# orchestrator-client  -c topology -i main:3306
main:3306             [0s,ok,5.7.36-log,rw,ROW,>>]
+ read-replica-1:3306 [0s,ok,5.7.36-log,rw,ROW,>>]
+ read-replica-2:3306 [0s,ok,5.7.36-log,rw,ROW,>>]
```


### Out of scopes
* How to set a main DB of a topology by the configuration

### Troubleshootings
#### Error when a read replica is promoted to a main DB
The example of an error message.
```
Desginated instance read-replica-2:3306 cannot take over all of its siblings. Error: 2021-11-04 04:34:36 ERROR Relocating 1 replicas of main:3306 below read-replica-2:3306 turns to be too complex; please do it manually
```

By cli, the same result happened.
```
bash-5.1# orchestrator-client -c graceful-master-takeover -a main:3306 -d read-replica-1
Desginated instance read-replica-1:3306 cannot take over all of its siblings. Error: 2021-11-04 05:37:45 ERROR Relocating 1 replicas of main:3306 below read-replica-1:3306 turns to be too complex; please do it manually
```

After GTID replication on MySQL turned on, this error was gone.

I also tried to turn on `log_slave_updates` on all MySQL servers by following one comment of [a GitHub issue](https://github.com/openark/orchestrator/issues/876), but didn't solve.
