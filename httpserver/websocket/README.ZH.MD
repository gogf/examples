---
title: WebSocket
slug: /examples/httpserver/websocket
keywords: [websocket, server, client, goframe, https]
description: 演示使用GoFrame实现WebSocket服务器和客户端用于双向实时通信。本示例展示了带有连接处理的WebSocket服务器设置、客户端连接建立和管理、双向消息发送和接收、连接生命周期管理(连接、断开、错误)、从HTTP到WebSocket的协议升级、以及与GoFrame服务器组件集成。主要功能包括:全双工通信、自动ping/pong心跳、连接池和广播、消息分帧和编码、TLS/SSL支持安全WebSocket(wss://)、生产级模式。非常适合构建实时聊天和消息应用、实现实时通知和更新、创建多玩家游戏服务器、实现协同编辑功能、以及构建实时数据流应用。通过本示例可以学习WebSocket协议、实时双向通信、以及构建高性能实时应用。
hide_title: true
sidebar_position: 5
---

# `WebSocket` 服务器和客户端

## 描述

本示例展示了如何使用 `GoFrame` 实现 `WebSocket` 通信。示例包含 `HTTP` 和 `HTTPS` 两种实现，主要组件包括：

1. 支持安全和非安全连接的 `WebSocket` 服务器
2. 支持 `HTTP` 和 `HTTPS` 的 `WebSocket` 客户端实现
3. `WebSocket` 消息交换示例

实现展示了以下功能：
- 基本的 `WebSocket` 服务器设置
- 安全 `WebSocket`（`WSS`）实现
- 客户端连接处理
- 客户端和服务器之间的消息交换
- 正确的连接清理

## 环境要求

- [Go](https://golang.org/dl/) `1.22` 或更高版本
- [Git](https://git-scm.com/downloads)
- [GoFrame](https://goframe.org)
- [Gorilla WebSocket](https://github.com/gorilla/websocket)

## 目录结构

```text
websocket/
├── http/
│   ├── client.go      # HTTP WebSocket客户端
│   ├── server.go      # HTTP WebSocket服务器
│   └── static/        # 静态文件目录
├── https/
│   ├── client.go      # HTTPS WebSocket客户端
│   ├── server.go      # HTTPS WebSocket服务器
│   ├── static/        # 静态文件目录
│   └── certs/         # SSL证书
└── README.MD
```

## 功能特性

- `WebSocket` 服务器实现
- 安全 `WebSocket`（`WSS`）支持
- 客户端连接处理
- 消息回显功能
- 连接生命周期管理
- 来源检查（可配置）
- 错误处理

## 安装设置

1. 克隆仓库：
    ```bash
    git clone https://github.com/gogf/examples.git
    cd examples/httpserver/websocket
    ```

2. 安装依赖：
    ```bash
    go mod tidy
    ```

3. 对于`HTTPS/WSS`支持，生成自签名证书（可选）：
    ```bash
    mkdir -p https/certs
    openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout https/certs/server.key -out https/certs/server.crt
    ```

## 使用方法

### HTTP WebSocket

1. 启动服务器：
   ```bash
   cd http
   go run server.go
   ```

2. 运行客户端：
   ```bash
   cd http
   go run client.go
   ```

3. 通过浏览器访问演示页面：

   在浏览器中打开 `http://127.0.0.1:8000` 查看 `WebSocket` 演示效果。页面提供了一个简单的聊天界面，可以实时发送和接收消息。

### HTTPS WebSocket

1. 启动安全服务器：
   ```bash
   cd https
   go run server.go
   ```

2. 运行安全客户端：
   ```bash
   cd https
   go run client.go
   ```

3. 通过浏览器访问演示页面：

   在浏览器中打开 `https://127.0.0.1:8000` 查看 `WebSocket` 演示效果。
   注意：由于使用自签名证书，浏览器可能会显示安全警告，这是正常的。

## 实现细节

### 服务器特性
- `WebSocket` 升级处理
- 消息回显功能
- 连接生命周期管理
- 可配置的来源检查
- 错误处理和日志记录

### 客户端特性
- 连接建立
- 消息发送和接收
- 安全连接的 `TLS` 配置
- 清理的连接关闭

## 注意事项

- `WebSocket` 服务器会回显收到的任何消息
- `HTTPS`/`WSS` 示例中使用了自签名证书
- 在生产环境中，需要实现适当的来源检查
- 正确处理连接关闭以避免资源泄漏

## 更多信息

更多关于`WebSocket`实现的详细信息，请参考：
- [GoFrame WebSocket指南](https://goframe.org/docs/web/senior-websocket)
- [Gorilla WebSocket文档](https://github.com/gorilla/websocket)

