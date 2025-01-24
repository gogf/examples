---
title: Consul
slug: /examples/registry/consul
keywords: [注册中心, consul, 服务发现, goframe]
description: GoFrame框架中的Consul服务注册与发现集成
hide_title: true
sidebar_position: 1
---

# 注册中心 - Consul集成

## 简介

本示例演示了如何在GoFrame应用程序中集成Consul服务注册中心。主要展示：
- 使用Consul注册服务
- 使用Consul发现服务
- 实现服务健康检查
- 构建分布式系统

## 目录结构

```
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

- Go 1.22 或更高版本
- Docker (用于运行Consul)
- Consul

## 使用说明

1. 启动Consul服务：
   ```bash
   cd docker-compose
   docker-compose up -d
   ```

2. HTTP服务示例：
   ```bash
   # 启动HTTP服务器
   cd http/server
   go run server.go

   # 运行HTTP客户端
   cd http/client
   go run client.go
   ```

3. gRPC服务示例：
   ```bash
   # 启动gRPC服务器
   cd grpc/server
   go run server.go

   # 运行gRPC客户端
   cd grpc/client
   go run client.go
   ```

## 实现说明

1. HTTP服务实现
   - 服务器使用`g.Server`创建HTTP服务
   - 通过`gsvc.SetRegistry`设置Consul注册中心
   - 客户端使用`g.Client()`自动发现并访问服务

2. gRPC服务实现
   - 服务器使用`grpcx.Server`创建gRPC服务
   - 通过`grpcx.Resolver.Register`设置Consul注册中心
   - 客户端使用`grpcx.Client`自动发现并访问gRPC服务

## 注意事项

1. 确保Consul服务正常运行在`127.0.0.1:8500`
2. 服务注册时会自动添加健康检查
3. 客户端会自动使用服务发现进行负载均衡
