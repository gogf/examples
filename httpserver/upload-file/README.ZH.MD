---
title: 文件上传
slug: /examples/httpserver/upload-file
keywords: [http, 服务器, 文件, 上传, goframe]
description: 使用 GoFrame 框架处理 HTTP 服务器的文件上传
hide_title: true
sidebar_position: 3
---

# HTTP 服务器文件上传

## 介绍

本示例展示了如何使用 `GoFrame` 在 HTTP 服务器中实现文件上传功能。它展示了：
- 现代化且用户友好的文件上传界面
- 服务器端文件上传处理
- 文件上传进度跟踪
- 适当的错误处理和验证
- 最大文件大小配置

本示例同时提供了 `REST API` 端点和 `Web` 界面用于文件上传。

## 环境要求

- [Go](https://golang.org/dl/) `1.22` 或更高版本
- [Git](https://git-scm.com/downloads)
- [GoFrame](https://goframe.org)

## 目录结构

- `go.mod`: Go 模块文件，用于依赖管理
- `main.go`: 主程序入口
- `resource/`: 静态资源目录
  - `public/`: 公共资源
    - `upload.html`: 上传页面
  - `upload/`: 上传文件存储目录

## 使用说明

1. 启动服务：
   ```bash
   go run main.go
   ```

2. 上传页面：http://localhost:8000/upload.html


