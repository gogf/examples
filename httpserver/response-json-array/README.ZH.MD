---
title: JSON 数组响应
slug: /examples/httpserver/response-json-array
keywords: [http, 服务器, json, 数组, goframe]
description: 使用 GoFrame 框架处理 HTTP 服务器的 JSON 数组响应
hide_title: true
sidebar_position: 4
---

# HTTP 服务器 JSON 数组响应

## 介绍

本示例展示了如何使用 `GoFrame` 实现一个返回 JSON 数组响应的 HTTP 服务器。它展示了如何：
- 将 API 响应构造为 JSON 数组
- 配置 OpenAPI/Swagger 文档
- 使用 GoFrame 的中间件实现一致的响应处理
- 定义类型安全的请求和响应结构

本示例实现了一个返回 JSON 数组格式用户列表的 `/user` 端点。

## 环境要求

- [Go](https://golang.org/dl/) `1.22` 或更高版本
- [Git](https://git-scm.com/downloads)
- [GoFrame](https://goframe.org)

## 目录结构

- `go.mod`: Go 模块文件
- `main.go`: 主程序入口

## 使用说明

1. 启动服务：
   ```bash
   go run main.go
   ```

2. 测试接口：http://localhost:8199/user

3. 接口文档：http://127.0.0.1:8199/swagger