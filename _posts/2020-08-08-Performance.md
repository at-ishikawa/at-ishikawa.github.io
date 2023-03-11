---
date: 2020-08-08
title: Performance
---

This document is written for MySQL 5.7, so these contents may be not correct for other versions.
In this page, [performance_schema](https://dev.mysql.com/doc/refman/5.7/en/performance-schema.html) is mainly discussed.

Latency monitorings
===

There are some tables in `performance_schema` to check the number of queries for the table.
If you are interested in other information, please see other web pages like ["What Does I/O Latencies and Bytes Mean in the Performance and sys Schemas?"](https://mysql.wisborg.dk/2018/08/05/what-does-io-latencies-and-bytes-mean-in-the-performance-and-sys-schemas/).

table_io_waits_summary_by_table table in performance_schema
---

The details of this table is described in [official page](https://dev.mysql.com/doc/refman/5.7/en/table-waits-summary-tables.html).
This table stores all I/O wait events, including select, insert, update, and delete DMLs.
- OBJECT_SCHEMA: Database name for a table I/O
- OBJECT_NAME: Table name for a table I/O
- COUNT_FETCH: The number of select queries
- COUNT_INSERT: The number of insert queries
- COUNT_UPDATE: The number of update queries
- COUNT_DELETE: The number of delete queries
- SUM\_TIMER\_\*, MIN\_TIMER\_\*, AVG\_TIMER\_\*, and MAX\_TIMER\_\*: picoseconds time

There is a function `sys.format_time` to convert latency to human readable format.

```
mysql root@127.0.0.1:performance_schema> select * from table_io_waits_summary_by_table \G
***************************[ 1. row ]***************************
OBJECT_TYPE      | TABLE
OBJECT_SCHEMA    | test
OBJECT_NAME      | test_table
COUNT_STAR       | 6
SUM_TIMER_WAIT   | 270431172
MIN_TIMER_WAIT   | 18428064
AVG_TIMER_WAIT   | 45071862
MAX_TIMER_WAIT   | 98592144
COUNT_READ       | 3
SUM_TIMER_READ   | 38676750
MIN_TIMER_READ   | 18428064
AVG_TIMER_READ   | 12891972
MAX_TIMER_READ   | 20248686
COUNT_WRITE      | 3
SUM_TIMER_WRITE  | 231754422
MIN_TIMER_WRITE  | 64047030
AVG_TIMER_WRITE  | 77251335
MAX_TIMER_WRITE  | 98592144
COUNT_FETCH      | 3
SUM_TIMER_FETCH  | 38676750
MIN_TIMER_FETCH  | 18428064
AVG_TIMER_FETCH  | 12891972
MAX_TIMER_FETCH  | 20248686
COUNT_INSERT     | 2
SUM_TIMER_INSERT | 162639174
MIN_TIMER_INSERT | 64047030
AVG_TIMER_INSERT | 81319587
MAX_TIMER_INSERT | 98592144
COUNT_UPDATE     | 1
SUM_TIMER_UPDATE | 69115248
MIN_TIMER_UPDATE | 69115248
AVG_TIMER_UPDATE | 69115248
MAX_TIMER_UPDATE | 69115248
COUNT_DELETE     | 0
SUM_TIMER_DELETE | 0
MIN_TIMER_DELETE | 0
AVG_TIMER_DELETE | 0
MAX_TIMER_DELETE | 0

1 row in set
Time: 0.011s
mysql root@127.0.0.1:performance_schema> select object_name, count_fetch, sys.format_time(max_timer_fetch), sys.format_time(min_timer_fetch) from table_io_waits_summary_by_table \G
***************************[ 1. row ]***************************
object_name                      | test_table
count_fetch                      | 19
sys.format_time(max_timer_fetch) | 24.73 s
sys.format_time(min_timer_fetch) | 13.46 us

1 row in set
Time: 0.011s
```


schema_table_statistics table in sys
---

The details is described in [this page](https://dev.mysql.com/doc/refman/5.7/en/sys-schema-table-statistics.html).
The view is to summarize table statistics.

```
mysql root@127.0.0.1:performance_schema> select * from sys.schema_table_statistics \G
***************************[ 1. row ]***************************
table_schema      | test
table_name        | test_table
total_latency     | 270.43 us
rows_fetched      | 3
fetch_latency     | 38.68 us
rows_inserted     | 2
insert_latency    | 162.64 us
rows_updated      | 1
update_latency    | 69.12 us
rows_deleted      | 0
delete_latency    | 0 ps
io_read_requests  | 14
io_read           | 1.54 KiB
io_read_latency   | 37.10 us
io_write_requests | 32
io_write          | 196.77 KiB
io_write_latency  | 300.59 us
io_misc_requests  | 29
io_misc_latency   | 11.26 ms

1 row in set
Time: 0.042s
```
