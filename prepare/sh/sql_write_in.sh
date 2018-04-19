#!/bin/bash
. ./utils.sh
if [ $1 ]; then
  DB_NAME=$1
  echo "Please enter mysql root password"
  sql_dir="./prepare/sql/data.sql"
  mysql -uroot -p $DB_NAME < "${sql_dir}"
else
  echo "Please enter database name first"
fi
