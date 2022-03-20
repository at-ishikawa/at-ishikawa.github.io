#! /bin/bash

set -ex
xtrabackup \
  --user=root \
  --password=$MYSQL_ROOT_PASSWORD \
  --backup \
  --stream=xbstream \
  --compress \
  --compress-threads=4 \
  --parallel=2 \
  --target-dir=/tmp/backups \
  | tee >(sha1sum > /tmp/source_checksum) \
  | nc backup 9999
# Next line doesn't work for some reasons
# cat /tmp/source_checksum | nc backup 9998
