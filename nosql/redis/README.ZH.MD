---
title: Redis 
slug: /examples/nosql/redis
keywords: [nosql, redis, cache, database, goframe]
description: 在 GoFrame 中使用 Redis 的示例
hide_title: true
sidebar_position: 1
---

# GoFrame Redis 示例

此示例演示了如何在 `GoFrame` 框架中使用 `Redis`。

## 概述

本示例展示了：
1. 如何使用 `YAML` 配置文件配置 `Redis` 连接
2. 如何创建 `Redis` 客户端
3. 基本的 `Redis` 操作（`SET/GET`）

## 环境要求

- `Go 1.15` 或更高版本
- `Redis` 服务器
- `GoFrame v2`

## 代码结构

- `main.go`: 包含主要逻辑和 `Redis` 客户端初始化
- `config.yaml`: `Redis` 配置文件

## 配置说明

`Redis` 配置存储在 `config.yaml` 文件中：

```yaml
redis:
  address: "127.0.0.1:6379"
  password:
```

您可以根据自己的 `Redis` 服务器配置修改这些设置。

## 使用 Docker 运行 Redis

如果您本地没有安装 `Redis`，可以使用 `Docker` 快速启动一个 `Redis` 实例：

```bash
# 运行 Redis 容器
docker run --name redis-test -p 6379:6379 -d redis:latest

# 验证容器是否正在运行
docker ps

# 如果需要停止容器
docker stop redis-test

# 如果需要删除容器
docker rm redis-test
```

如果需要带密码认证的 `Redis`：

```bash
# 运行带密码的 Redis
docker run --name redis-test -p 6379:6379 -d redis:latest redis-server --requirepass your_password

# 记得相应地更新 config.yaml：
# redis:
#   address: "127.0.0.1:6379"
#   password: "your_password"
```

## 运行示例

1. 确保 `Redis` 服务器正在运行
2. 根据需要更新 `config.yaml`
3. 运行示例：

```bash
go run main.go
```




## 更多参考

更多 `Redis` 的使用方法请参考第三方组件 [`github.com/redis/go-redis`](https://github.com/redis/go-redis)。
