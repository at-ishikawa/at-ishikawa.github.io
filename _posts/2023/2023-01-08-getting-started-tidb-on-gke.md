---
title: Getting Started TiDB on GKE
date: 2023-01-08T19:00:00
tags:
  - tidb
  - mysql
  - kubernetes
  - gke
---


See [another post](/2022/05/01/getting-started-with-tidb-by-kubernetes-operator/) also to set up a TiDB on minikube.

To get the basic concept and architecture of TiDB, there are many articles and videos on the Internet, like following:

<iframe width="560" height="315" src="https://www.youtube.com/embed/_0rGOo8fNbQ" title="YouTube video player" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share" allowfullscreen></iframe>


# Set up a TiDB cluster by Kubernetes Operator
Follow [the official document for GKE](https://docs.pingcap.com/tidb-in-kubernetes/stable/deploy-on-gcp-gke) and [the official document for kubernetes operator](https://docs.pingcap.com/tidb-in-kubernetes/stable/get-started#step-2-deploy-tidb-operator) for this document.


## ext4 file system configuration

It's recommended to configure following options for better I/O performance for ext4 file system.

- nodelalloc: Disable delayed allocation. See [this option](https://www.phoronix.com/news/EXT4-No-Delalloc-Perf-Fix) for why it's enabled in the first place.
- noatime: Disable to store access time on files. See [this article](https://opensource.com/article/20/6/linux-noatime) for the benefit for this performance improvement.

The options of ext4 file system can be seen in [this man document](https://www.kernel.org/doc/Documentation/filesystems/ext4.txt), though there is no noatime option.


## Install TiDB CRD and its operator

First, create the namespace at first.
```bash
kubectl create namespace tidb-cluster
```

Then following [this document](https://docs.pingcap.com/tidb-in-kubernetes/stable/get-started#step-2-deploy-tidb-operator), installing kubernetes operator.

```bash
kubectl create -f https://raw.githubusercontent.com/pingcap/tidb-operator/v1.4.0/manifests/crd.yaml
helm repo add pingcap https://charts.pingcap.org/
kubectl create namespace tidb-admin
helm install --namespace tidb-admin tidb-operator pingcap/tidb-operator --version v1.4.0
```

Confirm an installation on the operator
```bash
kubectl get pods --namespace tidb-admin -l app.kubernetes.io/instance=tidb-operator
```

## Create a cluster

Create followings
- the example of a cluster.
- tidb dashboard (for some reasons, it doesn't work)
- tidb monitoring like prometheus and Grafana

```bash
kubectl -n tidb-cluster apply -f
https://raw.githubusercontent.com/pingcap/tidb-operator/v1.4.0/examples/basic/tidb-cluster
kubectl -n tidb-cluster apply -f https://raw.githubusercontent.com/pingcap/tidb-operator/v1.4.0/examples/basic/tidb-dashboard.yaml
kubectl -n tidb-cluster apply -f https://raw.githubusercontent.com/pingcap/tidb-operator/v1.4.0/examples/basic/tidb-monitor.yaml
```

Confirm the setup
```bash
watch kubectl get po -n tidb-cluster
```

## Connect to TiDB

```bash
kubectl get svc -n tidb-cluster
kubectl port-forward -n tidb-cluster svc/basic-tidb 14000:4000 > pf14000.out &
mysql --comments -h 127.0.0.1 -P 14000 -u root mysql
```

Then create a table
```sql
use test;
create table hello_world (id int unsigned not null auto_increment primary key, v varchar(32));
select * from information_schema.tikv_region_status where db_name=database() and table_name='hello_world'\G
```

See the TiDB cluster information
```sql
select * from information_schema.cluster_info\G
```

## Check Grafana

```
kubectl port-forward -n tidb-cluster svc/basic-grafana 3000 > pf3000.out &
```

Then go to [http://localhost:3000][] to see Grafana.

## Check a TiDB Dashboard

I needed to update tidb-operator version to `v1.4.2` because it was v1.3.5 and TiDBDashboard wasn't created.

```bash
helm upgrade --namespace tidb-admin tidb-operator pingcap/tidb-operator --version v1.4.2 --set operatorImage='pingcap/tidb-operator:v1.4.2'
```

Once you make sure your TiDB Operator runs on v1.4.2, then run port-forward and access the TiDB dashboard

```bash
kubectl port-forward -n tidb-cluster svc/basic-tidb-dashboard-exposed 12333 > pf12333.out &
```

Access `localhost:12333` on a browser.

If you want to use TopSQL and Continuous Profiling features, then deploy TidbNGMonitoring described in [this document](https://docs.pingcap.com/tidb-in-kubernetes/stable/access-dashboard#enable-continuous-profiling)

# Basic operations

## Update the number of TiKV, TiDB, and PD nodes:

```bash
kubectl patch -n tidb-cluster tc basic --type merge --patch '{"spec":{"tikv":{"replicas":3}}}'
kubectl patch -n tidb-cluster tc basic --type merge --patch '{"spec":{"tidb":{"replicas":2}}}'
kubectl patch -n tidb-cluster tc basic --type merge --patch '{"spec":{"pd":{"replicas":2}}}'
```

## Increase the amount of storage sizes of Network PV

Follow [this document](https://docs.pingcap.com/tidb-in-kubernetes/stable/configure-storage-class#network-pv-configuration).

```bash
kubectl patch -n tidb-cluster tc basic --type merge --patch '{"spec":{"tikv":{"requests":{"storage":"100Gi"}}}}'
```

Note that `Shrinking persistent volumes is not supported` for Kubernetes Persistent Volumes.


# Load an example data

At first, we'll load the data from [this example](https://dev.mysql.com/doc/employee/en/).

Before loading it, at first, check a storage engine on TiDB:
```
mysql> select * from information_schema.engines;
+--------+---------+------------------------------------------------------------+--------------+------+------------+
| ENGINE | SUPPORT | COMMENT                                                    | TRANSACTIONS | XA   | SAVEPOINTS |
+--------+---------+------------------------------------------------------------+--------------+------+------------+
| InnoDB | DEFAULT | Supports transactions, row-level locking, and foreign keys | YES          | YES  | YES        |
+--------+---------+------------------------------------------------------------+--------------+------+------------+
1 row in set (0.04 sec)
```

Then expand storage sizes of TiKV nodes

```bash
kubectl patch -n tidb-cluster tc basic --type merge --patch '{"spec":{"tikv":{"requests":{"storage":"100Gi"}}}}'
```

Now it's time to load the data by

```bash
tar xzvf test_db-1.0.7.tar.gz
cd test_db/
mysql --comments -h 127.0.0.1 -P 14000 -u root < employees.sql
```

Validate if the data is loaded correctly

```bash
mysql --comments -h 127.0.0.1 -P 14000 -u root employees < test_employees_md5.sql
```

In order to manage some components directly, install tiup and a few CLIs

```bash
curl --proto '=https' --tlsv1.2 -sSf https://tiup-mirrors.pingcap.com/install.sh | sh
tiup ctl:v6.5.0 pd
```

## Control PD metadata

Following [this document](https://docs.pingcap.com/tidb-in-kubernetes/stable/tidb-toolkit#use-pd-control-on-kubernetes),

```shell
kubectl port-forward -n tidb-cluster svc/basic-pd 2379:2379 &>/tmp/portforward-pd.log &
```

Following commands are one of pd-ctl, listed in [this document](https://docs.pingcap.com/tidb/dev/pd-control), like `tiup ctl:v6.5.0 pd -u http://127.0.0.1:2379 cluster`.

### Manage a cluster

* Confirm the cluster: `cluster`
* Confirm the configuration: `config show all`

### Manage PD members

* Confirm the PD members: `member`
* Confirm the leader: `member leader show`


## Identify and analyze slow queries

See [this document](https://docs.pingcap.com/tidb/stable/identify-slow-queries#identify-slow-queries)

* Run queries on `information_schema.SLOW LOG`
* Run the query `ADMIN SHOW SLOW recent 10`
