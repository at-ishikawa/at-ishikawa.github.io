---
title: MySQL backup and restore
date: 2022-03-19T00:00:00Z
---

In this article, explain how to backup MySQL database using [Percona Xtrabackup](https://www.percona.com/doc/percona-xtrabackup/8.0/index.html).
There are two binaries, innobackupex and xtrabackup.
innobackupex is a wrapper for xtrabackup and it was going to be removed. It might have been deleted from version 8.0.


Getting Started
===

Installation
---

Follow [this tutorial](https://www.percona.com/doc/percona-xtrabackup/8.0/installation.html)

```
wget https://downloads.percona.com/downloads/Percona-XtraBackup-LATEST/Percona-XtraBackup-8.0.23-16/binary/tarball/percona-xtrabackup-8.0.23-16-Linux-x86_64.glibc2.17.tar.gz
tar xvf percona-xtrabackup-8.0.23-16-Linux-x86_64.glibc2.17.tar.gz
mv percona-xtrabackup-8.0.23-16-Linux-x86_64.glibc2.17 /usr/lib
export PATH=/usr/lib/percona-xtrabackup-8.0.23-16-Linux-x86_64.glibc2.17 /usr/lib:$PATH
```

Run backup
---

### Full backup
See [this page](https://www.percona.com/doc/percona-xtrabackup/8.0/backup_scenarios/full_backup.html) for more information about full backup and a streaming backup examples.

The backup can take a long time, depending on how large the database is.
It is safe to cancel at any time, because xtrabackup does not modify the database.
```
xtrabackup --user=root --password=password --backup /tmp/backups
```

In order to recovery from the above backup, at first it needs to prepare it
```
xtrabackup --user=root --password=password --prepare --target-dir=/tmp/backups/
```

Then recover a data from backup data
```
xtrabackup --copy-back --target-dir=/data/backups/
```

#### Streaming backups
Use a streaming option to send the backup data to other host with 9999 port.
See [this page](https://www.percona.com/doc/percona-xtrabackup/LATEST/xtrabackup_bin/backup.streaming.html) for more information about streaming backup.

In order to backup and send it to the
```
DESTINATION_HOST=backup
xtrabackup --user=root --password=password --backup --stream=xbstream --compress --compress-threads=4 --parallel=2 ./ | tee >(sha1sum > source_checksum) | nc $DESTINATION_HOST 9999
```

To receive a backup file from the port 9999 on the different host
```
nc -l -p 9999 | tee >(sha1sum > /tmp/destination_checksum) > /tmp/backup.xbstream
```

From this xbstream file, the original data can be extracted
```
xbstream -x --decompress -C /tmp/backups < /tmp/backup.xbstream
```

Then it can be recovered by the same way as the normal backup and restore.
```
xtrabackup --user=root --password=password --prepare --target-dir=/tmp/backups
xtrabackup --copy-back --target-dir=/tmp/backups
```



Use case: How to create a new (or repair a broken) GTID-based Replica
---

Follow [this tutorial](https://www.percona.com/doc/percona-xtrabackup/8.0/howtos/recipes_ibkx_gtid.html).

At first, backup and restore the data at first.

On the source host, backup a file
```
DESTINATION_HOST=backup
xtrabackup --user=root --password=password --backup --stream=xbstream --compress --compress-threads=4 --parallel=2 ./ | tee >(sha1sum > source_checksum) | nc $DESTINATION_HOST 9999
cat source_checksum
```

On the destination host, restore a file
```
nc -l -p 9999 | tee >(sha1sum > /tmp/destination_checksum) > /tmp/backup.xbstream
cat destination_checksum
xbstream -x --decompress -C /tmp/backups < /tmp/backup.xbstream
xtrabackup --user=root --password=password --prepare --target-dir=/tmp/backups
xtrabackup --copy-back --target-dir=/tmp/backups
```

Then setup a replication configuration correctly
At first, check GTID in `xtrabackup_binlog_info` file.

```
gtid=$(awk '{ printf $3 }' /tmp/backups/xtrabackup_binlog_info)
```

And these SQLs to start a replication.
```
RESET MASTER;
SET GLOBAL gtid_purged='$gtid';
CHANGE MASTER TO
    MASTER_HOST='$source_server',
    MASTER_USER='$replication_user',
    MASTER_PASSWORD='$replication_password',
    MASTER_AUTO_POSITION=1;
START SLAVE;
```
