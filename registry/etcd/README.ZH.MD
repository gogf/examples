---
title: Etcd
slug: /examples/registry/etcd
keywords: [注册中心, etcd, 服务发现, goframe]
description: GoFrame框架中的Etcd服务注册与发现集成
hide_title: true
sidebar_position: 1
---

# 注册中心 - Etcd集成

## 简介

本示例演示了如何在GoFrame应用程序中集成Etcd服务注册中心。主要展示：
- 使用Etcd注册服务
- 使用Etcd发现服务
- 实现服务健康检查
- 构建分布式系统

## 目录结构

```
.
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

## 功能特性

本示例展示了以下功能：

1. 服务注册
   - 自动服务注册
   - 基于TTL的健康检查
   - 元数据管理
   - 租约管理

2. 服务发现
   - 动态服务发现
   - 负载均衡
   - 故障转移支持
   - Watch机制监控

3. 协议支持
   - HTTP服务注册与发现
   - gRPC服务注册与发现

## 环境要求

- Go 1.22 或更高版本
- Docker (用于运行Etcd)
- Etcd 3.4+

## 使用说明

1. 启动Etcd服务：
   ```bash
   docker run -d --name etcd -p 2379:2379 -e ALLOW_NONE_AUTHENTICATION=yes bitnami/etcd:3.4.24
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
   - 通过`gsvc.SetRegistry`设置Etcd注册中心
   - 客户端使用`g.Client()`自动发现并访问服务
   - 支持自动的健康检查和租约续期

2. gRPC服务实现
   - 服务器使用`grpcx.Server`创建gRPC服务
   - 通过`grpcx.Resolver.Register`设置Etcd注册中心
   - 客户端使用`grpcx.Client`自动发现并访问gRPC服务
   - 支持服务健康检查和自动重连

## 注意事项

1. 确保Etcd服务正常运行在`127.0.0.1:2379`
2. 服务注册使用TTL机制确保服务健康状态
3. 客户端支持Watch机制，能及时感知服务变化
4. 默认提供负载均衡和故障转移能力
