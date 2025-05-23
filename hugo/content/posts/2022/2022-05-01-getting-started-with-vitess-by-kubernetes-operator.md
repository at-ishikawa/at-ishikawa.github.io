---
date: "2022-05-01T00:00:00Z"
tags:
- vitess
- kubernetes
- kubernetes operator
title: Getting Started with Vitess by Kubernetes Operator
---


There are fewer lock contentions to worry about, replication is a lot happier, production impact of outages become smaller, backups and restores run faster, and a lot more secondary advantages can be realized

Getting Started
===

Docker
---

Follow [this tutorial](https://vitess.io/docs/14.0/get-started/local-docker/)

```
git clone git@github.com:vitessio/vitess.git
cd vitess
make docker_local
./docker/local/run.sh
```



Kubernetes Operator
---

### Prerequisite

Follow [this tutorial](https://vitess.io/docs/14.0/get-started/operator/).

```
minikube start --kubernetes-version=v1.19.16 --cpus=4 --memory=4000 --disk-size=32g
```

Install vitess client.
```
version=6.0.20-20200818
file=vitess-${version}-90741b8.tar.gz
wget https://github.com/vitessio/vitess/releases/download/v${version}/${file}
tar -xzf ${file}
cd ${file/.tar.gz/}
sudo mkdir -p /usr/local/vitess
sudo cp -r * /usr/local/vitess/
export PATH=$PATH:/usr/local/vitess/bin
```

### Install Operator

```
git clone git@github.com:vitessio/vitess.git
cd vitess/examples/operator
kubectl apply -f operator.yaml -n vitess
```

Start the first cluster
```
kubectl apply -f 101_initial_cluster.yaml -n vitess
```

Connect to vitess by mysql client

```
> minikube kubectl -- port-forward svc/example-vtgate-ae7df4b6 15306:3306
Forwarding from 127.0.0.1:15306 -> 3306
Forwarding from [::1]:15306 -> 3306
Handling connection for 15306
```

Open a different session
```
> mysql -h 127.0.0.1 -P 15306 -u user
Welcome to the MySQL monitor.  Commands end with ; or \g.
Your MySQL connection id is 1
Server version: 5.7.9-vitess-14.0.0-SNAPSHOT Version: 14.0.0-SNAPSHOT (Git revision 270cf96cd2 branch 'main') built on Sat Mar 26 10:24:01 UTC 2022 by vitess@buildkitsandbox using go1.18 linux/amd64

mysql> show databases;
+--------------------+
| Database           |
+--------------------+
| commerce           |
| information_schema |
| mysql              |
| sys                |
| performance_schema |
+--------------------+
5 rows in set (0.00 sec)
```


Create a schema
```
minikube kubectl -- port-forward svc/example-vtctld-625ee430 15999:15999
```


```
set sql (cat create_commerce_schema.sql)
> vtctlclient -server=localhost:15999 ApplySchema -sql="$sql" commerce
E0326 19:20:53.480875  165620 main.go:64] E0327 02:20:53.480177 vtctl.go:3352] schema change failed, ExecuteResult: {
  "FailedShards": null,
  "SuccessShards": null,
  "CurSQLIndex": 0,
  "Sqls": [
    "create table product(   sku varbinary(128),   description varbinary(128),   price bigint,   primary key(sku) ) ENGINE=InnoDB",
    "create table customer(   customer_id bigint not null auto_increment,   email varbinary(128),   primary key(customer_id) ) ENGINE=InnoDB",
    "create table corder(   order_id bigint not null auto_increment,   customer_id bigint,   sku varbinary(128),   price bigint,   primary key(order_id) ) ENGINE=InnoDB"
  ],
  "UUIDs": null,
  "ExecutorErr": "rpc error: code = Unknown desc = TabletManager.PreflightSchema on zone1-2469782763 error: /usr/bin/mysql: exit status 1, output: ERROR 2013 (HY000) at line 3: Lost connection to MySQL server during query\n: /usr/bin/mysql: exit status 1, output: ERROR 2013 (HY000) at line 3: Lost connection to MySQL server during query\n",
  "TotalTimeSpent": 232941953
}
```

For some reasons, after I ran vschema at first, and then after waiting for a few minutes, then I was able to run the query.
```
set vschema (cat vschema_commerce_initial.json)
vtctlclient -server=localhost:15999 ApplyVSchema -vschema="$vschema" commerce
vtctlclient -server=localhost:15999 ApplySchema -sql="$sql" commerce
```

```
> mysql -h 127.0.0.1 -P 15306 -u user
Welcome to the MySQL monitor.  Commands end with ; or \g.
Your MySQL connection id is 1
Server version: 5.7.9-vitess-14.0.0-SNAPSHOT Version: 14.0.0-SNAPSHOT (Git revision 270cf96cd2 branch 'main') built on Sat Mar 26 10:24:01 UTC 2022 by vitess@buildkitsandbox using go1.18 linux/amd64

mysql> show tables;
+--------------------+
| Tables_in_commerce |
+--------------------+
| corder             |
| customer           |
| product            |
+--------------------+
3 rows in set (0.00 sec)
```

## Move a Table
Follow [this table](https://vitess.io/docs/13.0/user-guides/migration/move-tables/).

Insert a table
```
> mysql -h 127.0.0.1 -P 15306 -u user < ../common/insert_commerce_data.sql
> mysql -h 127.0.0.1 -P 15306 -u user < ../common/select_commerce_data.sql
Using commerce
Customer
customer_id     email
1       alice@domain.com
2       bob@domain.com
3       charlie@domain.com
4       dan@domain.com
5       eve@domain.com
Product
sku     description     price
SKU-1001        Monitor 100
SKU-1002        Keyboard        30
COrder
order_id        customer_id     sku     price
1       1       SKU-1001        100
2       2       SKU-1002        30
3       3       SKU-1002        30
4       4       SKU-1002        30
5       5       SKU-1002        30

> mysql -h 127.0.0.1 -P 15306 -u user --table --execute="show vitess_tablets"
+-------+----------+-------+------------+---------+------------------+-------------+----------------------+
| Cell  | Keyspace | Shard | TabletType | State   | Alias            | Hostname    | PrimaryTermStartTime |
+-------+----------+-------+------------+---------+------------------+-------------+----------------------+
| zone1 | commerce | -     | PRIMARY    | SERVING | zone1-2469782763 | 172.18.0.9  | 2022-03-27T01:45:19Z |
| zone1 | commerce | -     | REPLICA    | SERVING | zone1-2548885007 | 172.18.0.10 |                      |
+-------+----------+-------+------------+---------+------------------+-------------+----------------------+
```

```
> minikube kubectl -- apply -f 201_customer_tablets.yaml
vitesscluster.planetscale.com/example configured
```

Portforward
```
> vim pf.sh
# Edit alias kubectl="minikube kubectl --"

> ./pf.sh
```

```
> mysql -h 127.0.0.1 -P 15306 -u user --table --execute="show vitess_tablets"
+-------+----------+-------+------------+---------+------------------+-------------+----------------------+
| Cell  | Keyspace | Shard | TabletType | State   | Alias            | Hostname    | PrimaryTermStartTime |
+-------+----------+-------+------------+---------+------------------+-------------+----------------------+
| zone1 | commerce | -     | PRIMARY    | SERVING | zone1-2469782763 | 172.18.0.9  | 2022-03-27T01:45:19Z |
| zone1 | commerce | -     | REPLICA    | SERVING | zone1-2548885007 | 172.18.0.10 |                      |
| zone1 | customer | -     | PRIMARY    | SERVING | zone1-1250593518 | 172.18.0.11 | 2022-03-27T02:44:43Z |
| zone1 | customer | -     | REPLICA    | SERVING | zone1-3778123133 | 172.18.0.12 |                      |
+-------+----------+-------+------------+---------+------------------+-------------+----------------------+
```

Move a table into different keyspace
```
> vtctlclient -server=localhost:15999 MoveTables -source commerce -tables 'customer,corder' Create customer.commerce2customer
E0326 19:58:50.201219  262097 main.go:67] remote error: rpc error: code = Unknown desc = TabletManager.GetSchema on zone1-2469782763 error: EOF (errno 2013) (sqlstate HY000) during query: SELECT COLUMN_NAME as column_name
                FROM INFORMATION_SCHEMA.COLUMNS
                WHERE TABLE_SCHEMA = 'vt_commerce' AND TABLE_NAME = 'corder'
                ORDER BY ORDINAL_POSITION
EOF (errno 2013) (sqlstate HY000) during query: SELECT COLUMN_NAME as column_name
                FROM INFORMATION_SCHEMA.COLUMNS
                WHERE TABLE_SCHEMA = 'vt_commerce' AND TABLE_NAME = 'customer'
                ORDER BY ORDINAL_POSITION
EOF (errno 2013) (sqlstate HY000) during query: SELECT COLUMN_NAME as column_name
                FROM INFORMATION_SCHEMA.COLUMNS
                WHERE TABLE_SCHEMA = 'vt_commerce' AND TABLE_NAME = 'product'
                ORDER BY ORDINAL_POSITION
EOF (errno 2013) (sqlstate HY000) during query: SELECT table_name as table_name, ordinal_position as ordinal_position, COLUMN_NAME as column_name
                FROM INFORMATION_SCHEMA.COLUMNS
                WHERE TABLE_SCHEMA = 'vt_commerce'
                AND TABLE_NAME IN ('corder', 'customer', 'product')
                AND COLUMN_KEY = 'PRI'
                ORDER BY table_name, ordinal_position: EOF (errno 2013) (sqlstate HY000) during query: SELECT COLUMN_NAME as column_name
                FROM INFORMATION_SCHEMA.COLUMNS
                WHERE TABLE_SCHEMA = 'vt_commerce' AND TABLE_NAME = 'corder'
                ORDER BY ORDINAL_POSITION
EOF (errno 2013) (sqlstate HY000) during query: SELECT COLUMN_NAME as column_name
                FROM INFORMATION_SCHEMA.COLUMNS
                WHERE TABLE_SCHEMA = 'vt_commerce' AND TABLE_NAME = 'customer'
                ORDER BY ORDINAL_POSITION
EOF (errno 2013) (sqlstate HY000) during query: SELECT COLUMN_NAME as column_name
                FROM INFORMATION_SCHEMA.COLUMNS
                WHERE TABLE_SCHEMA = 'vt_commerce' AND TABLE_NAME = 'product'
                ORDER BY ORDINAL_POSITION
EOF (errno 2013) (sqlstate HY000) during query: SELECT table_name as table_name, ordinal_position as ordinal_position, COLUMN_NAME as column_name
                FROM INFORMATION_SCHEMA.COLUMNS
                WHERE TABLE_SCHEMA = 'vt_commerce'
                AND TABLE_NAME IN ('corder', 'customer', 'product')
                AND COLUMN_KEY = 'PRI'
                ORDER BY table_name, ordinal_position
```

Try to monitor the progress
```
> vtctlclient -server=localhost:15999 VReplicationExec zone1-1250593518 "select * from _vt.copy_state"
> vtctlclient -server=localhost:15999 VReplicationExec zone1-2469782763 "select * from _vt.copy_state"
> vtctlclient -server=localhost:15999 VReplicationExec zone1-2548885007 "select * from _vt.copy_state"
E0326 20:00:50.714505  268220 main.go:67] remote error: rpc error: code = Unknown desc = TabletManager.VReplicationExec on zone1-2548885007 error: vreplication engine is closed: vreplication engine is closed
```

Retry to run move table operations
```
> vtctlclient -server=localhost:15999 MoveTables -source commerce -tables 'customer,corder' Create customer.commerce2customer
E0326 20:02:11.167288  271560 main.go:67] remote error: rpc error: code = Unknown desc = TabletManager.VReplicationExec on zone1-1250593518 error: error in connecting to mysql db with connection <nil>, err net.Dial(/vt/socket/mysql.sock) to local server failed: dial unix /vt/socket/mysql.sock: connect: connection refused (errno 2002) (sqlstate HY000): error in connecting to mysql db with connection <nil>, err net.Dial(/vt/socket/mysql.sock) to local server failed: dial unix /vt/socket/mysql.sock: connect: connection refused (errno 2002) (sqlstate HY000)
> vtctlclient -server=localhost:15999 MoveTables -source commerce -tables 'customer,corder' Create customer.commerce2customer
Waiting for workflow to start:
0% ... 0% ... 0% ... 0% ... 0% ... 0% ... 0% ... 0% ... 0% ... 0% ... 0% ... 0% ... 0% ... 0% ... 0% ...
Workflow started successfully with 1 stream(s)

Following vreplication streams are running for workflow customer.commerce2customer:

id=1 on -/zone1-1250593518: Status: Copying. VStream Lag: 0s.
```

Validate the result
```
> vtctlclient -server=localhost:15999 VDiff customer.commerce2customer
Summary for table corder:
        ProcessedRows: 5
        MatchingRows: 5
        MismatchedRows: 0
        ExtraRowsSource: 0
        ExtraRowsTarget: 0
Summary for table customer:
        ProcessedRows: 5
        MatchingRows: 5
        MismatchedRows: 0
        ExtraRowsSource: 0
        ExtraRowsTarget: 0
```

Check routing routes
```
> vtctlclient -server=localhost:15999 GetRoutingRules commerce

{
  "rules": [
    {
      "fromTable": "customer.customer",
      "toTables": [
        "commerce.customer"
      ]
    },
    {
      "fromTable": "customer@replica",
      "toTables": [
        "commerce.customer"
      ]
    },
    {
      "fromTable": "commerce.customer@replica",
      "toTables": [
        "commerce.customer"
      ]
    },
    {
      "fromTable": "commerce.customer@rdonly",
      "toTables": [
        "commerce.customer"
      ]
    },
    {
      "fromTable": "corder@replica",
      "toTables": [
        "commerce.corder"
      ]
    },
    {
      "fromTable": "customer.customer@rdonly",
      "toTables": [
        "commerce.customer"
      ]
    },
    {
      "fromTable": "corder@rdonly",
      "toTables": [
        "commerce.corder"
      ]
    },
    {
      "fromTable": "customer",
      "toTables": [
        "commerce.customer"
      ]
    },
    {
      "fromTable": "customer.customer@replica",
      "toTables": [
        "commerce.customer"
      ]
    },
    {
      "fromTable": "commerce.corder@replica",
      "toTables": [
        "commerce.corder"
      ]
    },
    {
      "fromTable": "customer.corder@replica",
      "toTables": [
        "commerce.corder"
      ]
    },
    {
      "fromTable": "customer@rdonly",
      "toTables": [
        "commerce.customer"
      ]
    },
    {
      "fromTable": "commerce.corder@rdonly",
      "toTables": [
        "commerce.corder"
      ]
    },
    {
      "fromTable": "customer.corder",
      "toTables": [
        "commerce.corder"
      ]
    },
    {
      "fromTable": "customer.corder@rdonly",
      "toTables": [
        "commerce.corder"
      ]
    },
    {
      "fromTable": "corder",
      "toTables": [
        "commerce.corder"
      ]
    }
  ]
}
```

Switch non-primary traffic
```
> vtctlclient -server=localhost:15999 MoveTables -tablet_types=rdonly,replica SwitchTraffic customer.commerce2customer
SwitchTraffic was successful for workflow customer.commerce2customer
Start State: Reads Not Switched. Writes Not Switched
Current State: All Reads Switched. Writes Not Switched
```
Confirm the new routes
```
> vtctlclient -server=localhost:15999 GetRoutingRules commerce
{
  "rules": [
    {
      "fromTable": "customer@replica",
      "toTables": [
        "customer.customer"
      ]
    },
    {
      "fromTable": "customer",
      "toTables": [
        "commerce.customer"
      ]
    },
    {
      "fromTable": "customer.corder",
      "toTables": [
        "commerce.corder"
      ]
    },
    {
      "fromTable": "commerce.customer@replica",
      "toTables": [
        "customer.customer"
      ]
    },
    {
      "fromTable": "commerce.corder@replica",
      "toTables": [
        "customer.corder"
      ]
    },
    {
      "fromTable": "customer.corder@replica",
      "toTables": [
        "customer.corder"
      ]
    },
    {
      "fromTable": "customer@rdonly",
      "toTables": [
        "customer.customer"
      ]
    },
    {
      "fromTable": "commerce.corder@rdonly",
      "toTables": [
        "customer.corder"
      ]
    },
    {
      "fromTable": "corder",
      "toTables": [
        "commerce.corder"
      ]
    },
    {
      "fromTable": "commerce.customer@rdonly",
      "toTables": [
        "customer.customer"
      ]
    },
    {
      "fromTable": "corder@replica",
      "toTables": [
        "customer.corder"
      ]
    },
    {
      "fromTable": "customer.customer@rdonly",
      "toTables": [
        "customer.customer"
      ]
    },
    {
      "fromTable": "corder@rdonly",
      "toTables": [
        "customer.corder"
      ]
    },
    {
      "fromTable": "customer.customer",
      "toTables": [
        "commerce.customer"
      ]
    },
    {
      "fromTable": "customer.customer@replica",
      "toTables": [
        "customer.customer"
      ]
    },
    {
      "fromTable": "customer.corder@rdonly",
      "toTables": [
        "customer.corder"
      ]
    }
  ]
}
```
This moves commerce of `fromTable` to `customer` of `toTables`.


### Switch traffic to primary
```
> vtctlclient -server=localhost:15999 MoveTables -tablet_types=primary SwitchTraffic customer.commerce2customer

SwitchTraffic was successful for workflow customer.commerce2customer
Start State: All Reads Switched. Writes Not Switched
Current State: All Reads Switched. Writes Switched

```
> vtctlclient -server=localhost:15999 GetRoutingRules commerce
{
  "rules": [
    {
      "fromTable": "customer@replica",
      "toTables": [
        "customer.customer"
      ]
    },
    {
      "fromTable": "customer.corder@replica",
      "toTables": [
        "customer.corder"
      ]
    },
    {
      "fromTable": "commerce.corder@rdonly",
      "toTables": [
        "customer.corder"
      ]
    },
    {
      "fromTable": "commerce.corder",
      "toTables": [
        "customer.corder"
      ]
    },
    {
      "fromTable": "commerce.customer@replica",
      "toTables": [
        "customer.customer"
      ]
    },
    {
      "fromTable": "customer@rdonly",
      "toTables": [
        "customer.customer"
      ]
    },
    {
      "fromTable": "corder",
      "toTables": [
        "customer.corder"
      ]
    },
    {
      "fromTable": "customer.customer@replica",
      "toTables": [
        "customer.customer"
      ]
    },
    {
      "fromTable": "customer.corder@rdonly",
      "toTables": [
        "customer.corder"
      ]
    },
    {
      "fromTable": "commerce.customer",
      "toTables": [
        "customer.customer"
      ]
    },
    {
      "fromTable": "corder@replica",
      "toTables": [
        "customer.corder"
      ]
    },
    {
      "fromTable": "customer.customer@rdonly",
      "toTables": [
        "customer.customer"
      ]
    },
    {
      "fromTable": "customer",
      "toTables": [
        "customer.customer"
      ]
    },
    {
      "fromTable": "commerce.corder@replica",
      "toTables": [
        "customer.corder"
      ]
    },
    {
      "fromTable": "commerce.customer@rdonly",
      "toTables": [
        "customer.customer"
      ]
    },
    {
      "fromTable": "corder@rdonly",
      "toTables": [
        "customer.corder"
      ]
    }
  ]
}
```

Remove the data from the old keyspace.
```
> vtctlclient -server=localhost:15999 MoveTables Complete customer.commerce2customer

Complete was successful for workflow customer.commerce2customer
Start State: All Reads Switched. Writes Switched
Current State: Workflow Not Found
```


```
> mysql -h 127.0.0.1 -P 15306 -u user < ../common/select_commerce_data.sql
Using commerce
Customer
ERROR 1146 (42S02) at line 4: target: commerce.-.primary: vttablet: rpc error: code = NotFound desc = Table 'vt_commerce.customer' doesn't exist (errno 1146) (sqlstate 42S02) (CallerID: user): Sql: "select * from customer", BindVars: {}
```
