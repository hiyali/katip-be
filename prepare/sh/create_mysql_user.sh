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

create_mysql_user () {
  USERNAME=$1
  PASSWORD=$2
  PLACE="localhost"
  if [ $3 ]; then
    PLACE=$3
  fi

  echo "Ready to create user $1 @ $PLACE"
  echo "Please enter mysql root passowrd"
  echo "CREATE USER '$USERNAME'@'$PLACE' IDENTIFIED BY '$PASSWORD';GRANT ALL PRIVILEGES ON *.* TO '$USERNAME'@'$PLACE' WITH GRANT OPTION;" | mysql -u root --passowrd="${DB_ROOT_PASS}"
}

if [ $1 ]; then
  if [ $2 ]; then
    create_mysql_user $*
  else
    echo "Please enter user password at second argument"
  fi
else
  echo "Please enter user name at first argument"
fi
