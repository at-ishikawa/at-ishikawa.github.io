MySQL cluster example
===

MySQL 5.7 Cluster
* main: Main DB
* read-replica-1: Read replica DB 1
* read-replica-2: Read replica DB 2

[Orchestrator](https://github.com/openark/orchestrator)
* orchestrator: Application container for orchestrator
* orchestrator-mysql: MySQL for orchestrator container

Image
---
Each image contains
1. [gh-ost](https://github.com/github/gh-ost)
