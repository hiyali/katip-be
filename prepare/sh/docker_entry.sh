#!/bin/sh

DATEVAR=$(date +%Y-%m-%d_%H-%M-%S)

WORKDIR="/katip"
LOGDIR="${WORKDIR}/logs/entry.log"

date 2>&1 >> $LOGDIR

# --------------------------- Mysql
# mysql -u katip_mysql_admin -p katip_v1_pass katip_v1 < ./model-data.sql 2>&1 >> ./logs/katip_entry.log
mysql_install_db --user=mysql --datadir=${DB_DATA_PATH}
touch /run/openrc/softlevel
rc-service mariadb start

mysqld_safe &
echo 'sleep 5 seconds......'
sleep 5 # waiting for the mysql started

mysqladmin -u root password "${DB_ROOT_PASS}"
# /usr/bin/mysqld --basedir=/usr --datadir=/var/lib/mysql --plugin-dir=/usr/lib/mysql/plugin --user=mysql --pid-file=/run/mysqld/mysqld.pid --socket=/run/mysqld/mysqld.sock --port=3306

echo "GRANT ALL ON *.* TO ${DB_USER}@'127.0.0.1' IDENTIFIED BY '${DB_PASS}' WITH GRANT OPTION;" > /tmp/sql
echo "GRANT ALL ON *.* TO ${DB_USER}@'localhost' IDENTIFIED BY '${DB_PASS}' WITH GRANT OPTION;" >> /tmp/sql
echo "GRANT ALL ON *.* TO ${DB_USER}@'::1' IDENTIFIED BY '${DB_PASS}' WITH GRANT OPTION;" >> /tmp/sql
echo "DELETE FROM mysql.user WHERE User='';" >> /tmp/sql
echo "DROP DATABASE test;" >> /tmp/sql
echo "FLUSH PRIVILEGES;" >> /tmp/sql
cat /tmp/sql | mysql -u root --password="${DB_ROOT_PASS}"

rc-update add mariadb default
# reboot

# /usr/bin/mysqld_safe # --datadir=${DB_DATA_PATH}

# ./create_mysql_user.sh katip_mysql_admin katip_v1_pass 2>&1 >> $LOGDIR
./create_datebase.sh $DB_NAME 2>&1 >> $LOGDIR
./sql_write_in.sh $DB_NAME 2>&1 >> $LOGDIR
echo 'MySql complete......'

# --------------------------- Golang
go get github.com/hiyali/katip-be
cd $GOPATH/src/github.com/hiyali/katip-be/
export GO111MODULE=on
go get ./...
go build 2>&1 >> $LOGDIR
./katip-be

# --------------------------- Nginx
service nginx restart

# --------------------------- Finish
/bin/sh
