#!/bin/sh

generate_nginx_conf () {
  DOMAIN=$1

  sed "s/{{DOMAIN}}/${DOMAIN}/g" ./prepare/config/katip.conf.tpl > ./prepare/config/katip.conf
  echo "The nginx conf created,with $DOMAIN, in ./prepare/config/katip.conf"
}

if [ $1 ]; then
  generate_nginx_conf $*
else
  echo "Please enter your domain without www."
fi
