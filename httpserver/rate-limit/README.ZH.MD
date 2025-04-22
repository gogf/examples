---
title: 访问限流
slug: /examples/httpserver/rate-limit
keywords: [http, 服务器, 限流, 中间件, goframe]
description: 使用 GoFrame 框架实现 HTTP 服务器限流
hide_title: true
sidebar_position: 2
---

# HTTP 服务器限流

## 介绍

本示例展示了如何在使用 `GoFrame` 的 HTTP 服务器中实现限流。
它展示了如何使用 `golang.org/x/time/rate` 包实现的令牌桶算法来保护 `API` 端点免受过多请求的影响。


## 环境要求

- [Go](https://golang.org/dl/) `1.22` 或更高版本
- [Git](https://git-scm.com/downloads)
- [GoFrame](https://goframe.org)

## 目录结构

- `go.mod`: Go 模块文件
- `main.go`: 主程序入口

## 使用说明

1. 启动服务器：
   ```bash
   go run main.go
   ```

2. 测试限流：
   ```bash
   # 正常请求
   curl http://localhost:8000/hello?name=world
   
   # 快速发送多个请求测试限流
   for i in {1..20}; do curl http://localhost:8000/hello?name=world; done
   ```
