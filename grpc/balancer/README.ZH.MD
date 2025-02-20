---
title: 负载均衡
slug: /examples/grpc/balancer
keywords: [grpc, 负载均衡, 服务发现, goframe]
description: GoFrame 中的 gRPC 负载均衡
hide_title: true
sidebar_position: 2
---

# `gRPC` - 负载均衡

## 介绍

本示例展示了如何在 `GoFrame` 应用程序中实现 `gRPC` 负载均衡。


## 目录结构

```text
.
├── client/         # 客户端示例
│   └── client.go   # 带负载均衡的客户端实现
├── controller/     # 服务控制器
│   └── hello.go    # Hello 服务实现
├── protobuf/       # protobuf协议定义
├── server/         # 服务器示例
│   ├── config.yaml # 服务器配置
│   └── server.go   # 服务器实现
├── go.mod          # Go 模块文件
└── go.sum          # Go 模块校验和
```



## 环境要求

- [Go](https://golang.org/dl/) `1.22` 或更高版本
- [Git](https://git-scm.com/downloads)
- [GoFrame](https://goframe.org)
- [Protocol Buffers](https://developers.google.com/protocol-buffers)

## 使用说明

1. 启动多个服务器实例（使用随机端口）：
   ```bash
   cd server
   go run server.go

   # 启动第二个服务器实例
   go run server.go

   # 启动第三个服务器实例
   go run server.go
   ```

2. 运行客户端：
   ```bash
   cd client
   go run client.go
   ```
