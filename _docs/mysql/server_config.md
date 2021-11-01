---
title: MySQL server configuration
---

Replication
===

This configuration is for the version 5.7 and it's minimum configuration in [the official document](https://dev.mysql.com/doc/refman/5.7/en/replication-configuration.html).

There are 2 types of replication setup.

* Using binary log file positions
* Using Global Transaction Identifiers

The basic configurations for binary log file positions
---
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


Logs
===

About slow queries
---
```
slow_query_log = 1
slow_query_log_file = /var/log/mysql/slow_query.log
long_query_time = 0
log_queries_not_using_indexes
```

The field `long_query_time` is the time to output a query as a slow log. The meanings of other fields are obvious.


Performance
===

Recommended configuration
---
configurations should be changed based on server or application resources.
```
[mysqld]
# The max query size
query_cache_limit=16M
# The memory size for query cache
query_cache_size=512M
# The type of query cache (0:off, 1:ON except SELECT SQL_NO_CACHE, 2:only DEMAND SELECT SQL_CACHE)
query_cache_type=1

# The max size to open files simultanously
table_open_cache = 1M

# The buffer when using sort
sort_buffer_size=4M

# The buffer to cache rows for sorting by keys
read_rnd_buffer_size=2M
```

Cache
===

Query cache
---
```
[mysqld]
# The max query size
query_cache_limit=16M
# The memory size for query cache
query_cache_size=512M
# The type of query cache (0:off, 1:ON except SELECT SQL_NO_CACHE, 2:only DEMAND SELECT SQL_CACHE)
query_cache_type=1
# The max size to open files simultanously
table_open_cache = 1M
```

See following pages for more details.
- [Performance Tuning](https://qiita.com/mamy1326/items/9c5eaee3c986cff65a55) (in Japanese)
- [Table cache](https://qiita.com/kakuka4430/items/72dc5366c9cdf65e78e9) (in Japanese)


Buffer
===

Buffer sizes
---
```
[mysqld]
# The buffer when using sort
sort_buffer_size=4M

# The buffer to cache rows for sorting by keys
read_rnd_buffer_size=2M
```

See following pages for more details
- [Performance Tuning](https://qiita.com/mamy1326/items/9c5eaee3c986cff65a55) (in Japanese)
