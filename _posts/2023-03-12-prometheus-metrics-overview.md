---
title: Prometheus Metrics Overview
date: 2023-03-12
last_modified_at: 2023-04-03
tags:
  - prometheus
  - kubernetes
---


# Kubernetes Metrics

These metrics require installing some of followings:
- [Node Exporter](https://github.com/prometheus/node_exporter)
- [kube-state-metrics](https://github.com/kubernetes/kube-state-metrics/blob/main/docs/pod-metrics.md)
- [cAdvisor](https://github.com/google/cadvisor/blob/master/docs/storage/prometheus.md)

## Node metrics

- CPU utilization per node: `1 - (avg by (instance)(rate(node_cpu_seconds_total{mode="idle"}[5m])))`
- Memory utilization per node: `1 - (node_memory_MemAvailable_bytes / node_memory_MemTotal_bytes)`
- Disk utilization per node: `1 - (node_filesystem_avail_bytes / node_filesystem_size_bytes)`

- Number of pods with a certain phases on a node (from [this comment](https://github.com/kubernetes/kube-state-metrics/issues/332#issuecomment-355756863)):

    ```
    sum by(node)(kube_pod_info{} * on(pod, namespace) group_right(node) kube_pod_status_phase{phase="$phase"})
    ```

## Pod metrics

- CPU utilization per container: `sum by (container)(rate(container_cpu_usage_seconds_total{}[5m]))`
- CPU usages against request:

    ```
    sum by (namespace, container)(rate(container_cpu_usage_seconds_total{}[5m]))
    /
    sum by (namespace, container)(kube_pod_container_resource_requests{resource="cpu", unit="core"})
    ```

- CPU throttling:

    ```
    sum by (namespace, container)(rate(container_cpu_cfs_throttled_periods_total{}[5m]))
    /
    sum by (namespace, container)(rate(container_cpu_cfs_periods_total{}[5m]))
    ```

- CPU requests per namespace: `sum by (exported_namespace)(kube_pod_container_resource_requests{resource="cpu", unit="core"})`
- Memory utilization per container:
    - Max: `max by (namespace, container)(container_memory_working_set_bytes{})`
    - Median: `quantile by (namespace, container)(0.5, container_memory_working_set_bytes{})`
    - Min: `min by (namespace, container)(container_memory_working_set_bytes{})`
- Memory requests per namespace: `sum by (exported_namespace)(kube_pod_container_resource_requests{resource="memory", unit="byte"})`

## Persistent volumes

- The usage

    ```
    sum by (persistentvolumeclaim)(kubelet_volume_stats_used_bytes)
    /
    sum by (persistentvolumeclaim)(kubelet_volume_stats_capacity_bytes)
    ```


### Reference
Following some articles including followings
- [Stackoverflow: CPU usage for each node in prometheus](https://stackoverflow.com/a/66263640)
- [Prometheus cheetsheet](https://blog.ruanbekker.com/cheatsheets/prometheus/)
- [Average memory usage query prometheus](https://stackoverflow.com/questions/48835035/average-memory-usage-query-prometheus)
- [GitHub gist: max-rocket-internet/prom-k8s-request-limits.md](https://gist.github.com/max-rocket-internet/6a05ee757b6587668a1de8a5c177728b#queries-to-show-memory-and-cpu-as-percentage-of-both-request-and-limit)
