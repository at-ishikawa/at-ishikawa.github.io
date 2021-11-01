FROM mysql/mysql-server:5.7

RUN yum -y update && yum install -y tar gzip wget
RUN wget 'https://github.com/github/gh-ost/releases/download/v1.1.2/gh-ost-binary-linux-20210617134741.tar.gz' \
    && tar -xzvf gh-ost-binary-linux-20210617134741.tar.gz \
    && mv gh-ost /usr/local/bin/

ADD mysql-entrypoint.sh /

ENTRYPOINT "/mysql-entrypoint.sh"
