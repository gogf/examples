---
title: HTTP负载均衡
slug: /examples/balancer/http
keywords: [负载均衡, http, 服务发现, goframe]
description: GoFrame 中的 HTTP 服务负载均衡示例
hide_title: true
---

# HTTP 服务负载均衡示例

## 介绍

本示例展示了如何在 GoFrame 应用程序中实现 HTTP 服务的负载均衡。主要演示：
- 配置负载均衡器
- 注册和发现服务
- 请求路由和分发

## 目录结构

```
.
├── server/         # HTTP 服务器
├── client/         # HTTP 客户端
├── go.mod          # Go 模块文件
└── go.sum          # Go 模块校验
```

## 环境要求

- [Go](https://golang.org/dl/) 1.22 或更高版本
- [Git](https://git-scm.com/downloads)
- [GoFrame](https://goframe.org)

## 使用说明

1. 启动服务器：
   ```bash
   go run server/main.go
   ```

2. 在另一个终端启动客户端：
   ```bash
   go run client/main.go
   ```

