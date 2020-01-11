MySQL server
====

Logs
----
### About slow queries
```
slow_query_log = 1
slow_query_log_file = /var/log/mysql/slow_query.log
long_query_time = 0
log_queries_not_using_indexes
```

The field `long_query_time` is the time to output a query as a slow log. The meanings of other fields are obvious.


Performance
----
### Recommended configuration
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
----
### Query cache
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
----
### Buffer sizes
```
[mysqld]
# The buffer when using sort
sort_buffer_size=4M

# The buffer to cache rows for sorting by keys
read_rnd_buffer_size=2M
```

See following pages for more details
- [Performance Tuning](https://qiita.com/mamy1326/items/9c5eaee3c986cff65a55) (in Japanese)
