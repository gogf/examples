---
slug: '/examples/nosql/mongodb'
title: 'MongoDB'
hide_title: true
sidebar_position: 2
---

# GoFrame MongoDB 示例

此示例演示了如何在 `GoFrame` 框架中使用 `MongoDB`。

## 概述

本示例展示了：
1. 如何使用 `YAML` 配置文件配置 `MongoDB` 连接
2. 如何创建 `MongoDB` 客户端
3. 基本的 `MongoDB` 操作（`INSERT`、`FIND`、`UPDATE`）
4. 如何使用 `BSON` 进行文档操作

## 环境要求

- `Go 1.15` 或更高版本
- `MongoDB` 服务器
- `GoFrame v2`

## 配置说明

`MongoDB` 配置存储在 `config.yaml` 文件中：

```yaml
mongo:
  database: "user"
  address: "mongodb://127.0.0.1:27017/test?retryWrites=true"
```

您可以根据自己的 `MongoDB` 服务器配置修改这些设置。

## 使用 Docker 运行 MongoDB

如果您本地没有安装 `MongoDB`，可以使用 `Docker` 快速启动一个 `MongoDB` 实例：

```bash
# 运行 MongoDB 容器
docker run --name mongo-test -p 27017:27017 -d mongo:latest

# 验证容器是否正在运行
docker ps

# 如果需要停止容器
docker stop mongo-test

# 如果需要删除容器
docker rm mongo-test
```

如果需要带认证的 `MongoDB`：

```bash
# 运行带认证的 MongoDB
docker run --name mongo-test -p 27017:27017 -e MONGO_INITDB_ROOT_USERNAME=admin -e MONGO_INITDB_ROOT_PASSWORD=password -d mongo:latest

# 记得相应地更新 config.yaml：
# mongo:
#   database: "user"
#   address: "mongodb://admin:password@127.0.0.1:27017/test?retryWrites=true"
```

## 运行示例

1. 确保 `MongoDB` 服务器正在运行
2. 根据需要更新 `config.yaml`
3. 运行示例：

```bash
go run main.go
```

## 代码结构

- `main.go`: 包含主要逻辑和 `MongoDB` 客户端初始化
- `config.yaml`: `MongoDB` 配置文件

## 功能特点

1. `MongoDB` 客户端初始化及错误处理
2. 使用 `GoFrame` 的配置系统进行配置管理
3. 基本 `MongoDB` 操作演示
4. 完善的错误处理和日志记录
5. 使用 `BSON` 标签进行文档映射

## 更多参考

更多 `MongoDB` 的使用方法请参考官方 `MongoDB` Go 驱动 [`github.com/mongodb/mongo-go-driver`](https://github.com/mongodb/mongo-go-driver)。

## 注意事项

- 示例使用官方 `MongoDB` Go 驱动
- 所有操作都使用 context 以实现正确的取消和超时处理
- 代码包含适当的错误处理和日志记录
