---
date: "2022-12-26T19:00:00Z"
tags:
- containerd
title: Getting Started Containerd
---

# Overview

containerd is designed to built with a larger system like kubernetes.
To see the overview, it's better to check something like

<iframe width="560" height="315" src="https://www.youtube.com/embed/q0xt_JrJiIg" title="YouTube video player" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture" allowfullscreen></iframe>

As the above video shows, it depends on plugins or runtime outside of its core by gRPC or the light weight gRPC (ttRPC), like
- [`runc`](https://github.com/opencontainers/runc) for container runtime
    - runc is a tool to manage containers on Linux following the OCI specification



# Install

Following [the official document](https://github.com/containerd/containerd/blob/main/docs/getting-started.md), install a containerd from the official binaries.

```bash
curl -sLO https://github.com/containerd/containerd/releases/download/v1.6.14/containerd-1.6.14-linux-amd64.tar.gz
sudo tar Cxzvf /usr/local containerd-1.6.14-linux-amd64.tar.gz
```

Setup the systemd for the containerd
```bash
curl -sLO https://raw.githubusercontent.com/containerd/containerd/main/containerd.service
sudo mkdir -p /usr/local/lib/systemd/system/
sudo mv containerd.service /usr/local/lib/systemd/system/
sudo systemctl daemon-reload
sudo systemctl enable --now containerd
```

Install runc
```bash
curl -sLO https://github.com/opencontainers/runc/releases/download/v1.1.4/runc.amd64
sudo install -m 755 runc.amd64 /usr/local/sbin/runc
```

Install cni plugins
```bash
curl -sLO https://github.com/containernetworking/plugins/releases/download/v1.1.1/cni-plugins-linux-amd64-v1.1.1.tgz
sudo mkdir -p /opt/cni/bin
sudo tar Cxzvf /opt/cni/bin cni-plugins-linux-amd64-v1.1.1.tgz
```

# Run a container by CLI

## ctr CLI
Use a `ctr` CLI to confirm followings:

* Pull an image by `ctr images pull [image ref]`
* Confirm the list of local images by `ctr images list`
* Run a container by `ctr run [image ref] [container name]`
* See namespaces by `ctr namespaces list`. Output result is like

    ```bash
    sudo ctr namespaces list
    NAME    LABELS
    default
    k8s.io
    ```

* See containers or images on the namespace, add an option `--namespace [namespace]`

There are other subcommands for
* plugins
* tasks
* snapshots
* events
* content
* leases

## nerdctr
[nerdctl](https://github.com/containerd/nerdctl) is a CLI for user friendly CLI compatible with Docker.

To install the version 1.1 of it,
```
curl -sLO https://github.com/containerd/nerdctl/releases/download/v1.1.0/nerdctl-1.1.0-linux-amd64.tar.gz
tar xzvf nerdctl-1.1.0-linux-amd64.tar.gz
sudo mv nerdctl /usr/local/bin
sudo nerdctl --version
```

* See the list of running containers: `nerdctl container list`

For example, to show the running containers with some information like its ip address, run following command

```bash
nerdctl --namespace k8s.io container list | awk '{ print $1 }' | tail -n +2 | xargs nerdctl --namespace k8s.io container inspect | jq -r '.[] | [.Id,.Path,.Image,.NetworkSettings.IPAddress] | @tsv'
```


# Network configuration

There is an [example configuration](https://kubernetes.io/docs/tasks/administer-cluster/migrating-from-dockershim/troubleshooting-cni-plugin-related-errors/#an-example-containerd-configuration-file) of CNI for k8s.

```json
sudo su -
cat << EOF | tee /etc/cni/net.d/10-containerd-net.conflist
{
 "cniVersion": "1.0.0",
 "name": "containerd-net",
 "plugins": [
   {
     "type": "bridge",
     "bridge": "cni0",
     "isGateway": true,
     "ipMasq": true,
     "promiscMode": true,
     "ipam": {
       "type": "host-local",
       "ranges": [
         [{
           "subnet": "10.88.0.0/16"
         }]
       ],
       "routes": [
         { "dst": "0.0.0.0/0" },
         { "dst": "::/0" }
       ]
     }
   },
   {
     "type": "portmap",
     "capabilities": {"portMappings": true},
     "externalSetMarkChain": "KUBE-MARK-MASQ"
   }
 ]
}
EOF
```

After adding the above configuration, it can be checked on a `nerdctl network ls` command.

```bash
$ nerdctl network ls
NETWORK ID    NAME              FILE
              containerd-net    /etc/cni/net.d/10-containerd-net.conflist
              host
              none

$ nerdctl network inspect containerd-net
nerdctl network inspect containerd-net
[
    {
        "Name": "containerd-net",
        "IPAM": {
            "Config": [
                {
                    "Subnet": "<IP ADDRESS>"
                }
            ]
        },
        "Labels": null
    }
]
```

I couldn't confirm how to confirm port mapping was configured correctly.

The configuration of the CNI is described in these documents:
- [CNI documentation](https://www.cni.dev/docs/spec/#section-1-network-configuration-format)
- [CNI plugin README](https://github.com/containernetworking/plugins#plugins-supplied)

## Start a container

```
ctr images pull docker.io/library/redis :alpine
ctr run docker.io/library/redis:alpine redis
```

With `nerdctl`, we can run a container like
```
nerdctl run -d docker.io/library/redis:alpine redis
```

To deploy a container into a network,
```
nerdctl run -d --network containerd-net --ip 10.128.128.16 docker.io/library/redis:alpine redis
```

But I got errors
```
$ sudo nerdctl run -d --network containerd-net docker.io/library/redis:alpine redis
FATA[0000] failed to create shim task: OCI runtime create failed: runc create failed: unable to start container process: error during container init: error running hook #0: error running hook: exit status 1, stdout: , stderr: time="2023-01-02T05:39:33Z" level=fatal msg="failed to call cni.Setup: plugin type=\"bridge\" failed (add): failed to set bridge addr: \"cni0\" already has an IP address different from 10.128.128.1/24"
Failed to write to log, write /var/lib/nerdctl/1935db59/containers/default/8dd436c2b4bfe5f3312aa872e00f61c6253801f25fcd1dfb313858fd625cdc57/oci-hook.createRuntime.log: file already closed: unknown


$ sudo nerdctl run -d --net containerd-net docker.io/library/redis:alpine redis
FATA[0000] failed to create shim task: OCI runtime create failed: runc create failed: unable to start container process: error during container init: error running hook #0: error running hook: exit status 1, stdout: , stderr: time="2023-01-02T05:39:43Z" level=fatal msg="failed to call cni.Setup: plugin type=\"bridge\" failed (add): failed to set bridge addr: \"cni0\" already has an IP address different from 10.128.128.1/24"
Failed to write to log, write /var/lib/nerdctl/1935db59/containers/default/8c454b3e839413472a3bdde5351b3772eba644ceb4adfcb6e8871d3837e941e2/oci-hook.createRuntime.log: file already closed: unknown

$ sudo nerdctl run -d --network containerd-net --ip 10.128.128.16 docker.io/library/redis:alpine redis
FATA[0000] failed to create shim task: OCI runtime create failed: runc create failed: unable to start container process: error during container init: error running hook #0: error running hook: exit status 1, stdout: , stderr: time="2023-01-02T05:35:48Z" level=fatal msg="failed to call cni.Setup: plugin type=\"bridge\" failed (add): failed to set bridge addr: \"cni0\" already has an IP address different from 10.128.128.1/24"
Failed to write to log, write /var/lib/nerdctl/1935db59/containers/default/77211e02cb8fed8eba40e3922cae75341da2018c219bbc18d5a292561e36fa88/oci-hook.createRuntime.log: file already closed: unknown
plugin type=\"bridge\" failed (add): failed to set bridge addr: \"cni0\" already has an IP address different from 10.128.128.1/24
```

The `cni0` is assigned to NIC.
```
ip addr
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
    inet6 ::1/128 scope host
       valid_lft forever preferred_lft forever
2: ens4: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1460 qdisc mq state UP group default qlen 1000
    inet 10.128.15.225/32 brd 10.128.15.225 scope global dynamic ens4
       valid_lft 65941sec preferred_lft 65941sec
3: cni0: <NO-CARRIER,BROADCAST,MULTICAST,PROMISC,UP> mtu 1500 qdisc noqueue state DOWN group default qlen 1000
    link/ether 5e:e7:6c:72:30:eb brd ff:ff:ff:ff:ff:ff
    inet 10.128.0.1/20 brd 10.128.15.255 scope global cni0
       valid_lft forever preferred_lft forever
```
