---
title: Polaris
slug: /examples/config/polaris
keywords: [配置中心, polaris, goframe]
description: GoFrame 框架中 Polaris 配置中心的集成示例
hide_title: true
sidebar_position: 1
---

# Polaris 配置中心示例

## 介绍

本示例展示了如何在 GoFrame 应用程序中通过配置管理组件集成 Polaris 配置中心。

## 目录结构

```
.
├── boot/           # 启动配置
│   └── boot.go     # Polaris 客户端初始化
├── main.go         # 主程序入口
├── go.mod          # Go 模块文件
└── go.sum          # Go 模块校验和
```

## 环境要求

- [Go](https://golang.org/dl/) 1.22 或更高版本
- [Git](https://git-scm.com/downloads)
- [GoFrame](https://goframe.org)
- [GoFrame Polaris Config](https://github.com/gogf/gf/tree/master/contrib/config/polaris)
- [Polaris Server](https://polarismesh.cn/)

## 使用说明

1. 启动 Polaris：
   ```bash
   docker run -d --name polaris \
      -p 8090:8090 -p 8091:8091 -p 8093:8093 -p 9090:9090 -p 9091:9091 \
      polarismesh/polaris-standalone:v1.17.2
   ```

2. 启动服务：
   ```bash
   go run main.go
   ```
