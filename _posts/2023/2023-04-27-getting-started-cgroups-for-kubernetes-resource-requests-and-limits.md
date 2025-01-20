---
title: Getting started cgroups for kubernetes resource requests and limits
date: 2023-04-27
tags:
  - kubernetes
  - cgroup
---

# Overview

There are many excellent articles or videos describing what is Kubernetes CPU resources requests and limits and how they are implemented.
Watch followings to understand how the CPU resource requests and limits of a container on the Pod spec works on a Linux host with cgroup.

- Kaslin and Kohei's video in CNCF in 2021
    - [A YouTube video](https://youtu.be/WB3_sV_EQrQ?t=788)
    - [The slides of this video](https://speakerdeck.com/inductor/resource-requests-and-limits-under-the-hood-the-journey-of-a-pod-spec?slide=34)
- Shon's articles
    - [CPU resource requests](https://medium.com/directeam/kubernetes-resources-under-the-hood-part-2-6eeb50197c44)
    - [CPU resource limits](https://medium.com/directeam/kubernetes-resources-under-the-hood-part-3-6ee7d6015965)

{% comment %}
https://blog.kintone.io/entry/2022/03/08/170206
https://medium.com/omio-engineering/cpu-limits-and-aggressive-throttling-in-kubernetes-c5b20bd8a718
https://engineering.squarespace.com/blog/2017/understanding-linux-container-scheduling
{% endcomment %}

In short about CPU resources,

* K8s CPU resources are for CPU time, and not CPU cores.
* At a moment, a k8s pod might use all of CPU cores, beyond one set by CPU requests.
* If CPU limit is set
    * Even if there is idle CPUs, they are not used if CPU usages hit the CPU limits
* k8s CPU resources are mapped on cgroups
    * requests: `cpu.shares`
    * limits: `cpu.cfs_quota_us` with `cpu.cfs_period_us`

In this article, I'm going to check these cgroups configurations


# Cgroups v1

Cgroups is to allow processes to have CPUs, memories, and network bandwidth.

## Commands to work around Cgroups

1. Check if my kernel supports it by confirming `/proc/cgroups`. It's supported if CPU or others are enabled there.
1. Install required tools by `sudo apt install cgroup-tools`.
1. Create a cgroup by `sudo cgcreate -g cpu,memory:/my_cgroup`
    1. We can confirm the cgroup under `/sys/fs/cgroup/{cpu,memory}/my_cgroup`
1. We can set some resource parameters there
    - `sudo cgset -r cpu.cfs_quota_us=50000 my_cgroup` to limit the CPU usages
        ```
        > cat /sys/fs/cgroup/cpu/my_group/cpu.cfs_quota_us
        50000
        ```

    - `sudo cgset -r memory.limit_in_bytes=1G my_cgroup` to limit the memory usages
        ```
        > cat /sys/fs/cgroup/memory/my_group/memory.limit_in_bytes
        1073741824
        ```
1. Add processes to the cgroup by adding some IDs on the file under the cgroup directory
    - tasks: PIDs
    - cgroup.procs: TGIDs (Thread Group IDs)
1. Monitor usages
    - CPU: `/sys/fs/cgroup/memory/my_group/cpu.stat`. It's described in the [page](https://access.redhat.com/documentation/en-us/red_hat_enterprise_linux/6/html/resource_management_guide/sec-cpu), for example
       - nr_periods: number of period intervals
       - nr_throttled: number of times tasks in a cgroup have been throttled
       - throttled_time: the total time duration
    - Memory: `/sys/fs/cgroup/memory/my_cgroup/memory.usage_in_bytes`
1. Delete the cgroup by `sudo cgdelete -g cpu,memory:/my_cgroup`

## Testings

There are good articles to test these configurations, including
* [awinecki's article](https://andywine.dev/constrain-process-cpu-usage-with-cgroups/)

I referred the above articles to test something further.

### Load test CPUs
At first, prepare 2 cgroups, fast and slow.
Install [`stress`](https://linux.die.net/man/1/stress) beforehand.

* Create 2 cgroups
    * `sudo cgcreate -g cpu:/fast`
    * `sudo cgcreate -g cpu:/slow`
* Set cpu.shares to assign 3:1 ratios between fast and slow
    * `sudo cgset -r cpu.shares=750 fast`
    * `sudo cgset -r cpu.shares=250 slow`

Then load test CPUs on each cgroups for how it works.

1. Run a stress command with [cgexec](https://linux.die.net/man/1/cgexec) on the slow cgroup by next command. Note that 12 is the number of CPU cores on my machine.

    ```zsh
    > sudo cgexec -g cpu:slow stress --cpu 12
    stress info: [1065] dispatching hogs: 12 cpu 0 io, 0 vm, 0 hdd
    ```

1. See CPU usages by `htop`. It's confirmed that 12 CPUs are used for 100%.

    ![htop after only slow cgroup](/assets/images/posts/2023/04/27/getting_started_cgroups_for_kubernetes_resource_requests_and_limits/02_01_htop_after_slow_runs.png)

1. Run a stress on the fast cgroup by

    ```zsh
    > sudo cgexec -g cpu:fast stress --cpu 12
    stress: info: [1305] dispatching hogs: 12 cpu, 0 io, 0 vm, 0 hdd
    ```

1. See CPU usages by `htop`.
   The PID output by `stress` can be seen on the column PPID on htop to recognize which one is for fast cgroup and which one is for slow cgroup.
   And also, the fast cgroup uses about 75% (9 CPUs) while the slow cgroup uses 25% (3 CPUs), which is similar to the configuration of `cpu.shares` ratios.

    ![htop after 2 cgroups](/assets/images/posts/2023/04/27/getting_started_cgroups_for_kubernetes_resource_requests_and_limits/02_02_htop_after_fast_runs.png)

1. Confirm `cpu.cfs_quota_us` and `cpu.cfs_period_us` on the fast cgroup.

    ```zsh
    > sudo cgget -r cpu.cfs_quota_us fast
    fast:
    cpu.cfs_quota_us: -1

    > sudo cgget -r cpu.cfs_period_us fast
    fast:
    cpu.cfs_period_us: 100000
    ```

1. Set `cpu.cfs_quota_us to 400000` to the fast cgroup so 4 CPUs are used per a second.

    ```zsh
    > sudo cgset -r cpu.cfs_quota_us=400000 fast
    >
    ```

1. See htop. CPU usages on the stress against the fast cgroup which PID=1305, was reduced and those for the slow cgroup was increased.

    ![htop after setting cfs_quota_us](/assets/images/posts/2023/04/27/getting_started_cgroups_for_kubernetes_resource_requests_and_limits/02_03_htop_after_fast_set_cgroup_cfs_quota_us.png)

1. Stop the stress against the slow cgroup, which is without cpu.cfs_quota_us.
1. See htop. CPU usages on the stress against the fast cgroup cannot use 100% CPU usages. There are 12 CPUs and they are used about 30%, so 4 CPUs are roughly used in a second.

    ![htop after stopping slow cgroup](/assets/images/posts/2023/04/27/getting_started_cgroups_for_kubernetes_resource_requests_and_limits/02_04_htop_after_stopped_slow_cgroup.png)


## Cgroups for Memory

* Memory requests: It seems this is not used on Cgroup v1. I couldn't find any document to use this.
    * On Cgroup v2, it seems it sets `memory.min` and `memory.low`
* Memory limit: sets `memory.limit_in_bytes`


# Questions

* How is it related to [QoS](https://kubernetes.io/docs/tasks/configure-pod-container/quality-service-pod/)?
    * [QoS](https://kubernetes.io/docs/concepts/workloads/pods/pod-qos/) is used to decide priorities to evict pods in the case of [NodePressure](https://kubernetes.io/docs/concepts/scheduling-eviction/node-pressure-eviction/)
