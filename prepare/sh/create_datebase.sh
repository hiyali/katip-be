#!/bin/sh
# shell example:
# function how_many {
#   echo "$# arguments were supplied."
#   echo $1
#   return 10
# }
# how_many "$*" # args all to one
# how_many "$@" # args all
# how_many "$#" # args count
# $? # => 10

create_datebase () {
  if [ $1 ]; then
    DATABASE_NAME=$1
    echo "create_datebase $DATABASE_NAME"
    echo "create database $DATABASE_NAME" | mysql -u $DB_USER -p # --password="${DB_PASS}"
  else
    echo "Please enter database name follow sh file first"
  fi
}
create_datebase $*
