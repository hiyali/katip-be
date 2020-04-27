#!/bin/sh
. ./utils.sh
if [ $1 ]; then
  DATABASE_NAME=$1
  echo "sql_write_in $DATABASE_NAME"
  sql_dir="./data.sql"
  mysql -u $DB_USER -p $DATABASE_NAME < "${sql_dir}"
else
  echo "Please enter database name first"
fi
