---
title: 原生 gRPC 示例
slug: /examples/grpc/rawgrpc
keywords: [grpc, 原生, 实现, goframe]
description: GoFrame 中的原生 gRPC 实现
hide_title: true
sidebar_position: 9
---

## gRPC - 原生示例

### 介绍

本示例展示了如何在 GoFrame 中实现原生 gRPC 服务，不使用额外的抽象层。



### 目录结构

```
.
├── client/           # 客户端示例
│   └── client.go     # 原生 gRPC 客户端实现
├── helloworld/       # protobuf协议定义
│   └── helloworld.proto # 服务和消息定义
├── server/           # 服务器示例
│   └── server.go     # 原生 gRPC 服务器实现
├── go.mod            # Go 模块文件
└── go.sum            # Go 模块校验和
```

### 环境要求

- [Go](https://golang.org/dl/) 1.22 或更高版本
- [Git](https://git-scm.com/downloads)
- [GoFrame](https://goframe.org)
- [Protocol Buffers](https://developers.google.com/protocol-buffers)

### 使用说明

1. 启动服务器：
   ```bash
   cd server
   go run server.go
   ```

2. 运行客户端：
   ```bash
   cd client
   go run client.go
   ```
