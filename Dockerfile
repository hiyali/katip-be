# Author: Salam Hiyali
# Created at: 2019-05-14

# ------------------------- Base & info & deps
# 2019-05-09 Alpine 3.9.4 released
# 54.175.43.85    registry-1.docker.io
FROM alpine:3.9.4
MAINTAINER Salam Hiyali "hiyali920@gmail.com"
ENV REFRESHED_AT 2019-05-15

# remove this block (if you're not in China)
RUN echo "#aliyun" > /etc/apk/repositories
RUN echo "https://mirrors.aliyun.com/alpine/v3.6/main/" >> /etc/apk/repositories
RUN echo "https://mirrors.aliyun.com/alpine/v3.6/community/" >> /etc/apk/repositories

# update
RUN apk update

# add basic deps
RUN apk add openrc --no-cache
RUN apk add git make
RUN apk add ca-certificates  # ca-certificates & ssh for git clone
RUN apk add tzdata  # for set timezone
RUN apk add vim # The reason is that for edit file in docker sometimes
# RUN apk add curl
# RUN apk add cron

# set timezone
ENV TZ Asia/Urumqi

# ------------------------- Add go & mysql deps
RUN apk add --no-cache musl-dev go
# -- config
ENV GOROOT /usr/lib/go
ENV GOPATH /go
ENV PATH /go/bin:$PATH
RUN mkdir -p ${GOPATH}/src ${GOPATH}/bin
# -- mysql
RUN apk add mariadb mariadb-client
# -- config
ENV DB_DATA_PATH="/var/lib/mysql"
ENV DB_ROOT_PASS="root"
ENV DB_NAME="katip_v1"
ENV DB_USER="katip_mysql_admin"
ENV DB_PASS="katip_v1_pass"
RUN addgroup mysql mysql
VOLUME ["/var/lib/mysql"]

# ------------------------- Prepare apps
RUN mkdir -p /katip
WORKDIR /katip
RUN mkdir -p ./logs

COPY prepare/sh/* ./
RUN chmod +x *.sh
COPY prepare/sql/model-data.sql ./data.sql

# ------------------------- Finish
# clear
RUN rm -rf /var/cache/apk/*

# expose
EXPOSE 80
EXPOSE 443

ENTRYPOINT ["./docker_entry.sh"]
