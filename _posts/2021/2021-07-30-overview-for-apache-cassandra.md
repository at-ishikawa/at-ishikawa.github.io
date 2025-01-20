---
date: 2021-07-30
title: Overviews for Apache cassandra
tags:
  - apache cassandra
---

Apache Cassandra

<iframe width="560" height="315" src="https://www.youtube.com/embed/iDhIjrJ7hG0" title="YouTube video player" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture" allowfullscreen></iframe>

* Column Store
* Structure
    * Clusters: Container for Keyspaces
    * Keyspace = DB in RDBMS
	     * Confiure replication strategies, for example
	* Column Family = Table in RDBMS
	     * Column can be added at any given time


<iframe width="560" height="315" src="https://www.youtube.com/embed/oawc4doC76U" title="YouTube video player" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture" allowfullscreen></iframe>

* Architecture explanation
* Read
    * Anti-Entropy: use latest data from one node
* Write
    * Append only even for updates
    	* Compaction is used not to blow disk out
