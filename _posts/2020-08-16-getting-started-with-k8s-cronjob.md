---
date: 2020-08-16
title: Getting Started with Kubernetes Cronjob
tags:
  - kubernetes
---

Configurations
===

Important fields
---
1. `.spec.jobTemplate.spec.template.spec.restartPolicy`: How to handle when a container in a pod. Note that this doesn't mean a job restarts or not (it's commented on [this github issue](https://github.com/kubernetes/kubernetes/issues/20255#issuecomment-310540940)).
  - `onFailure`: Restart a container in the same pod when it fails.
  - `Never`: Restart a container with new pod.
1. `.spec.concurrencyPolicy`: How to handle concurrent jobs
  - `Allow` (Default): Allows concurrent runs
  - `Forbid`: Does not allow concurrent runs
  - `Replace`: Replace old job with new job if previous job hasn't finished when it's time to schedule new job
1. `.spec.successfulJobsHistoryLimit`, `.spec.failedJobsHistoryLimit`: How many pods is kept after it succeeded or failed
1. `.spec.startingDeadlineSeconds`: How long jobs keep trying to start. Also, see a [pitfall](#startingdeadlineseconds) section.

Pitfalls
===

StartingDeadlineSeconds
---
A cronjob stops starting jobs if it "misses" jobs more than 100 times in a specific period related with `.spec.startingDeadlineSeconds`.
The details are explained in [here](https://kubernetes.io/docs/concepts/workloads/controllers/cron-jobs/#cron-job-limitations).

Here is the brief explanation, though they might be wrong or incomplete.
- If `startingDeadlineSeconds` isn't set, the number of missed jobs is counted since last scheduled time. And if it exceeds more than 100 times, then jobs won't start. If `startingDeadlineseconds` is set, then missed jobs counted in last `startingDeadlineSeconds`.
- The number of missed jobs includes when a new job cannot run because a `concurrencyPolicy` is `Forbid` and previous job is still running.


Other pitfalls
---
1. Jobs should be "idempotent", because the same jobs might be created more than once, or no job might be created.
1. CronJobs may run more than once even if `.spec.parallelism` = 1, `.spec.completions` = 1, and `.spec.template.spec.restartPolicy` = "Never"
  - See [official page](https://kubernetes.io/docs/concepts/workloads/controllers/job/#handling-pod-and-container-failures) for more details.
1. CronJobs usually fail to start more than 30 or 60 seconds (in 2018).
  - See [this slide (Japanese)](https://speakerdeck.com/potsbo/kubernetes-cronjob-implementation-in-detail-number-k8sjp?slide=3) for more details.
1. Workaround is required for cronjobs with istio sidecars. There is an option to update manifests of istio sidecar to stop the sidecar after corresponding job container stops, or disable injecting istio sidecar.
  - See [this GitHub issues](https://github.com/istio/istio/issues/11659#issuecomment-479547294) for workarounds to insert istio sidecar for cronjobs
  - See [this GitHub issues](https://github.com/kubernetes/enhancements/issues/753) to support a sidecar as a kubernetes 1st citizen. After this issue is solved, no workaround won't be required anymore.


kubectl commands
===
1. How to run a job from cronjob: `kubectl create job --from=cronjob/[cronjob name] [job name]`


References
====
1. [Kubernetes documentation: Running Automated Tasks with a CronJob](https://kubernetes.io/docs/tasks/job/automated-tasks-with-cron-jobs/)
1. [Kubernetes documentation: Jobs](https://kubernetes.io/docs/concepts/workloads/controllers/job/)
1. [How we learned to improve Kubernetes CronJobs at Scale (Part 1 of 2)](https://eng.lyft.com/improving-kubernetes-cronjobs-at-scale-part-1-cf1479df98d4)
1. [What does Kubernetes cronjob’s `startingDeadlineSeconds` exactly mean?](https://medium.com/@hengfeng/what-does-kubernetes-cronjobs-startingdeadlineseconds-exactly-mean-cc2117f9795f)
