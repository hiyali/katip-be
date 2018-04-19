#!/bin/bash
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
  USERNAME="root"
  if [ $2 ]
  then
    USERNAME=$2
  fi

  if [ $1 ]
  then
    echo "create database $1" | mysql -u $USERNAME -p
  else
    echo "Please enter database name follow sh file first"
  fi
}
create_datebase $*
