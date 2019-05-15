# Katip - Back End
The KATIP is a tool which will provide a securable text storage service for individuals

## Demo
[Katip product](https://katip.hiyali.org)

## Front end project
[github.com/hiyali/katip-fe](https://github.com/hiyali/katip-fe)

## Requirements
* Golang 1.9+
* MySQL 5.7+

## Install
```shell
go get github.com/hiyali/katip-be/...
```

## Config
```shell
cd $GOPATH/src/github.com/hiyali/katip-be
cp config_example.yml config.yml
```

#### Write your configurations in `config.yml`

## Docker
> Default user name and password is `hiyali920@gmail.com` `non-secure`, Change `prepare/sql/model-data.sql` if necessry.
```shell
docker build --rm -t katip-be:v1 ./
docker run -d katip-be:v1 # -it # for foreground
```

## TODO
* [x] User login & register with an email address
* [x] Send a email over TLS
* [x] Add & edit & show records
* [x] Filter the records by title & type
* [x] User update information & change password & avatar
* [x] User reset password with email address (not tested)
* [ ] Design a logo for Katip
* [ ] Security (maybe use Symmetric Searchable Encryption)
* [x] Dockerize

## Contribute
> Feel free

## Screenshot
![black-theme](https://raw.githubusercontent.com/hiyali/katip-be/master/screenshot/katip-black-theme.png "black-theme")
![record-list](https://raw.githubusercontent.com/hiyali/katip-be/master/screenshot/katip-record-list.png "record-list")

## License
MIT
