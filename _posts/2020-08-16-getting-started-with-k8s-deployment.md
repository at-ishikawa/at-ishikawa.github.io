---
date: 2020-08-16
title: k8s Deployment
tags:
  - k8s
  - deployment
---

The deployment is many use cases and in this page, they're not described.
For the details for those use cases or the concept of deployment, see [official page](https://kubernetes.io/docs/concepts/workloads/controllers/deployment/).

kubectl commands
===

Revisions and rollbacks
---
1. How to get revision history of deployments.
  - To see revision history: `kubectl rollout history deployment.v1.apps/[deployment name]`. For example:
    ```
    kubectl rollout history deployment.v1.apps/test-deployment
    deployment.apps/test-deployment
    REVISION  CHANGE-CAUSE
    135       <none>
    137       <none>
    138       <none>
    139       <none>
    140       <none>
    141       <none>
    142       <none>
    143       <none>
    144       <none>
    145       <none>
    146       <none>
    ```
  - To see the detail of a revision: `kubectl rollout history deployment.v1.apps/[deployment name] --revision [revision number]`
    ```
    > kubectl rollout history deployment.apps/test-deployment --revision 145
    deployment.apps/test-deployment with revision #145
    Pod Template:
      Labels:       app=test-deployment
            env=prod
            pod-template-hash=c7d84d6fc
      Annotations:  prometheus.io/port: 9100
            prometheus.io/scrape: true
      Containers:
       test-deployment:
        Image:      gcr.io/test-project/test-deployment:tag
        Port:       <none>
        Host Port:  <none>
        Limits:
          cpu:      1
          memory:   256Mi
        Requests:
          cpu:      1
          memory:   256Mi
        Environment:
          DEBUG:    false
        Mounts:     <none>
      Volumes:      <none>
    ```
1. How to rollback deployment: `kubectl rollout undo deployment [deployment name] (--to-revision [revision number])`. Examples:
  - Rollback to the previous version.
    ```
    > kubectl rollout undo deployment test-deployment
    deployment.extensions/test-deployment rolled back
	```
  - Rollback to a specific version.
    ```
    > kubectl rollout undo deployment test-deployment --to-revision 2
	```
