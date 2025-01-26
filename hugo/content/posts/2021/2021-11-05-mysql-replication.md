---
date: "2021-11-05T00:00:00Z"
tags:
- mysql
title: MySQL Replication
---

This configuration is for the version 5.7 and it's minimum configuration in [the official document](https://dev.mysql.com/doc/refman/5.7/en/replication-configuration.html).

There are 2 types of replication setup.

* Using binary log file positions
* Using Global Transaction Identifiers

Binary logging replication
---

There are a couple of important configurations
* `server-id`: The unique ID on each server, and must be a positive integer between 1 and 2^32-1. The default value is 0.
* `log-bin`: the file name for binary logs. This is required to enable replications using binary loggings.
    Binary loggings are not required on read replicas, but they can be used for data backups and crash recovery.

Besides,
* Create a user who can connect to the main db from each replica and grant **REPLICATION SLAVE** permission to them.
  Note that it's better to create a separate user because **this username and password are stored in plain text in the replication metadata repositories.**

### Main DB configuration

```
[mysqld]
server-id = 1
log-bin = mysql-bin
```

Create a user for replication.

```
CREATE USER 'repl'@'%' IDENTIFIED BY 'password';
GRANT REPLICATION SLAVE ON *.* TO 'repl'@'%';
```

### Replica DB configuration

```
[mysqld]
server-id = 11
log-bin = mysql-bin
```

In order to connect to the main DB,
```
STOP SLAVE;
CHANGE MASTER TO
    MASTER_HOST='$mysql_main_server',
    MASTER_USER='repl',
    MASTER_PASSWORD='password';
START SLAVE;
```


Replication with GTID
---

Next YouTube video may be helpful to understand the operation and configuration on GTID replication
<iframe width="560" height="315" src="https://www.youtube.com/embed/cVymVWZ7SPw" title="YouTube video player" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture" allowfullscreen></iframe>

There are a couple of configurations to enable GTID based replication
* [gtid_mode](https://dev.mysql.com/doc/refman/5.7/en/replication-options-gtids.html#sysvar_gtid_mode) = on: GTID based logging is enabled
* [enforce_gtid_consistency](https://dev.mysql.com/doc/refman/5.7/en/replication-options-gtids.html#sysvar_enforce_gtid_consistency): allow execution of only statements that can be safely logged using a GTID



The other configuration
---

### [Replication format]
See an [official document](https://dev.mysql.com/doc/refman/5.7/en/replication-sbr-rbr.html) for the details.
This can be configured as a [`binlog_format`](https://dev.mysql.com/doc/refman/5.7/en/replication-options-binary-log.html#sysvar_binlog_format) in a configuration file.

- Statement-based replication (SBR) (binlog_format=STATEMENT)
    - Some statements cannot be replicated
    - Less data
    - If replicas have different triggers from main DB, then a SBR is required. ([Ref](https://dev.mysql.com/doc/refman/5.7/en/replication-features-triggers.html))
- Row-based replication (RBR) (binlog_format=ROW)
    - All changes can be replicated
    - More data
- Mixed (binlog_format=MIXED)
    - Use statement-based replication as default
    - Use row-based replication only if the statement-based replication cannot guarantee the proper result

### TODOs
* Add [log_slave_updates](https://dev.mysql.com/doc/refman/5.7/en/replication-options-binary-log.html#sysvar_log_slave_updates)
* Add [binlog_row_image](https://dev.mysql.com/doc/refman/5.7/en/replication-options-binary-log.html#sysvar_binlog_row_image)
