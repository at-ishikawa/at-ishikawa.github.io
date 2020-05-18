# Getting Started
## Install packages
```
> go get all
```

## Generate files from protocol buffers
```
> make proto
```

## Run a program
### Hello world
#### Start a server
```
> make run/helloworld/server
2020/05/17 19:49:06 receive name: client name in SayHello
2020/05/17 19:49:06 receive name: client name in KeepReplyingHello, took: 56770
```

#### Connect to a server
```
> make run/helloworld/client
response: message:"Hello client name"
header: map[content-type:[application/grpc] start:[2020-05-17T19:49:06-07:00]]
response: message:"1: Hello client name"
response: message:"2: Hello client name"
response: message:"3: Hello client name"
response: message:"4: Hello client name"
response: message:"5: Hello client name"
response: message:"6: Hello client name"
response: message:"7: Hello client name"
response: message:"8: Hello client name"
response: message:"9: Hello client name"
trailer: map[duration:[56.77µs]]
client duration: 498.955µs
```

### Server streaming benchmark
1. Setup environment
```
make setup
```
1. Run benchmark
```
> make benchmark
go test -v -bench=. -benchmem -cpuprofile=cpu.out -memprofile=mem.out ./user
goos: darwin
goarch: amd64
pkg: github.com/at-ishikawa/at-ishikawa.github.io/examples/golang/grpc/user
BenchmarkPaginateUsers
BenchmarkPaginateUsers-16              1        5596781834 ns/op        15926000 B/op     226702 allocs/op
BenchmarkStreamUsers
BenchmarkStreamUsers/concurrency_=_1
BenchmarkStreamUsers/concurrency_=_1-16                        1        5758503770 ns/op        15508784 B/op     219892 allocs/op
BenchmarkStreamUsers/concurrency_=_10
BenchmarkStreamUsers/concurrency_=_10-16                       1        2026171094 ns/op        15503736 B/op     220055 allocs/op
BenchmarkStreamUsers/concurrency_=_100
BenchmarkStreamUsers/concurrency_=_100-16                      1        1906163681 ns/op        15501696 B/op     220161 allocs/op
PASS
ok      github.com/at-ishikawa/at-ishikawa.github.io/examples/golang/grpc/user  15.989s
```
1. Clean an environment
```
> make clean
docker-compose rm -s -f
Going to remove grpc_mysql_1
Removing grpc_mysql_1 ... done
```
