#!/bin/sh
OS="`uname`"
case $OS in
  'Linux')
    OS='Linux'
    home='/home/'
    ;;
  'FreeBSD')
    OS='FreeBSD'
    ;;
  'WindowsNT')
    OS='Windows'
    ;;
  'Darwin')
    OS='Mac'
    home='/Users/'
    ;;
  'SunOS')
    OS='Solaris'
    ;;
  'AIX') ;;
  *) ;;
esac

if [ $home ]; then
  user_dir="${home}`whoami`"
else
  echo "This OS is not support"
  exit 1
fi
