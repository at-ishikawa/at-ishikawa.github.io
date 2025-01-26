---
date: "2020-03-19T00:00:00Z"
tags:
- mysql
- mysql tuner
title: MySQL Tuner
---

[MySQL Tuner tool](https://github.com/major/MySQLTuner-perl)
---
This is a tool to review a configuration for MySQL server.

### Getting started
To run this tool, we should run following commands.

```shell
wget http://mysqltuner.pl/ -O mysqltuner.pl
wget https://raw.githubusercontent.com/major/MySQLTuner-perl/master/basic_passwords.txt -O basic_passwords.txt
wget https://raw.githubusercontent.com/major/MySQLTuner-perl/master/vulnerabilities.csv -O vulnerabilities.csv
perl mysqltuner.pl
```
