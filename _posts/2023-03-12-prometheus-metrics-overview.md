---
title: Prometheus Metrics Overview
date: 2023-03-12
tags:
  - prometheus
  - kubernetes
---


# Kubernetes Metrics


## Node metrics

These metrics require [Node Exporter]()

- CPU utilization per node: `1 - (avg by (instance)(rate(node_cpu_seconds_total{mode="idle"}[5m])))`
- Memory utilization per node: `1 - (node_memory_MemAvailable_bytes / node_memory_MemTotal_bytes)`
- Disk utilization per node: `1 - (node_filesystem_avail_bytes / node_filesystem_size_bytes)`

## Pod metrics

- CPU requests per namespace: `sum by (exported_namespace)(kube_pod_container_resource_requests{resource="cpu", unit="core"})`
- Memory requests per namespace: `sum by (exported_namespace)(kube_pod_container_resource_requests{resource="memory", unit="byte"})`


### Reference
Following some articles including followings
- [Stackoverflow: CPU usage for each node in prometheus](https://stackoverflow.com/a/66263640)
- [Prometheus cheetsheet](https://blog.ruanbekker.com/cheatsheets/prometheus/)
- [Average memory usage query prometheus](https://stackoverflow.com/questions/48835035/average-memory-usage-query-prometheus)
