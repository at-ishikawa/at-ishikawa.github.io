---
date: "2023-02-11T19:00:00Z"
tags:
- tidb
title: Manage TiKV regions on TiDB
---

TiDB data is split into multiple nodes and they're called the name as a region.

In this document, I just played around to confirm the places and behaviors for regions on TiDB.


## Manage TiKV regions through PD metadata

To handle PD metadata from CLI, follow [this document](https://docs.pingcap.com/tidb-in-kubernetes/stable/tidb-toolkit#use-pd-control-on-kubernetes).

At first, a port-forward to access PD pods

```shell
kubectl port-forward -n tidb-cluster svc/basic-pd 2379:2379 &>/tmp/portforward-pd.log &
```

Then run following sub commands to check each.
These commands are one of pd-ctl, listed in [this document](https://docs.pingcap.com/tidb/dev/pd-control), and to execute them, it has to add a command like `tiup ctl:v6.5.0 pd -u http://127.0.0.1:2379 store`:

* Show stores (TiKV nodes): `store`
* Confirm regions: `region`
* Confirm a specific region by its key: `region key 7480000000000000FF5E00000000000000F8`
* Confirm the hot spots by `hot (read|write)`
* Confirm an abnormal region: `region check (miss-peer|extra-peer|down-peer|pending-peer)`


## Confirm regions on a table

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

## Split regions
There is a query to split by `SPLIT TABLE table_name [INDEX index_name] BETWEEN (lower_value) AND (upper_value) REGIONS region_num`

## Add new node and distribute regions

Following [this document](https://docs.pingcap.com/tidb/v3.0/horizontal-scale#add-a-node-dynamically-1), just adding nodes dynamically.
Then regions are automatically distributed by PD.

```bash
kubectl patch -n tidb-cluster tc basic --type merge --patch '{"spec":{"tikv":{"replicas":5}}}'
kubectl get pods -w
```

Then run `tiup ctl:v6.5.0 pod  -u http://127.0.0.1:2379 store` a few times and see the number of region and leader counts would increase on the new node.

## Confirm multiple regions

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

**TODO: But still all replicas show that role name is Voter, and no replica shows a Follower.**

```bash
tiup ctl:v6.5.0 pd  -u http://127.0.0.1:2379 region | jq . | less
```
