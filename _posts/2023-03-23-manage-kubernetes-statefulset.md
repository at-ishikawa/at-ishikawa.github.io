---
title: Manage kubernetes StatefulSet
date: 2023-03-23
tags:
  - kubernetes
  - statefulset
---

# Operations

## Increase the storage sizes without downtime

We cannot update the resources.storageSize of the StatefulSet to increase the volume sizes.
So, we have to follow a few steps to update them.

See [this article](https://itnext.io/resizing-statefulset-persistent-volumes-with-zero-downtime-916ebc65b1d4) for more details.

1. Confirm if the storageclass of persistent volumes has `.allowVolumeExpansion=true`
    ````
    > kubectl get storageclass standard -o yaml | yq .allowVolumeExpansion
    true
    ```

1. Update all PVC to have more storage sizes
1. Delete the statefulset objects without deleting pods by `kubectl delete statefulset $statefulset_name --cascade=orphan`
1. Recreate the statefulset with the new storage size

## Decrease the storage size

Shrinking disk sizes is not supported. [official document](https://kubernetes.io/blog/2018/07/12/resizing-persistent-volumes-using-kubernetes/).
