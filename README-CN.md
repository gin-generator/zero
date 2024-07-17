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

## 🤷‍ 什么是 gin-zero ？

[English](README.md) | 简体中文

gin-zero是一个基于gin的golang应用程序脚手架，支持一键生成HTTP, GRPC, Websocket 应用，也可以用来构建微服务应用。
它在内部集成了许多优秀的开源包，还提供了一些模块化的微服务治理包。

gin-zero拥有名为`ginctl`的代码生成工具。

你还在为编写数据库的映射结构而烦恼吗？你是否仍在努力编写受数据库字段限制的验证请求结构？使用`ginctl`，它可以帮助你通过数据库注释编写数据库映射结构和验证结构。

## 能力

* 只需要一条命令就可以生成HTTP、GRPC、Websocket服务
* 一条命令生成Restful API，或者生成一个单个接口，同时帮你生成好了接口验证结构体和业务校验代码
* 一条命令生成你的http服务中间件或全局中间件
* 一条命令生成你的服务配置文件etc/env.yaml
* 一条命令生成你的服务部署文件夹，它包含docker镜像生成文件和k8s部署文件
* 支持全量SQL日志，提供记录日志方法logger

## 使用包

* [Gin](https://github.com/gin-gonic/gin)
* [Gorm](https://github.com/go-gorm/gorm)
* [Viper](https://github.com/spf13/viper)
* [Cobra](https://github.com/spf13/cobra)
* [Validator](https://github.com/go-playground/validator)
* ...

## 环境

* golang版本尽量高于1.18

## 快速开始

1. 一个关于http服务的使用案例: [demo](https://github.com/gin-generator/zero/tree/master/app/http/demo)
2. 安装

```shell
git clone https://github.com/gin-generator/zero
cd zero
go install github.com/gin-generator/ginctl@latest
ginctl -h
# 示例
ginctl apply --app admin --module http # 一条命令生成admin应用的http服务 
# 修改你的配置文件，他在app/http/admin/etc/env.yaml 
ginctl api -a admin -m http -l user -o ping -d test # 为admin应用生成一个user模块的ping接口，并指定接口描述为test
# 接下来在你的接口中补充你的业务代码，它的位置应该在app/http/admin/logic/user/logic.go，相同位置的types.go文件存放的是你的请求接收结构体和验证逻辑
```

3. http服务目录结构

* 服务目录

```
app/http/demo # 示例http服务demo的目录结构
├── demo.go # 服务启动入口文件
├── deploy  # 部署文件夹
│   ├── DockerFile
│   ├── Makefile
│   ├── certificate.yaml # 为你的服务订阅免费ssl
│   ├── gateway.yaml     # k8s负载均衡gateway api配置文件
│   └── k8s.yaml         # k8s部署文件
├── etc
│   └── env.yaml         # 服务配置文件
├── logic
│   ├── demo             # 业务代码目录
│   │   ├── logic.go    # http接口
│   │   └── types.go    # http接口的请求结构体存放文件和业务校验文件
│   └── logic.go # 所有http接口的能力继承文件
├── middleware  # 局域中间件
│   └── auth.go
├── route  # 路由文件夹
│   └── route.go
└── storage # 日志
    └── logs-2024-07-06.log
```

## 行为准则

为了确保gin-zero社区欢迎所有人，请审阅并遵守 [Code of Conduct](https://github.com/gin-generator/zero/blob/master/code-of-conduct).

## 安全漏洞

如果你发现gin-zero存在安全漏洞，请通过[golang2024@163.com](mailto:golang2024@163.com)发邮件给Joey。所有安全漏洞将被及时解决。

## 许可证

gin-zero框架是开源软件，在 [MIT license](https://opensource.org/licenses/MIT).
