#! /bin/bash


mysql_user=root
mysql_password=$MYSQL_ROOT_PASSWORD

# Wait for main db
cd /test_db && \
    mysql -u $mysql_user -p$mysql_password --silent < employees.sql && \
    mysql -u $mysql_user -p$mysql_password --silent -t < test_employees_sha.sql
