---
date: "2023-02-11T19:00:00Z"
last_modified_at: "2023-04-02"
tags:
- tidb
title: Online Schema Change on TiDB
---

There is a good video to describe a algorithms of TiDB:

<iframe width="560" height="315" src="https://www.youtube-nocookie.com/embed/IiiFVr6UEIc?start=1871" title="YouTube video player" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share" allowfullscreen></iframe>

Also, there are a few documents related to online DDL
- The Online DDL design is descrbed in [this design document](https://github.com/pingcap/tidb/blob/master/docs/design/2018-10-08-online-DDL.md)
- Google F1 algorithm, which is refered from TiDB online schema change, is described in this [research paper](https://static.googleusercontent.com/media/research.google.com/en//pubs/archive/41376.pdf)


# Online Schema change

It seems it's unable to confirm the schema version on TiKV node.
Inside a PD node, it's stored in the `/tidb/ddl/global_schema_version` in the etcd according to [this article](https://docs.pingcap.com/tidb/stable/tidb-computing#metadata-management).

