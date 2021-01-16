---
title:
date: 2021-01-15T19:00:00-07:00
categories:
  - http
  - websocket
tags:
  - http
  - websocket
---

There are some use cases that server side want to send data to clients, like real time chat application.
In this post, I looked around what the possible solutions like web socket or [Comet](https://en.wikipedia.org/wiki/Comet_(programming)) and what the pros and cons for them.


Quick comparisons
===
There are a few ideas that I came up with and also knew when I searched for the Internet.

* Web socket
* gRPC streaming
* HTTP/2 stream (without gRPC)
* HTTP/2 + SSE

Note that there is little information for HTTP/2 stream, but I guessed it can be used for real time communications.

|                | Web socket                          | gRPC        | HTTP/2 stream                 | HTTP/2 + SSE       |
| Direction      | Bidirection                         | Bidirection | Client/Server + Server stream |                    |
| Data Format    | Text or binary                      | Binary      | Binary                        | Binary             |
| End of message | Fin bit                             | Supported   | No native support             |                    |
| Reconnect      | Recreate the connection by yourself | ?           | ?                             | ?                  |
| Mobile app     | OK                                  | ?           | OK                            | No standardization |
| Web support    | OK                                  | Limited     | ?                             | OK                 |

Besides the above, it's better to check whether a load balancer supports a web socket or gRPC protocol.

Notes for each technology
===

Websocket
---
* Communication: Bidirectional
* Buffered data: Fragmented frame (Fin bit) is used to recognize the end of a message from buffered data.

Example codes of Golang and JavaScript is [here](./examples/golang/websocket).
Also, web socket API on a browser doesn't support sending a header, but it's included in the protocol spec. So, we can solve it by a workaround liike [this page](https://yeti.co/blog/token-based-header-authentication-for-websockets-behind-nodejs/).

HTTP/2 stream
---
* Communication: Client/Server + Server Push


gRPC streaming
---
* Support all of client, server, or bidirectional streaming
* Web supports server streaming only when grpc web and base64 encoded content-type is used.


Server Sent Events
---
Server sent events are the technology to keep the persistent connection between a clients and a server side.
This is a standardized for W3C for web but no standardization for a mobile app.
See [this page](https://javascript.info/server-sent-events)


Other technologies
===

HTTP/2 server push
---
HTTP/2 server push is the technology for a web browser and not for an application code.
And server push is used to fetch assets before clients ask, like CSS or JavaScript files for a requested HTML file.
See [wiki](https://en.wikipedia.org/wiki/HTTP/2_Server_Push) for more details about HTTP/2 server push.


References
===
* [Wikipedia: HTTP/2 Server Push](https://en.wikipedia.org/wiki/HTTP/2_Server_Push)
* [Rambling Comments: WebSockets is a stream, not a message based protocol](http://www.lenholgate.com/blog/2011/07/websockets-is-a-stream-not-a-message-based-protocol.html)
* [SessionStack Blog: How JavaScript works: Deep dive into WebSockets and HTTP/2 with SSE + how to pick the right path](https://blog.sessionstack.com/how-javascript-works-deep-dive-into-websockets-and-http-2-with-sse-how-to-pick-the-right-path-584e6b8e3bf7)
* [API friends: What is Server-Sent Events?](https://apifriends.com/api-streaming/server-sent-events/)
* [appfutura: XMPP or Websockets: What should you choose for your Mobile Chat Application](https://www.appfutura.com/blog/xmpp-or-websockets-what-should-you-choose-for-your-mobile-chat-application/)
