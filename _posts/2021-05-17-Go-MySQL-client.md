---
date: 2021-05-17
title: Go MySQL client
---

This page is written on May 2021.


MySQL connections
===

See also [this article](https://www.programmersought.com/article/50355442570/) for further details.

MySQL Driver configurations
---

The repository for MySQL driver is https://github.com/go-sql-driver/mysql

### DSN parameters

The [DSN parameters](https://github.com/go-sql-driver/mysql#parameters) for timeout is either application layer and TCP layer and they're on client sides.

1. timeout: a dial timeout.
    - [Actual code](https://github.com/go-sql-driver/mysql/blob/bcc459a906419e2890a50fc2c99ea6dd927a88f2/connector.go#L41-L46)
1. readTimeout: Used to set timeout for reading data from I/O. This is used for [`net.Conn.SetReadDeadline`](https://golang.org/pkg/net/#Conn)
	- Usages
        - [Set ReadDeadline for a buffer](https://github.com/go-sql-driver/mysql/blob/bcc459a906419e2890a50fc2c99ea6dd927a88f2/buffer.go#L84-L88)
        - [Check connection for a first query from a connection pool](https://github.com/go-sql-driver/mysql/blob/bcc459a906419e2890a50fc2c99ea6dd927a88f2/packets.go#L113-L118)
1. writeTimeout: Used to set timeout for writing data from I/O. This is used for [`net.Conn.SetWriteDeadline`](https://golang.org/pkg/net/#Conn)
    - Usages
	    - [Code for writting a packet](https://github.com/go-sql-driver/mysql/blob/bcc459a906419e2890a50fc2c99ea6dd927a88f2/packets.go#L144-L151)


MySQL server configurations
---
Even if the client connection is disconnected, queries on MySQL servers can keep running.
On MySQL server side, some configurations like `wait_timeout` is used to close a non-interactive "idle" connection.
See [this answer](https://dba.stackexchange.com/questions/10578/client-times-out-while-mysql-query-remains-running), for example.

In order to stop running queries for longer time, [`max_execution_time`](https://dev.mysql.com/doc/refman/5.7/en/server-system-variables.html#sysvar_max_execution_time) can be used, though it's only applied to read-only SELECT statements.
See [this article](https://www.thegeekdiary.com/mysql-how-to-kill-a-long-running-query-using-max_execution_time/) for more details.
