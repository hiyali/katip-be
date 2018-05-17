# Katip - Back End - prepare
The KATIP is a tool it will provide a securable text storage service support for individuals

## Requirements
* Golang > 1.9
* MySQL > 5.7

## Installation
```shell
go get github.com/hiyali/katip-be/...
cd ${GOHOME}/src/github.com/hiyali/katip-be
```

## Database prepare
```shell
./prepare/sh/create_datebase.sh katip_v1
./prepare/sh/sql_write_in.sh katip_v1
```

## Build & Run
```shell
go build
./katip-be
```

## License
MIT
