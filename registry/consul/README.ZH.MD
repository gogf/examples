---
title: Consul
slug: /examples/registry/consul
keywords: [注册中心, consul, 服务发现, goframe]
description: GoFrame框架中的Consul服务注册与发现集成
hide_title: true
sidebar_position: 1
---

# 注册中心 - `Consul` 集成

## 简介

本示例演示了如何在 `GoFrame` 应用程序中集成 `Consul` 服务注册中心。主要展示：
- 使用 `Consul` 注册服务
- 使用 `Consul` 发现服务
- 实现服务健康检查
- 构建分布式系统

## 目录结构

```text
.
├── docker-compose/      # 运行Consul的Docker配置文件
├── grpc/                # gRPC服务示例
│   ├── client/          # gRPC客户端实现
│   ├── controller/      # gRPC服务控制器
│   ├── protobuf/        # Protocol buffer定义
│   └── server/          # gRPC服务器实现
│       ├── main.go      # 服务器启动代码
│       └── config.yaml  # 服务器配置
├── http/                # HTTP服务示例
│   ├── client/          # HTTP客户端实现
│   └── server/          # HTTP服务器实现
├── go.mod               # Go模块文件
└── go.sum               # Go模块校验和
```


## 环境要求

- `Go` 1.22 或更高版本
- `Docker` (用于运行 `Consul`)
- `Consul`

## 使用说明

1. 启动 `Consul` 服务：
   ```bash
   cd docker-compose
   docker-compose up -d
   ```

2. `HTTP` 服务示例：
   ```bash
   # 启动HTTP服务器
   cd http/server
   go run server.go

   # 运行HTTP客户端
   cd http/client
   go run client.go
   ```

3. `gRPC` 服务示例：
   ```bash
   # 启动gRPC服务器
   cd grpc/server
   go run server.go

   # 运行gRPC客户端
   cd grpc/client
   go run client.go
   ```

## 实现说明

1. `HTTP` 服务实现
   - 服务器使用 `g.Server` 创建 `HTTP` 服务
   - 通过 `gsvc.SetRegistry` 设置 `Consul` 注册中心
   - 客户端使用 `g.Client()` 自动发现并访问服务

2. `gRPC` 服务实现
   - 服务器使用 `grpc.Server` 创建 `gRPC` 服务
   - 通过 `gsvc.SetRegistry` 设置 `Consul` 注册中心
   - 客户端使用 `grpc.Client` 自动发现并访问服务

## 注意事项

1. 确保 `Consul` 服务正常运行在 `127.0.0.1:8500`
2. 服务注册时会自动添加健康检查
3. 客户端会自动使用服务发现进行负载均衡
