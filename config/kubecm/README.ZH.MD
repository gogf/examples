---
title: Kubernetes ConfigMap
slug: /examples/config/kubecm
keywords: [配置中心, kubernetes, configmap, goframe]
description: GoFrame 框架中 Kubernetes ConfigMap 的集成示例
hide_title: true
sidebar_position: 9
---

# `Kubernetes ConfigMap` 配置示例

## 介绍

本示例展示了如何在 `GoFrame` 应用程序中通过配置管理组件集成 `Kubernetes ConfigMap`。


## 目录结构

```text
.
├── boot_in_pod/     # Pod内部署的启动配置
│   └── boot.go      # Pod内客户端初始化
├── boot_out_pod/    # Pod外部署的启动配置
│   └── boot.go      # Pod外客户端初始化
├── main.go          # 主程序入口
├── go.mod           # Go 模块文件
└── go.sum           # Go 模块校验和
```


## 环境要求

- [Go](https://golang.org/dl/) 1.22 或更高版本
- [Git](https://git-scm.com/downloads)
- [GoFrame](https://goframe.org)
- [GoFrame Kubernetes ConfigMap Config](https://github.com/gogf/gf/tree/master/contrib/config/kubecm)




