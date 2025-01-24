---
title: 服务发现
slug: /examples/grpc/resolver
keywords: [grpc, 服务发现, etcd, goframe]
description: GoFrame 中使用 etcd 的 gRPC 服务发现
hide_title: true
sidebar_position: 1
---

# `gRPC` - 服务发现

## 介绍

本示例展示了如何在 `GoFrame` 的 `gRPC` 服务中使用 `etcd` 进行服务发现。


## 目录结构

```text
.
├── client/              # 客户端示例
│   └── client.go        # 带服务发现的客户端
├── controller/          # 服务控制器
│   └── helloworld.go    # Hello 服务实现
├── protobuf/            # protobuf协议定义
│   └── helloworld.proto # 服务和消息定义
├── server/              # 服务器示例
│   └── server.go        # 带服务注册的服务器
├── go.mod               # Go 模块文件
└── go.sum               # Go 模块校验和
```



## 环境要求

- [Go](https://golang.org/dl/) 1.22 或更高版本
- [Git](https://git-scm.com/downloads)
- [GoFrame](https://goframe.org)
- [Protocol Buffers](https://developers.google.com/protocol-buffers)
- [etcd](https://etcd.io/)

## 使用说明

1. 启动 `etcd`：
   ```bash
   docker run -d --name etcd -p 2379:2379 -e ALLOW_NONE_AUTHENTICATION=yes bitnami/etcd:3.4.24
   ```

2. 启动服务器：
   ```bash
   cd server
   go run server.go
   ```

3. 运行客户端：
   ```bash
   cd client
   go run client.go
   ```

