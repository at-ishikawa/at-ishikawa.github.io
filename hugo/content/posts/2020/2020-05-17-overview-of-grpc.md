---
date: "2020-05-17T00:00:00Z"
title: Overview of gRPC
---

Written in March, 2020.

gRPC is the protocol using HTTP/2. For the details, please take a look for a [official page](https://grpc.io/docs/guides/).
This page explains how to use it in go and how it behaves.

# Types
There are some types for RPC.

1. unary RPC
2. server stream RPC
3. client stream RPC
4. bidirectional stream RPC

Using stream RPC, multiple messages can be sent on a single TCP connection.

# Examples
Examples include unary RPC(SayHello) and server stream RPC(KeepReplyingHello).
1. [server/main.go](/examples/golang/grpc/cmd/helloworld/server/main.go)
2. [client/main.go](/examples/golang/grpc/cmd/helloworld/client/main.go)


# Reference
- [Go support for Protocol Buffers](https://github.com/golang/protobuf)
- [The examples of gRPC for Go](https://github.com/grpc/grpc-go/blob/master/examples/README.md)
- [The example gist of jzelinskie/client.go for server streaming](https://gist.github.com/jzelinskie/10ceca82f4f5085c106d)
