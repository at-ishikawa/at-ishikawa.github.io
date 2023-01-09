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

## Online DDL

There are a few algorithms described in the next video.

<iframe width="560" height="315" src="https://www.youtube-nocookie.com/embed/IiiFVr6UEIc?start=1871" title="YouTube video player" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share" allowfullscreen></iframe>

Also, there are a few documents related to online DDL
- The Online DDL design is descrbed in [this design document](https://github.com/pingcap/tidb/blob/master/docs/design/2018-10-08-online-DDL.md)
- Google F1 algorithm, which is refered from TiDB online schema change, is described in this [research paper](https://static.googleusercontent.com/media/research.google.com/en//pubs/archive/41376.pdf)



# Walking through by the example data

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

### Manage TiKV regions

* Show stores (TiKV nodes): `store`
* Confirm regions: `region | jq .`
* Confirm a specific region by its key: `region key 7480000000000000FF5E00000000000000F8`
* Confirm the hot spots by `hot (read|write)`
* Confirm an abnormal region: `region check (miss-peer|extra-peer|down-peer|pending-peer)`


## Manage regions

### Confirm regions on a table

To confirm regions via TiDB:

    ```sql
    show table [table] regions;
    show table [table] index [index] regions;
    ```

The examples to look around regions:
```sql
mysql> show table salaries regions;
+-----------+----------------------------------------------+----------------------------------------------+-----------+-----------------+------------------+------------+---------------+------------+----------------------+------------------+------------------------+------------------+
| REGION_ID | START_KEY                                    | END_KEY                                      | LEADER_ID | LEADER_STORE_ID | PEERS            | SCATTERING | WRITTEN_BYTES | READ_BYTES | APPROXIMATE_SIZE(MB) | APPROXIMATE_KEYS | SCHEDULING_CONSTRAINTS | SCHEDULING_STATE |
+-----------+----------------------------------------------+----------------------------------------------+-----------+-----------------+------------------+------------+---------------+------------+----------------------+------------------+------------------------+------------------+
|      5161 | t_113_                                       | t_113_r_03800000000003112104194fb40000000000 |      5164 |            5004 | 5162, 5163, 5164 |          0 |             0 |          0 |                   77 |           988924 |                        |                  |
|      5165 | t_113_r_03800000000003112104194fb40000000000 | t_113_r_0380000000000623c70419599e0000000000 |      5167 |            5003 | 5166, 5167, 5168 |          0 |            39 |          0 |                   72 |           942081 |                        |                  |
|      5149 | t_113_r_0380000000000623c70419599e0000000000 | t_115_                                       |      5151 |            5003 | 5150, 5151, 5152 |          0 |            42 |          0 |                   72 |           941966 |                        |                  |
+-----------+----------------------------------------------+----------------------------------------------+-----------+-----------------+------------------+------------+---------------+------------+----------------------+------------------+------------------------+------------------+
3 rows in set (0.04 sec)


-- Check regions for a secondary index
mysql> create index salaries_to_date on salaries(to_date);
mysql> show table salaries index salaries_to_date regions;
+-----------+--------------------------------------------------------------------+--------------------------------------------------------------------+-----------+-----------------+------------------+------------+---------------+------------+----------------------+------------------+------------------------+------------------+
| REGION_ID | START_KEY                                                          | END_KEY                                                            | LEADER_ID | LEADER_STORE_ID | PEERS            | SCATTERING | WRITTEN_BYTES | READ_BYTES | APPROXIMATE_SIZE(MB) | APPROXIMATE_KEYS | SCHEDULING_CONSTRAINTS | SCHEDULING_STATE |
+-----------+--------------------------------------------------------------------+--------------------------------------------------------------------+-----------+-----------------+------------------+------------+---------------+------------+----------------------+------------------+------------------------+------------------+
|      6001 | t_113_                                                             | t_113_i_2_041934020000000000038000000000005185041933e20000000000   |      6004 |            5004 | 6002, 6003, 6004 |          0 |            39 |          0 |                   25 |           328620 |                        |                  |
|      6021 | t_113_i_2_041934020000000000038000000000005185041933e20000000000   | t_113_i_2_04195b0c0000000000038000000000037e2e041957cc0000000000   |      6024 |            5004 | 6022, 6023, 6024 |          0 |            39 |          0 |                   52 |                0 |                        |                  |
|      6009 | t_113_i_2_04195b0c0000000000038000000000037e2e041957cc0000000000   | t_113_i_2_04195b7400000000000380000000000652cb041958340000000000   |      6010 |               1 | 6010, 6011, 6012 |          0 |           427 |          0 |                    2 |            30720 |                        |                  |
|      6017 | t_113_i_2_04195b7400000000000380000000000652cb041958340000000000   | t_113_i_2_04196c5a000000000003800000000001677d0419691a0000000000   |      6020 |            5004 | 6018, 6019, 6020 |          0 |            39 |          0 |                   52 |                0 |                        |                  |
|      6013 | t_113_i_2_04196c5a000000000003800000000001677d0419691a0000000000   | t_113_i_2_047ef1020000000000038000000000007c3e04196bd80000000000   |      6015 |            5003 | 6014, 6015, 6016 |          0 |           427 |          0 |                    2 |            30720 |                        |                  |
|      6005 | t_113_i_2_047ef1020000000000038000000000007c3e04196bd80000000000   | t_113_i_2_047ef102000000000003800000000007a11f04196a3a000000000000 |      6008 |            5004 | 6006, 6007, 6008 |          0 |           613 |          0 |                   18 |           222607 |                        |                  |
|      5161 | t_113_i_2_047ef102000000000003800000000007a11f04196a3a000000000000 | t_113_r_03800000000003112104194fb40000000000                       |      5164 |            5004 | 5162, 5163, 5164 |          0 |           517 |          0 |                   76 |           985862 |                        |                  |
+-----------+--------------------------------------------------------------------+--------------------------------------------------------------------+-----------+-----------------+------------------+------------+---------------+------------+----------------------+------------------+------------------------+------------------+
7 rows in set (0.07 sec)

--- Check region peers and if it's leader
mysql> select * from information_schema.tikv_region_peers order by region_id;
+-----------+---------+----------+------------+-----------+--------+--------------+
| REGION_ID | PEER_ID | STORE_ID | IS_LEARNER | IS_LEADER | STATUS | DOWN_SECONDS |
+-----------+---------+----------+------------+-----------+--------+--------------+
|         2 |       3 |        1 |          0 |         0 | NORMAL |         NULL |
|         2 |    5048 |     5003 |          0 |         1 | NORMAL |         NULL |
|         2 |    5092 |     5004 |          0 |         0 | NORMAL |         NULL |
|        10 |    5088 |     5004 |          0 |         0 | NORMAL |         NULL |
|        10 |      11 |        1 |          0 |         1 | NORMAL |         NULL |
|        10 |    5044 |     5003 |          0 |         0 | NORMAL |         NULL |
|      5129 |    5132 |     5004 |          0 |         0 | NORMAL |         NULL |
|      5129 |    5130 |        1 |          0 |         0 | NORMAL |         NULL |
|      5129 |    5131 |     5003 |          0 |         1 | NORMAL |         NULL |
|      5141 |    5143 |     5003 |          0 |         1 | NORMAL |         NULL |
|      5141 |    5144 |     5004 |          0 |         0 | NORMAL |         NULL |
|      5141 |    5142 |        1 |          0 |         0 | NORMAL |         NULL |
...
```

### Split regions
There is a query to split by `SPLIT TABLE table_name [INDEX index_name] BETWEEN (lower_value) AND (upper_value) REGIONS region_num`

### Add new node and distribute regions

Following [this document](https://docs.pingcap.com/tidb/v3.0/horizontal-scale#add-a-node-dynamically-1), just adding nodes dynamically.
Then regions are automatically distributed by PD.

```bash
kubectl patch -n tidb-cluster tc basic --type merge --patch '{"spec":{"tikv":{"replicas":5}}}'
kubectl get pods -w
```

Then run `tiup ctl:v6.5.0 pd  -u http://127.0.0.1:2379 store` a few times and see the number of region and leader counts would increase on the new node.

### Confirm multiple regions

Follow
- [Placement rules](https://docs.pingcap.com/tidb/stable/configure-placement-rules)
- [Follower read](https://docs.pingcap.com/tidb/dev/follower-read)

```bash
tiup ctl:v6.5.0 pd  -u http://127.0.0.1:2379 config set replication.max-replicas 5
```

On TiDB, update tidb_replica_read to `follower` to enable a follower read

```sql
set @@tidb_replica_read = 'follower';
```

TODO: But still all replicas show that role name is Voter.

```bash
tiup ctl:v6.5.0 pd  -u http://127.0.0.1:2379 region | less
```


## Identify and analyze slow queries

See [this document](https://docs.pingcap.com/tidb/stable/identify-slow-queries#identify-slow-queries)

* Run queries on `information_schema.SLOW LOG`
* Run the query `ADMIN SHOW SLOW recent 10`
