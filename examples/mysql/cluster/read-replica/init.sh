#! /bin/bash

set -x

mysql_main_server=main
mysql_user=root
mysql_password=$MYSQL_ROOT_PASSWORD

# Wait for main db
while ! mysqladmin ping -h "$mysql_main_server" -u $mysql_user -p$mysql_password --silent; do
    sleep 10
done

# Backup a main db
mysqldump -h $mysql_main_server -u$mysql_user -p$mysql_password --all-databases --master-data > /tmp/db.dump

# Restore a main db
mysql -h localhost -u$mysql_user -p$mysql_password < /tmp/db.dump

# Set up a replication connection
cat <<EOF > /tmp/init.sql
STOP SLAVE;
CHANGE MASTER TO
    MASTER_HOST='$mysql_main_server',
    MASTER_USER='repl',
    MASTER_PASSWORD='password',
    MASTER_AUTO_POSITION=1;
START SLAVE;
EOF

mysql -h localhost -u$mysql_user -p$mysql_password < /tmp/init.sql
