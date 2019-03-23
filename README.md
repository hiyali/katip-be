# Katip - Back End
The KATIP is a tool which will provide a securable text storage service for individuals

## Demo
[Katip product](https://katip.hiyali.org)

## Front end project
[github.com/hiyali/katip-fe](https://github.com/hiyali/katip-fe)

## Requirements
* Golang >= 1.9
* MySQL >= 5.7

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
* [x] User login & register with an email address
* [x] Send a email over TLS
* [x] Add & edit & show records
* [x] Filter the records by title & type
* [x] User update information & change password & avatar
* [ ] User reset password with email address
* [ ] Design a logo for Katip
* [ ] Security (maybe use Symmetric Searchable Encryption)
* [ ] Dockerize

## Contribute
> Feel free

## Screenshot
![black-theme](https://raw.githubusercontent.com/hiyali/katip-be/master/screenshot/katip-black-theme.png "black-theme")
![record-list](https://raw.githubusercontent.com/hiyali/katip-be/master/screenshot/katip-record-list.png "record-list")

## License
MIT
