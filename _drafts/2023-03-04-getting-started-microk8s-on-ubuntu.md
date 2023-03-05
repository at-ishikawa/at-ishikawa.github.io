---
title: Set up a microk8s locally on Ubuntu
date: 2023-03-04T00:00:00
draft: true
tags:
  - kubernetes
  - ubuntu
  - microk8s
---

This post is an attempt and fail to complete.

This post is written by following [this tutorial](https://ubuntu.com/tutorials/install-a-local-kubernetes-with-microk8s#1-overview).

# Getting Started

Install microk8s by snap.
At first, confirm the latest stable version.

```sh
snap info microk8s
```

After confirm the latest version, install the microk8s.

```sh
sudo snap install microk8s --classic --channel=1.26/stable
```

After installing it, run following commands
```bash
sudo usermod -a -G microk8s $USER
sudo chown -f -R $USER ~/.kube
```


Next, enable a few addons
```sh
microk8s enable dns dashboard storage
```

Next, confirm the progress of installing addons by `microk8s kubectl get all --all-namespaces`.
In my case, all pods were Pending statuses. And noticed there was an error

```sh
> microk8s kubectl get events -n kube-system
LAST SEEN   TYPE      REASON             OBJECT                                                MESSAGE
55s         Warning   FailedScheduling   pod/coredns-588fd544bf-kfmnq                          0/1 nodes are available: 1 node(s) had taint {node.kubernetes.io/not-ready: }, that the pod didn't tolerate.
55s         Warning   FailedScheduling   pod/dashboard-metrics-scraper-db65b9c6f-bm9rs         0/1 nodes are available: 1 node(s) had taint {node.kubernetes.io/not-ready: }, that the pod didn't tolerate.
55s         Warning   FailedScheduling   pod/heapster-v1.5.2-58fdbb6f4d-r655s                  0/1 nodes are available: 1 node(s) had taint {node.kubernetes.io/not-ready: }, that the pod didn't tolerate.
55s         Warning   FailedScheduling   pod/hostpath-provisioner-75fdc8fccd-hkrsb             0/1 nodes are available: 1 node(s) had taint {node.kubernetes.io/not-ready: }, that the pod didn't tolerate.
55s         Warning   FailedScheduling   pod/kubernetes-dashboard-67765b55f5-22qtp             0/1 nodes are available: 1 node(s) had taint {node.kubernetes.io/not-ready: }, that the pod didn't tolerate.
55s         Warning   FailedScheduling   pod/monitoring-influxdb-grafana-v4-6dc675bf8c-rtjq4   0/1 nodes are available: 1 node(s) had taint {node.kubernetes.io/not-ready: }, that the pod didn't tolerate.
```

Apparantly, a node is not ready for some reasons.

```bash
> microk8s kubectl get nodes
NAME              STATUS     ROLES    AGE     VERSION
ishikawa-ubuntu   NotReady   <none>   4h21m   v1.18.20

> microk8s kubectl describe nodes ishikawa-ubuntu | grep Warning
  Warning  InvalidDiskCapacity      2s     kubelet  invalid capacity 0 on image filesystem
```


## Troubleshootings

Mainly following a [document](https://microk8s.io/docs/troubleshooting) for troubleshooting.

### Trial1: Run microk8s inspect

```sh
> microk8s inspect
...
# Warning: iptables-legacy tables present, use iptables-legacy to see them
 WARNING:  IPtables FORWARD policy is DROP. Consider enabling traffic forwarding with: sudo iptables -P FORWARD ACCEPT
The change can be made persistent with: sudo apt-get install iptables-persistent
File "/etc/docker/daemon.json" does not exist.
You should create it and add the following lines:
{
    "insecure-registries" : ["localhost:32000"]
}
and then restart docker with: sudo systemctl restart docker
Building the report tarball
  Report tarball is at /var/snap/microk8s/2271/inspection-report-20230304_204000.tar.gz
```

So as the above message said
- Add the above json in `/etc/docker/daemon.json`
- Run the commands that warnings proposed

    ```sh
    sudo iptables -P FORWARD ACCEPT
    sudo apt install iptables-persistent
    ```

