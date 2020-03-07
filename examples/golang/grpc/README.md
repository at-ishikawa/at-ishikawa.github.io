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
### Start a server
```
> make run-server
```

### Connect to a server
```
> make run-client
response: message:"Hello client name"
header: map[content-type:[application/grpc] start:[2020-03-06T23:32:09-08:00]]
response: message:"1: Hello client name"
response: message:"2: Hello client name"
response: message:"3: Hello client name"
response: message:"4: Hello client name"
response: message:"5: Hello client name"
response: message:"6: Hello client name"
response: message:"7: Hello client name"
response: message:"8: Hello client name"
response: message:"9: Hello client name"
trailer: map[duration:[63.069µs]]
client duration: 316.54µs
```

## Reference
- Go support for Protocol Buffers: https://github.com/golang/protobuf
- The examples of gRPC for Go: https://github.com/grpc/grpc-go/blob/master/examples/README.md
- stream gRPC: https://gist.github.com/jzelinskie/10ceca82f4f5085c106d
