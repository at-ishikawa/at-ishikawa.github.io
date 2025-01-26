---
date: "2023-01-16T00:00:00Z"
tags:
- mysql
title: Working around MySQL lock metadata
---

There are multiple documents about innodb locks on MySQL 5.7:
- [InnoDB locking](https://dev.mysql.com/doc/refman/5.7/en/innodb-locking.html)
- [Locks Set by Different SQL Statements in InnoDB](https://dev.mysql.com/doc/refman/5.7/en/innodb-locks-set.html)
- [Using InnoDB Transaction and Locking Information](https://dev.mysql.com/doc/refman/5.7/en/innodb-information-schema-examples.html)


# Granularity of locks

There are 4 types of granularity of locks

- Shared lock
- Exclusive lock
- Intention shard lock
- Intention exclusive lock


# Types of locks

- Record locks
- Gap locks
- Next-Key locks
    - the value in the index for the next-key lock indicates as "supreme" pseudo-record, which is not a real value
- Insert Intention locks
- AUTO-INC locks
    - This is a table lock, and how to lock the table depending on the configuration of [`innodb_autoinc_lock_mode`](https://dev.mysql.com/doc/refman/5.7/en/innodb-auto-increment-handling.html#innodb-auto-increment-lock-modes)


# Lock tables

- information_schema.innodb_trx
- information_schema.innodb_locks
- information_schema.innodb_lock_waits


# Use cases

## Update queries without an index
Create a test table and insert a few records.

```sql
create database test;
use test;

create table test1(id INTEGER AUTO_INCREMENT PRIMARY KEY, name VARCHAR(255) NOT NULL, created_at DATETIME NOT NULL, updated_at DATETIME NOT NULL);
insert into test1(name, created_at, updated_at) values('test', NOW(), NOW());
insert into test1(name, created_at, updated_at) values('test 2', NOW(), NOW());
insert into test1(name, created_at, updated_at) values('test 3', NOW(), NOW());
insert into test1(name, created_at, updated_at) values('test 4', NOW(), NOW());
select * from test1;
```

Then the last selct result would be like:
```sql
+----+--------+---------------------+---------------------+
| id | name   | created_at          | updated_at          |
+----+--------+---------------------+---------------------+
|  1 | test   | 2023-01-17 02:02:35 | 2023-01-17 02:02:35 |
|  2 | test 2 | 2023-01-17 02:02:36 | 2023-01-17 02:02:36 |
|  3 | test 3 | 2023-01-17 02:02:38 | 2023-01-17 02:02:38 |
|  4 | test 4 | 2023-01-17 02:02:40 | 2023-01-17 02:02:40 |
+----+--------+---------------------+---------------------+
4 rows in set (0.01 sec)
```

Checking what the lock results will be:

1. session A
    ```sql
    BEGIN;
    update test1 set name = 'test 2 changed' where name = 'test 2';
    ```

1. session B: This will be locked because of the query on session A.
    ```sql
    insert into test1(name, created_at, updated_at) values('test 5', NOW(), NOW());
    ```

Then we can see lock information like followings:

```sql
mysql> select * from information_schema.innodb_locks \G
*************************** 1. row ***************************
    lock_id: 1814:24:3:1
lock_trx_id: 1814
  lock_mode: X
  lock_type: RECORD
 lock_table: `test`.`test1`
 lock_index: PRIMARY
 lock_space: 24
  lock_page: 3
   lock_rec: 1
  lock_data: supremum pseudo-record
*************************** 2. row ***************************
    lock_id: 1811:24:3:1
lock_trx_id: 1811
  lock_mode: X
  lock_type: RECORD
 lock_table: `test`.`test1`
 lock_index: PRIMARY
 lock_space: 24
  lock_page: 3
   lock_rec: 1
  lock_data: supremum pseudo-record
2 rows in set, 1 warning (0.00 sec)
```

You can see lock_data is the supremum pseudo-record, which is not a real record created for next-key lock.

Other tables look like followings:

```sql
mysql> select * from information_schema.innodb_trx \G
*************************** 1. row ***************************
                    trx_id: 1814
                 trx_state: LOCK WAIT
               trx_started: 2023-01-17 02:24:59
     trx_requested_lock_id: 1814:24:3:1
          trx_wait_started: 2023-01-17 02:24:59
                trx_weight: 2
       trx_mysql_thread_id: 4
                 trx_query: insert into test1(name, created_at, updated_at) values('test 5', NOW(), NOW())
       trx_operation_state: inserting
         trx_tables_in_use: 1
         trx_tables_locked: 1
          trx_lock_structs: 2
     trx_lock_memory_bytes: 1136
           trx_rows_locked: 1
         trx_rows_modified: 0
   trx_concurrency_tickets: 0
       trx_isolation_level: REPEATABLE READ
         trx_unique_checks: 1
    trx_foreign_key_checks: 1
trx_last_foreign_key_error: NULL
 trx_adaptive_hash_latched: 0
 trx_adaptive_hash_timeout: 0
          trx_is_read_only: 0
trx_autocommit_non_locking: 0
*************************** 2. row ***************************
                    trx_id: 1811
                 trx_state: RUNNING
               trx_started: 2023-01-17 02:03:26
     trx_requested_lock_id: NULL
          trx_wait_started: NULL
                trx_weight: 3
       trx_mysql_thread_id: 3
                 trx_query: NULL
       trx_operation_state: NULL
         trx_tables_in_use: 0
         trx_tables_locked: 1
          trx_lock_structs: 2
     trx_lock_memory_bytes: 1136
           trx_rows_locked: 5
         trx_rows_modified: 1
   trx_concurrency_tickets: 0
       trx_isolation_level: REPEATABLE READ
         trx_unique_checks: 1
    trx_foreign_key_checks: 1
trx_last_foreign_key_error: NULL
 trx_adaptive_hash_latched: 0
 trx_adaptive_hash_timeout: 0
          trx_is_read_only: 0
trx_autocommit_non_locking: 0
2 rows in set (0.00 sec)
```

```sql
mysql> select * from information_schema.innodb_lock_waits \G
*************************** 1. row ***************************
requesting_trx_id: 1815
requested_lock_id: 1815:24:3:1
  blocking_trx_id: 1811
 blocking_lock_id: 1811:24:3:1
1 row in set, 1 warning (0.01 sec)
```
