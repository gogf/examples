---
title: Consul
slug: /examples/config/consul
keywords: [配置中心, consul, goframe]
description: GoFrame 框架中 Consul 配置中心的集成示例
hide_title: true
sidebar_position: 1
---

# `Consul` 配置中心示例

## 介绍

本示例展示了如何在 `GoFrame` 应用程序中通过配置管理组件集成 `Consul` 配置中心。


## 目录结构

```text
.
├── boot/           # 启动配置
│   └── boot.go     # Consul 客户端初始化
├── main.go         # 主程序入口
├── go.mod          # Go 模块文件
└── go.sum          # Go 模块校验和
```

## 环境要求

- [Go](https://golang.org/dl/) `1.22` 或更高版本
- [Git](https://git-scm.com/downloads)
- [GoFrame](https://goframe.org)
- [GoFrame Consul Config](https://github.com/gogf/gf/tree/master/contrib/config/consul)

## 使用说明

1. 确保 `Consul` 服务端已启动

2. 配置 `Consul` 连接信息：
   ```yaml
   consul:
     address: "localhost:8500"
     scheme: "http"
     datacenter: "dc1"
   ```

3. 运行示例：
   ```bash
   go run main.go
   ```
