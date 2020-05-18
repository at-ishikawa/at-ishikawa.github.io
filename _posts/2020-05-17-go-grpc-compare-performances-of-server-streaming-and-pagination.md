---
title: Compare performances of gRPC server streaming and pagination of unary RPC
date: 2020-05-17T19:00:00-07:00
categories:
  - golang
tags:
  - golang
  - grpc
  - server-streaming
---

There are some cases that we wanna fetch all records that are matching with a certain condition from an other gRPC server.
In these cases, there are at least 2 possible solutions to implement it, by gRPC server streaming or pagination.
In this post, in order to check the benefit of the performance of gRPC server streaming over pagination, running benchmarks for the both of pagination and server streaming, and checked the performances.

Summaries of a result and analysis
===
Result wasn't very good, and it might be better to run benchmarks on multiple machines and have more data with faster queries.

* Contrary to expectations that using server streaming has better performances, the performance wasn't improved under benchmark environment.
    * It might be because the benchmark environment was local machine and it couldn't check the advantages of gRPC multiplexing responses.
* Not surprisingly, if a gRPC server fetches records concurrently only on server streaming, it has significant advantages. And this might be the most important advantage for server streaming against pagination.
    * For pagination, a client side has to implement concurrent logics because the number of data that can be fetched by one request is limited. However, in many cases, pagination needs the previous result to get the data for next pagination request, like next token. So it's hard to implement concurrent logics.
	* For server streaming, all of logics are is on a server side. Because a server can control the logics and how many data they can reply to the client, it's easier to optimze.


Details
===

Benchmark environments
---
* Run MySQL server as a docker container, and gRPC servers locally
* In MySQL, there are 110K records.
* Get multiple responses from gRPC RPCs to fetch all data that are matching with a condition
* Fetch 1000 items at once for one query of pagination, and one response of gRPC server streaming
* Run tests 10 times. One test took too much time so couldn't increase more.
* gRPC servers and clients are implemented in Golang.


Benchmark implementation details
---
* From MySQL, a gRPC server fetches records by the query `SELECT * FROM users WHERE name LIKE '%${keyword}%' ORDER BY id DESC LIMIT {$limit} OFFSET ${offset}`.
* In order to enable concurrency on gRPC server streaming, a gRPC server runs the query to fetch the total count `SELECT COUNT(1) AS count FROM users WHERE name LIKE '%${keyword}%'`.
    * For pagination, it runs only on the 1st request for pagination.
* The proto for pagination and server streaming RPCs are [here](../../examples/golang/grpc/proto/server_streaming_benchmark.proto).
* The details of codes are
    * [server](/examples/golang/grpc/internal/user/server.go)
	* [benchmark](/examples/golang/grpc/internal/user/benchmark_test.go)
	* [seed](/examples/golang/grpc/cmd/seed/user/main.go)

Benchmark result
---

### The case when 18K records are fetched
---
```
mysql root@127.0.0.1:test> select count(1) from users where name like '%ff%' limit 5;
+----------+
| count(1) |
+----------+
| 18081    |
+----------+
1 row in set
Time: 0.926s
```

```
> go test -v -bench=. -benchtime=10x -benchmem -cpuprofile=cpu.out -memprofile=mem.out ./user
goos: darwin
goarch: amd64
pkg: github.com/at-ishikawa/at-ishikawa.github.io/examples/golang/grpc/user
BenchmarkPaginateUsers
BenchmarkPaginateUsers-16             10        5729014144 ns/op        15748736 B/op     222302 allocs/op
BenchmarkStreamUsers
BenchmarkStreamUsers/concurrency_=_1
BenchmarkStreamUsers/concurrency_=_1-16                       10        5767206314 ns/op        15387366 B/op     219532 allocs/op
BenchmarkStreamUsers/concurrency_=_10
BenchmarkStreamUsers/concurrency_=_10-16                      10        2086846980 ns/op        15504884 B/op     219983 allocs/op
BenchmarkStreamUsers/concurrency_=_100
BenchmarkStreamUsers/concurrency_=_100-16                     10        1954722516 ns/op        15444297 B/op     219882 allocs/op
PASS
ok      github.com/at-ishikawa/at-ishikawa.github.io/examples/golang/grpc/user  171.140s
```

### The case when 1039K records are fetched
In this case, the benchmark couldn't run even 5 times, so run the test only once.
```
mysql root@127.0.0.1:test> select count(1) from users where name like '%e%' limit 5;
+----------+
| count(1) |
+----------+
| 1039693  |
+----------+
1 row in set
Time: 0.936s
```

```
> go test -v -bench=. -benchmem -cpuprofile=cpu.out -memprofile=mem.out ./user
goos: darwin
goarch: amd64
pkg: github.com/at-ishikawa/at-ishikawa.github.io/examples/golang/grpc/user
BenchmarkPaginateUsers
BenchmarkPaginateUsers-16              1        191732946675 ns/op      889266160 B/op  12768022 allocs/op
BenchmarkStreamUsers
BenchmarkStreamUsers/concurrency_=_1
BenchmarkStreamUsers/concurrency_=_1-16                        1        189976218377 ns/op      865303704 B/op  12596571 allocs/op
BenchmarkStreamUsers/concurrency_=_10
BenchmarkStreamUsers/concurrency_=_10-16                       1        41810204876 ns/op       870448160 B/op  12613615 allocs/op
BenchmarkStreamUsers/concurrency_=_100
BenchmarkStreamUsers/concurrency_=_100-16                      1        39546694589 ns/op       870049200 B/op  12618855 allocs/op
PASS
ok      github.com/at-ishikawa/at-ishikawa.github.io/examples/golang/grpc/user  463.391s
```
