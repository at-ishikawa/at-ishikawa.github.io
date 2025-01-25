---
date: "2020-08-08T00:00:00Z"
tags:
- mysql
title: Overview about MySQL Lock
---

This document is written for MySQL 5.7, so these contents may be not correct for other versions.

Lock timeout
===

When you run DDL or DML, in some cases, you may get next error.
```
ERROR 1205 (HY000): Lock wait timeout exceeded; try restarting transaction
```

This is related with `lock_wait_timeout` variable.
```
mysql root@127.0.0.1:test> show variables like 'lock_wait_timeout';
+-------------------+----------+
| Variable_name     | Value    |
+-------------------+----------+
| lock_wait_timeout | 31536000 |
+-------------------+----------+
1 row in set
Time: 0.015s
```

This variable is for timeout in seconds to acquire metadata locks.
To see more details, check [official page](https://dev.mysql.com/doc/refman/5.7/en/server-system-variables.html#sysvar_lock_wait_timeout).

### Reference
* [mysql error 1205 lock wait timeout – Quick ways to fix it](https://bobcares.com/blog/mysql-error-1205-lock-wait-timeout/)


Lock debugging and monitoring
===

show open tables command
---
Locked tables are found by `SHOW OPEN TABLES` command.
This is explained more in [the page](https://oracle-base.com/articles/mysql/mysql-identify-locked-tables).


```
mysql root@127.0.0.1:performance_schema> show open tables where in_use > 0;
+----------+-------+--------+-------------+
| Database | Table | In_use | Name_locked |
+----------+-------+--------+-------------+
0 rows in set
Time: 0.017s

mysql root@127.0.0.1:performance_schema> show open tables where `table` = 'events_waits_summary_global_by_event_name';
+--------------------+-------------------------------------------+--------+-------------+
| Database           | Table                                     | In_use | Name_locked |
+--------------------+-------------------------------------------+--------+-------------+
| performance_schema | events_waits_summary_global_by_event_name | 0      | 0           |
+--------------------+-------------------------------------------+--------+-------------+
1 row in set
Time: 0.017s
```


metadata_lock table in performance_schema
---

Locking information can be seen on `performance_schema.metadata_lock` table.
The details of this table is written in [official page](https://dev.mysql.com/doc/refman/5.7/en/performance-schema-metadata-locks-table.html).
Also, in [this post](https://www.nivas.hr/blog/2017/08/04/mysql-get_lock-problem-debug-help-mysql-performance_schema/), there is an example to use this table.

Next columns look important in `metadata_locks` table.
- OBJECT_TYPE: The type of lock. For lock acquired by users, it's `USER LEVEL LOCK`.
- OBJECT_SCHEMA, OBJECT_NAME: The same as `performance_schema.table_io_waits_summary_by_table`
- LOCK_STATUS: The status of lock like PENDING or GRANTED. See official page for the details.
- OWNER_THREAD_ID: The ID of threads in `performance_schema.threads`. In that table, we can see the process information and the query.

The example of output is like next.
```
mysql root@127.0.0.1:test> SELECT * FROM performance_schema.metadata_locks WHERE OBJECT_TYPE='USER LEVEL LOCK' \G
*************************** 1. row ***************************
          OBJECT_TYPE: USER LEVEL LOCK
        OBJECT_SCHEMA: NULL
          OBJECT_NAME: test_table
OBJECT_INSTANCE_BEGIN: 140233702308832
            LOCK_TYPE: EXCLUSIVE
        LOCK_DURATION: EXPLICIT
          LOCK_STATUS: GRANTED
               SOURCE:
      OWNER_THREAD_ID: 4586436204
       OWNER_EVENT_ID: 1
1 row in set (0.01 sec)
```

From the thread_id, we can see who's access for what query in `threads` table.

```
mysql root@127.0.0.1:test> select * from performance_schema.threads limit 30 \G
***************************[ 26. row ]***************************
THREAD_ID           | 36
NAME                | thread/sql/one_connection
TYPE                | FOREGROUND
PROCESSLIST_ID      | 11
PROCESSLIST_USER    | root
PROCESSLIST_HOST    | 172.21.0.1
PROCESSLIST_DB      | performance_schema
PROCESSLIST_COMMAND | Sleep
PROCESSLIST_TIME    | 44
PROCESSLIST_STATE   | <null>
PROCESSLIST_INFO    | select * from performance_schema.metadata_locks limit 10
PARENT_THREAD_ID    | <null>
ROLE                | <null>
INSTRUMENTED        | YES
HISTORY             | YES
CONNECTION_TYPE     | TCP/IP
THREAD_OS_ID        | 74
***************************[ 27. row ]***************************
THREAD_ID           | 41
NAME                | thread/sql/one_connection
TYPE                | FOREGROUND
PROCESSLIST_ID      | 16
PROCESSLIST_USER    | root
PROCESSLIST_HOST    | 172.21.0.1
PROCESSLIST_DB      | test
PROCESSLIST_COMMAND | Query
PROCESSLIST_TIME    | 0
PROCESSLIST_STATE   | Sending data
PROCESSLIST_INFO    | select * from performance_schema.threads limit 30
PARENT_THREAD_ID    | <null>
ROLE                | <null>
INSTRUMENTED        | YES
HISTORY             | YES
CONNECTION_TYPE     | TCP/IP
THREAD_OS_ID        | 75

27 rows in set
Time: 0.012s
```
