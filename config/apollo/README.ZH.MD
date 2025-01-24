---
title: Apollo
slug: /examples/config/apollo
keywords: [配置中心, apollo, goframe]
description: GoFrame 框架中 Apollo 配置中心的集成示例
hide_title: true
sidebar_position: 1
---

# `Apollo` 配置中心示例

## 介绍

本示例展示了如何在 `GoFrame` 应用程序中通过配置管理组件集成 `Apollo` 配置中心。


## 目录结构

```text
.
├── boot/           # 启动配置
│   └── boot.go     # Apollo 客户端初始化
├── main.go         # 主程序入口
├── go.mod          # Go 模块文件
└── go.sum          # Go 模块校验和
```

## 环境要求

- [Go](https://golang.org/dl/) 1.22 或更高版本
- [Git](https://git-scm.com/downloads)
- [GoFrame](https://goframe.org)
- [GoFrame Apollo Config](https://github.com/gogf/gf/tree/master/contrib/config/apollo)

## 使用说明

1. 确保 `Apollo` 服务端已启动

2. 配置 `Apollo` 连接信息：
   ```yaml
   apollo:
     appId: "your-app-id"
     cluster: "default"
     namespaceName: "application"
     ip: "localhost:8080"
   ```

3. 运行示例：
   ```bash
   go run main.go
   ```
