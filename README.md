# gin-zero

[//]: # (<p align="center">)

[//]: # (<img align="center" width="150px" src="https://raw.githubusercontent.com/zeromicro/zero-doc/main/doc/images/go-zero.png">)

[//]: # (</p>)

[//]: # (go-zero is a web and rpc framework with lots of builtin engineering practices. It’s born to ensure the stability of the busy services with resilience design and has been serving sites with tens of millions of users for years.)

<div align=center>

[//]: # ([![Go]&#40;https://github.com/zeromicro/go-zero/workflows/Go/badge.svg?branch=master&#41;]&#40;https://github.com/gin-generator/zero/actions&#41;)

[//]: # ([![codecov]&#40;https://codecov.io/gh/zeromicro/go-zero/branch/master/graph/badge.svg&#41;]&#40;https://codecov.io/gh/zeromicro/go-zero&#41;)

[//]: # ([![Go Report Card]&#40;https://goreportcard.com/badge/github.com/zeromicro/go-zero&#41;]&#40;https://goreportcard.com/report/github.com/zeromicro/go-zero&#41;)
[![Release](https://img.shields.io/github/v/release/gin-generator/zero.svg?style=flat-square)](https://github.com/gin-generator/zero)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

[//]: # ([![Go Reference]&#40;https://pkg.go.dev/badge/github.com/zeromicro/go-zero.svg&#41;]&#40;https://pkg.go.dev/github.com/zeromicro/go-zero&#41;)

[//]: # ([![Awesome Go]&#40;https://cdn.rawgit.com/sindresorhus/awesome/d7305f38d29fed78fa85652e3a63e154dd8e8829/media/badge.svg&#41;]&#40;https://github.com/avelino/awesome-go&#41;)

[//]: # ([![Discord]&#40;https://img.shields.io/discord/794530774463414292?label=chat&logo=discord&#41;]&#40;https://discord.gg/4JQvC5A4Fe&#41;)

</div>

## 🤷‍ What is gin-zero?

English | [简体中文](README-CN.md)

Gin-zero is a gin-based golang application scaffolding that supports one-click generation of HTTP, GRPC, Websocket
applications, and can also be used to build microservice applications.
It integrates many excellent open source packages internally, and also provides some modular microservices governance
packages.

Gin-zero has a code generation tool called `ginctl`.

Are you still worried about writing the mapping structure of the database? Are you still struggling to write database field-restricted validation request structs? Get started with `ginctl`, which helps you write database mapping structures and validation structures through database annotations.

## Ability statement

* One command generates the entire http,grpc,websocket service
* A command generates a Restful API, or a custom interface
* A command generates middleware or global middleware
* A command generates the service configuration file
* One command to generate the service deployment folder

## Support package

* [Gin](https://github.com/gin-gonic/gin)
* [Gorm](https://github.com/go-gorm/gorm)
* [Viper](https://github.com/spf13/viper)
* [Cobra](https://github.com/spf13/cobra)
* [Validator](https://github.com/go-playground/validator)
* ...

## Environment
* golang 1.18^

## Quick start

1. An example of http application: [demo](https://github.com/gin-generator/zero/tree/master/app/http/demo)
2. Install

```shell
git clone https://github.com/gin-generator/zero
cd zero
go install github.com/gin-generator/ginctl@latest
ginctl -h
# example
ginctl apply --app admin --module http # A command generates the admin application of the http request
# Next modify your etc/env.yaml file 
ginctl api -a admin -m http -l user -o ping -d test # Generate an http request interface 'ping' for the user module in the admin application
# Then go to the route directory and fill in your request method for the ping interface
```

3. Directory description

* app: application directory contains http, grpc, and websocket
* Application directory

```
app/http/demo
├── demo.go # main.go
├── deploy  # deployment document
│   ├── DockerFile
│   ├── Makefile
│   ├── certificate.yaml # Subscribe to ssl encrypted files for domain names for free
│   ├── gateway.yaml     # Kubernetes gateway configuration file
│   └── k8s.yaml         # Kubernetes deployment file
├── etc
│   └── env.yaml         # Modify your application configuration
├── logic
│   ├── demo             # Service code
│   │   ├── logic.go    # http request controller
│   │   └── types.go    # http request checksum request structure
│   └── logic.go
├── middleware  # middleware
│   └── auth.go
├── route  # Routing file
│   └── route.go
└── storage # Log file
    └── logs-2024-07-06.log
```

## Code of Conduct

In order to ensure that the gin-zero community is welcoming to all, please review and abide by
the [Code of Conduct](https://github.com/gin-generator/zero/blob/master/code-of-conduct).

## Security Vulnerabilities

If you discover a security vulnerability within gin-zero, please send an e-mail to Joey via [golang2024@163.com](mailto:golang2024@163.com). All security vulnerabilities will be promptly addressed.

## License

The gin-zero framework is open-sourced software licensed under the [MIT license](https://opensource.org/licenses/MIT).
