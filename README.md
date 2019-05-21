# Katip - Back End
The KATIP is a tool which will provide a securable text storage service for individuals

## Demo
[Katip product](https://katip.hiyali.org)

## Front end project
[github.com/hiyali/katip-fe](https://github.com/hiyali/katip-fe)

## Requirements
* Docker

## Install
```shell
curl -fsSL https://get.docker.com -o get-docker.sh
sh get-docker.sh
```

```shell
git clone https://github.com/hiyali/katip-be
```

## Config
```shell
cd katip-be
./prepare/sh/generate_nginx_conf.sh hiyali.org # replace with your domain (without www.)
cp config_example.yml config.yml
```

#### Write your configurations in `config.yml`

## Docker
> Default user name and password is `hiyali920@gmail.com` `non-secure`, Change `prepare/sql/model-data.sql` if necessary.
```shell
docker build --rm -t katip:v1 ./
docker run --name katip --rm -id -p 80:80 -p 443:443 katip:v1
```
> replace `-id` to `-it` for foreground run and use `Ctrl+p` - `Ctrl+q` to detach
It will take a few minutes... and `curl -G localhost/api/ping`

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
