#! /bin/bash

set -x

mysql_main_server=main
mysql_user=root
mysql_password=$MYSQL_ROOT_PASSWORD

# Restore a data from a read replica
nc -l -p 9999 | tee >(sha1sum > /tmp/destination_checksum) > /tmp/backup.xbstream &

## Can't receive a checksum from a backup server. Use sleep instead
# nc -l -p 9998 > /tmp/source_checksum &
# while [ ! -s /tmp/source_checksum ]; do
#     sleep 10
# done
# cat /tmp/source_checksum
sleep 300

cat /tmp/destination_checksum

mkdir -p /tmp/backups
xbstream -x --decompress -C /tmp/backups < /tmp/backup.xbstream
xtrabackup --user=$mysql_user --password=$mysql_password --prepare --target-dir=/tmp/backups
rm -rf /var/lib/mysql/*
xtrabackup --copy-back --target-dir=/tmp/backups
# Or rsync -avprP /tmp/backups/* /var/lib/mysql

# Set up a replication connection
gtid=$(awk '{ printf $3 }' /tmp/backups/xtrabackup_binlog_info)
cat <<EOF > /tmp/init.sql
RESET MASTER;
SET GLOBAL gtid_purged='$gtid';
CHANGE MASTER TO
    MASTER_HOST='$mysql_main_server',
    MASTER_USER='repl',
    MASTER_PASSWORD='password',
    MASTER_AUTO_POSITION=1;
START SLAVE;
EOF

mysql --ssl-mode=DISABLED -h localhost -u$mysql_user -p$mysql_password < /tmp/init.sql
