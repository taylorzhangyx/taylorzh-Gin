# taylorzh-go

A simple monolithic app serving RESTful APIs.
This repo contains some best practices that could be referred in production and learning purpose.
And all the codes are original written by Yuxin Zhang ([taylorzyx@hotmail.com](https://www.linkedin.com/in/yxzh/))

# Building Blocks

- [Gin Web Framework](https://github.com/gin-gonic/gin)
- [Air - Live reload for Go apps](https://github.com/cosmtrek/air)
- [MySql](https://www.mysql.com/)
- [GORM - The fantastic ORM library for Golang](https://gorm.io/)
- [redis](https://redis.io/)

# Before start
A lot of components are used to make this server more complete and fancy. So be sure to install them as you needed.

## Required components
The following components are required for you to run this app in your local environment:

### MySql DB

Please refer to this address to install MySql on your local machine:
https://dev.mysql.com/downloads/mysql/

Your server should be ready in the following configurations or change the `Makefile` to meet your settings:

```text
ip: localhost(127.0.0.1)
port: 3306
password: 1qaz!QAZ
DB schema: taylorzh
```

## Nice to have components
The following components are used to either improve your dev experience or to enhance the server.

### Air
It's a good idea to do TDD or to verify your code in real-time. This tool offers the ability to auto rebuild and reload your app every time the file is changed.
Be sure to use it to shorten the feedback loop of the dev cycle.

```shell
go install github.com/cosmtrek/air@latest
```

### Redis

```shell
https://redis.io/docs/getting-started/
```


# Quick Start

Run the following command to run this app on your local machine: 
```shell
make dev  # if you installed air

make run  # if you don't have air
```

check the server is running

```shell
curl localhost:8080/healthcheck
# OK 2022-08-20T17:02:23+08:00
```

```shell
curl localhost:8080/hello
# {"count":1,"message":"hello world"}
```

# Feature List

- load counter
- async task runner
- logger

## Load counter and metrics
Sometimes we want to load test the performance of our backend service and to collect the request metrics of the service.
These apis give you the metrics of the request in various dimensions.

TODO
 
## Async Task Runner

TODO

## logger