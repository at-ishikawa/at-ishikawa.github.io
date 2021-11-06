---
title: Getting started
date: 2021-11-06
---


[Consul](https://www.consul.io/) is developed by Hashicorp to provide a few features like a service discovery.

This document is to touch with the consul to understand as a first step, but not for production purpose.

These pages are something I refered to.
- [HashiCorp Learn: Create a Secure Local Consul Datacenter with Docker Compose](https://learn.hashicorp.com/tutorials/consul/docker-compose-datacenter?in=consul/docker)


The first step for consul servers
---

At first, use Docker Compose to set up a few consul servers.
The `docker-compose.yml` looks like next.

```
version: '3.3'

services:
  consul-server1:
    image: consul
    networks:
      - cluster
    ports:
      - 8500:8500
    volumes:
      - ./server.json:/consul/config/server.json:ro
    command:
      - "agent"
      - "--retry-join"
      - "consul-server2"

  consul-server2:
    image: consul
    networks:
      - cluster
    volumes:
      - ./server.json:/consul/config/server.json:ro
    command:
      - "agent"
      - "--retry-join"
      - "consul-server2"

networks:
  cluster:
```

The server.json which is mounted to two containers are something like this.
```
{
  "server": true,
  "bootstrap_expect": 2,
  "ui_config": {
    "enabled": true
  },
  "addresses": {
    "http": "0.0.0.0"
  }
}
```

Then, after containers started, the Web UI can be seen on the `http://127.0.0.1/8500`.

![1st web ui](/assets/images/docs/consul/getting_started/1st_web_ui.png)


