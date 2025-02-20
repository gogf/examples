---
title: 多进程示例
slug: /examples/observability/trace/processes
keywords: [链路跟踪, 多进程, gcmd, gproc, goframe]
description: 使用GoFrame不同进程管理方式实现的多进程分布式链路跟踪示例
hide_title: true
sidebar_position: 1
---

# 多进程链路跟踪示例

## 简介

本目录包含了使用GoFrame不同进程管理方式实现多进程分布式链路跟踪的示例。包括：

1. 命令行进程管理 (`gcmd/`)
   - 使用`gcmd`包进行进程管理
   - 演示基于命令行的进程创建
   - 展示父子进程间的链路跟踪上下文传播

2. 程序化进程管理 (`gproc/`)
   - 使用`gproc`包进行进程管理
   - 演示程序化的进程创建
   - 展示进程层级中的链路跟踪传播

## 目录结构

```
.
├── gcmd/           # 命令行进程管理示例
│   ├── main.go     # 主进程实现
│   └── sub/        # 子进程实现
│       └── sub.go  # 子进程代码
├── gproc/          # 进程管理示例
│   ├── main.go     # 主进程实现
│   └── sub/        # 子进程实现
│       └── sub.go  # 子进程代码
├── go.mod          # Go模块文件
└── go.sum          # Go模块校验和
```

## 环境要求

- Go `1.22` 或更高版本
- GoFrame框架
- GoFrame链路跟踪支持

## 功能特性

本示例展示了以下功能：

1. 进程管理
   ```go
   // gcmd方式
   Main = &gcmd.Command{
       Name:  "main",
       Brief: "main process",
       Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
           return gproc.ShellRun(ctx, `go run sub/sub.go`)
       },
   }

   // gproc方式
   if err := gproc.ShellRun(ctx, `go run sub/sub.go`); err != nil {
       panic(err)
   }
   ```

2. 链路跟踪上下文传播
   ```go
   // 主进程
   ctx := gctx.GetInitCtx()
   g.Log().Debug(ctx, `this is main process`)

   // 子进程
   ctx := gctx.GetInitCtx()
   g.Log().Debug(ctx, `this is sub process`)
   ```

3. 日志和调试
   - 进程标识
   - 调试日志
   - 错误报告

## 管理方式对比

### 命令行管理 (gcmd/)
1. 特点：
   - 命令行界面
   - 结构化命令处理
   - 内置帮助和文档
   - 命令层级支持

2. 使用场景：
   - CLI应用程序
   - 命令驱动工具
   - 交互式应用

### 程序化管理 (gproc/)
1. 特点：
   - 程序化进程控制
   - 直接进程操作
   - Shell命令执行
   - 进程同步

2. 使用场景：
   - 后台进程
   - 服务管理
   - 进程自动化

## 使用说明

### 命令行示例
1. 进入gcmd示例目录：
   ```bash
   cd gcmd
   ```

2. 运行示例：
   ```bash
   go run main.go
   ```

### 程序化管理示例
1. 进入gproc示例目录：
   ```bash
   cd gproc
   ```

2. 运行示例：
   ```bash
   go run main.go
   ```

## 实现说明

两个示例都演示了：

1. 进程创建
   - 主进程初始化
   - 子进程启动
   - 进程环境设置

2. 上下文管理
   ```go
   // 创建和初始化上下文
   ctx := gctx.GetInitCtx()

   // 在进程间传播上下文
   g.Log().Debug(ctx, `process message`)
   ```

3. 错误处理
   ```go
   // 进程执行错误处理
   if err := gproc.ShellRun(ctx, command); err != nil {
       panic(err)
   }

   // 命令执行错误处理
   if err := cmd.Run(ctx); err != nil {
       return err
   }
   ```
