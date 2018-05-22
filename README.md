# Katip - Back End
The KATIP is a tool which will provide a securable text storage service for individuals

## Demo
[Katip product](https://katip.hiyali.org)

## Front end project
[github.com/hiyali/katip-fe](https://github.com/hiyali/katip-fe)

## Requirements
* Golang > 1.9
* MySQL > 5.7

## Install
```shell
go get github.com/hiyali/katip-be/...
```

## Database prepare
```shell
./prepare/sh/create_mysql_user.sh katip_mysql_admin katip_v1_pass
./prepare/sh/create_datebase.sh katip_v1
./prepare/sh/sql_write_in.sh katip_v1
```

## Config
```shell
cd $GOPATH/src/github.com/hiyali/katip-be
cp config_example.yml config.yml
```

Write your configurations in `config.yml`

## Build & Run
```shell
go build
./katip-be
```

## TODO
[] User update information & change password
[] User avatar

## Contribute
> Feel free

## License
MIT
