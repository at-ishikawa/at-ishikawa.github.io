---
date: 2020-03-20
title: Getting Started with MySQL DDL
tags:
  - mysql
---

This document is described based on MySQL 5.6.

Online DDL
===
On MYSQL 5.6, DDL can be ran online, without maintenance.
There are two important types of online DDL.

1. [Algorithm](https://dev.mysql.com/doc/refman/5.6/en/alter-table.html#alter-table-performance)
    * COPY: Concurrent DML is not supported. New table is copied from original table and.
    * INPLACE: Concurrent DML may be supported, but sometimes, exclusive lock is taken. Avoid copying tables to update schema.
	* DEFAULT: Use `INPLACE` if it is supported for running DDL. Otherwise, `COPY`.
1. [LOCK](https://dev.mysql.com/doc/refman/5.6/en/alter-table.html#alter-table-concurrency)
    * NONE: Permit concurrent reads and writes, or error occurs.
	* SHARED: Permit concurrent reads but block writes, or error occurs.
	* EXCLUSIVE: Block reads and writes.
	* DEFAULT: Choose one way from `NONE`, `SHARED`, or `EXCLUSIVE` by this priority.

It's recommended to read [the official document](https://dev.mysql.com/doc/refman/5.6/en/innodb-online-ddl-operations.html#online-ddl-table-operations) at first to understand the more details for what happens in each operation.
