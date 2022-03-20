# 8.0.20 or higher MySQL doesn't support Percona Xtrabackup 8.0 https://www.percona.com/doc/percona-xtrabackup/8.0/index.html
FROM mysql:8.0.19

# Fix an error
# W: GPG error: http://repo.mysql.com/apt/debian buster InRelease: The following signatures couldn't be verified because the public key is not available: NO_PUBKEY 467B942D3A79BD29
RUN apt-key adv --keyserver keyserver.ubuntu.com --recv-keys 467B942D3A79BD29

RUN apt update && apt install -y wget netcat rsync
RUN wget https://downloads.percona.com/downloads/Percona-XtraBackup-LATEST/Percona-XtraBackup-8.0.23-16/binary/tarball/percona-xtrabackup-8.0.23-16-Linux-x86_64.glibc2.17.tar.gz \
    && tar xvf percona-xtrabackup-8.0.23-16-Linux-x86_64.glibc2.17.tar.gz \
    && mv percona-xtrabackup-8.0.23-16-Linux-x86_64.glibc2.17 /usr/lib \
    && apt install -y libgssapi-krb5-2

ENV PATH="/usr/lib/percona-xtrabackup-8.0.23-16-Linux-x86_64.glibc2.17/bin:$PATH"

# RUN apt-get update && apt-get install -y tar gzip wget
# RUN wget 'https://github.com/github/gh-ost/releases/download/v1.1.2/gh-ost-binary-linux-20210617134741.tar.gz' \
#     && tar -xzvf gh-ost-binary-linux-20210617134741.tar.gz \
#     && mv gh-ost /usr/local/bin/
