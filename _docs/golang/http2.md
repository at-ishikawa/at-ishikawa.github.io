---
title: HTTP/2
---
Written in March, 2020.

`http` package in golang supports HTTP/2 protocols.
It's automatically configured.


# The features of HTTP/2
In order to understand the benefits of HTTP/2, [this document](https://developers.google.com/web/fundamentals/performance/http2), provided by google, is helpful.
It supports a lot of features including followings.
1. binary format communications
1. streaming messages
1. multiplexing
1. server push

Some examples are written in [this page](https://posener.github.io/http2/), and I checked it so much.

## Streaming
[main.go](../../examples/golang/http2/main.go) is an example for server streaming, including a server and a client codes.
The important part is calling `Flush` method, whose interface is [`http.Flusher`](https://golang.org/pkg/net/http/#Flusher) and implemented by `http.ResponseWriter`.
When this method is used, the buffered data on the server is sent to the client.


## Server push
Like `http.Flusher`, there is an interface [`http.Pusher`](https://golang.org/pkg/net/http/#Pusher), which is implemented by `http.ResponseWriter`, and we can use it for server push.
A path, http method, and headers can be configured to get requests from clients.
However, as of March 2020, `http.Client` in Go does not support server push, neither does `curl` cli.
If clients do not support server push but servers try to do it, then [`http.ErrNotSupported`](https://golang.org/src/net/http/request.go?s=1006:1204) is made.

For the details of this in Go, there is a [page](https://blog.golang.org/h2push) in an official blog.


# About examples in this page
In order to use HTTP/2, it seems TLS must be configured.
I tried to find the way to configure HTTP servers without TLS, but I couldn't find it.


# Reference
- [Introduction to HTTP/2, written by Google](https://developers.google.com/web/fundamentals/performance/http2)
- [HTTP/2 Adventure in the Go World](https://posener.github.io/http2/)
- [HTTP/2 Streaming in Golang](https://www.codemio.com/2018/03/http2-streaming-golang.html)
