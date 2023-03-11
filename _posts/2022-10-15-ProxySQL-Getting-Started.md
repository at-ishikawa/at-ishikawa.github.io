---
title: ProxySQL Getting Started
date: 2022-10-15T00:00:00Z
---

The official document:
- [Docker image and initial configuration](https://hub.docker.com/r/proxysql/proxysql)

Getting Started
===

See [this docker-compose.yml](/examples/proxysql/simple/docker-compose.yml) and [proxysql.cnf](/examples/proxysql/simple/proxysql.cnf) for the example of proxysql.

With this configuration, you can access a proxysql by
```
docker compose up -d
mysql -h 127.0.0.1 -u radmin -P 16032-pradmin --prompt "ProxySQL Admin> "
```

There are multiple ports
- 6032: MySQL port to access ProxySQL configurations
- 6033: Backend MySQL ports, defined in `mysql_variables.interfaces`
- 6070: Rest API port, including Prometheus endpoints. [Official document](https://proxysql.com/documentation/prometheus-exporter/)

Basics of a configuration file
---
See an [official article](https://proxysql.com/documentation/getting-started/).

There are important separated configurations
- admin_variables: for the admin interface
- [mysql_variables](): for handling the incoming MySQL traffic
- [mysql_servers](https://proxysql.com/documentation/main-runtime/#mysql_servers): for the backend servers towards which the incoming MySQL traffic
```
mysql_servers =
(
  {
    address="mysql"
    port=3306
    hostgroup=1
    max_connections=200
  }
)
```

- mysql_users: users which can connect to the proxy, and the users with which the proxy can connect to the backend servers
```
mysql_users =
(
  {
    username = "root"
    password = "password"
    default_hostgroup = 1
    max_connections=1000
    default_schema="information_schema"
    active = 1
  }
)
```

Proxy to a MySQL cluster
---

See [proxysql.cnf](/exampels/proxysql/cluster/proxysql.cnf) for the example of a configuration.

- [mysql_query_rules](https://proxysql.com/documentation/main-runtime/#mysql_query_rules): Rules to route queries to each backend MySQL.
  For example, in [this official page](articled), this is used to show how to split a read and write query though it's not completed.

Confirming which query goes to which host group, we can see [stats tables](https://proxysql.com/documentation/proxysql-configuration/).
```
SELECT * FROM stats_mysql_query_digest;
```

Also, see [an article](https://severalnines.com/blog/running-proxysql-kubernetes-service/) for configurations of ProxySQL.