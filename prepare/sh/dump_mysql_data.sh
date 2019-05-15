#!/bin/sh
. ./utils.sh

if [ $1 ]; then
  sql_dir="./dumped_data.sql"
  mysqldump -u $DB_USER --password="${DB_PASS}" $DB_NAME > "${sql_dir}"
else
  echo "Please enter database name first"
fi
