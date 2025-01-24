---
title: 反向代理
slug: /examples/httpserver/proxy
keywords: [http, 服务器, 代理, 反向代理, goframe]
description: 使用 GoFrame 框架实现反向代理服务器
hide_title: true
sidebar_position: 1
---

# HTTP 服务器代理

## 介绍

本示例展示了如何使用 `GoFrame` 创建一个反向代理服务器。示例包含两个服务器：

1. 在 8198 端口运行的后端服务器，提供实际服务
2. 在 8199 端口运行的代理服务器，将请求转发到后端服务器

代理服务器实现了以下功能：
- 使用 `httputil.NewSingleHostReverseProxy` 实现反向代理功能
- 自定义代理失败的错误处理
- URL 路径重写
- 请求体处理
- 详细的代理操作日志


## 环境要求

- [Go](https://golang.org/dl/) 1.22 或更高版本
- [Git](https://git-scm.com/downloads)
- [GoFrame](https://goframe.org)

## 使用说明

1. 运行服务:
   ```bash
   go run main.go
   ```

2. 服务监听两个端口:
   - 后端服务器在 http://127.0.0.1:8198
   - 代理服务器在 http://127.0.0.1:8199

3. 测试代理:
   - 通过代理访问：http://127.0.0.1:8199/proxy/user/1
   - 直接访问后端：http://127.0.0.1:8198/user/1

