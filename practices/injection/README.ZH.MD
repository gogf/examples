---
title: 依赖注入
slug: /examples/practices/injection
keywords: [依赖注入, 测试, goframe, di]
description: 使用GoFrame框架和do包实现的依赖注入示例
hide_title: true
---

# 依赖注入示例

## 简介

本示例展示了如何在 `GoFrame` 应用程序中实现依赖注入。主要演示：

1. 基础依赖注入设置
2. 使用DI的服务层实现
3. 集成gRPC的控制器层
4. 使用模拟依赖的单元测试

实现重点在于使核心业务逻辑易于测试：
- 使用依赖注入分离关注点
- 为服务依赖提供清晰的接口
- 使依赖易于在测试中模拟

## 环境要求

- [Go](https://golang.org/dl/) `1.22` 或更高版本
- [Git](https://git-scm.com/downloads)
- [`GoFrame`](https://goframe.org)
- [`do`](https://github.com/samber/do)
- [MongoDB](https://www.mongodb.com)
- [Redis](https://redis.io)

## 目录结构

```text
injection/
├── app/
│   ├── gateway/          # API网关服务
│   │   ├── api/         # API定义
│   │   ├── internal/    # 内部实现
│   │   │   ├── cmd/    # 命令行工具
│   │   │   ├── controller/  # 控制器
│   │   │   ├── model/  # 数据模型
│   │   │   └── service/  # 业务逻辑
│   │   └── manifest/   # 配置文件
│   └── user/           # 用户服务
│       ├── api/        # API定义
│       │   ├── entity/ # 实体定义
│       │   └── user/   # 用户API proto
│       ├── internal/   # 内部实现
│       │   ├── cmd/    # 命令行工具
│       │   ├── controller/  # 使用DI的控制器
│       │   ├── dao/    # 数据访问对象
│       │   ├── model/  # 数据模型
│       │   └── service/  # 使用DI的业务逻辑
│       └── manifest/   # 配置文件
├── hack/               # 开发工具
└── utility/           # 通用工具
    ├── injection/     # DI工具
    └── mongohelper/   # MongoDB辅助工具
```

## 功能特性

- 依赖注入使用
- MongoDB和Redis集成
- gRPC服务实现
- 完整的单元测试
- 清理资源的关闭处理
- 命名依赖支持

## 安装设置

1. 克隆仓库：
    ```bash
    git clone https://github.com/gogf/examples.git
    cd examples/practices/injection
    ```

2. 安装依赖：
    ```bash
    go mod tidy
    ```

3. 使用Docker启动所需服务：
    ```bash
    # 启动MongoDB
    docker run -d --name mongo -p 27017:27017 mongo:latest

    # 启动Redis
    docker run -d --name redis -p 6379:6379 redis:latest
    ```

## 使用方法

1. 运行应用：
   ```bash
   go run main.go
   ```

2. 运行测试：
   ```bash
   go test ./...
   ```

## 实现细节

### 依赖注入设置
- 使用 `do` 包进行依赖管理
- 支持命名和未命名依赖
- 提供常用操作的辅助函数

### 服务层
- 清晰的关注点分离
- 基于接口的设计
- 易于使用模拟实现进行测试

### 控制器层
- 支持DI的gRPC集成
- 清晰的错误处理
- 合理的资源管理

### 测试
- 完整的单元测试
- 依赖模拟
- 清晰的测试设置和清理

## 注意事项

- 始终使用接口定义依赖，使其可模拟
- 在 `Shutdown` 方法中正确清理资源
- 当需要同类型的多个实例时使用命名依赖
- 在测试中考虑使用外部服务的模拟实现

