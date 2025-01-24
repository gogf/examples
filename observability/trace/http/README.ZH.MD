---
title: HTTP服务
slug: /examples/observability/trace/http
keywords: [链路跟踪, http, goframe]
description: GoFrame中HTTP服务的分布式跟踪实现
hide_title: true
sidebar_position: 1
---

# 链路跟踪 - HTTP服务

## 简介

本示例演示了如何在GoFrame中实现HTTP服务的分布式跟踪。主要展示：
- 配置HTTP服务的跟踪
- 跟踪HTTP请求和响应
- 传播跟踪上下文
- 可视化分布式跟踪

## 环境要求

- Go 1.22 或更高版本
- GoFrame框架
- GoFrame OpenTelemetry跟踪

## 目录结构

```
.
├── client/          # 客户端示例
│   └── client.go    # 带跟踪的客户端
├── server/          # 服务端示例
│   └── server.go    # 带跟踪的服务端
├── go.mod          # Go模块文件
└── go.sum          # Go模块校验和
```


## 前置条件

1. 运行Jaeger实例：
   ```bash
   docker run --rm --name jaeger \
   -e COLLECTOR_ZIPKIN_HOST_PORT=:9411 \
   -p 6831:6831/udp \
   -p 6832:6832/udp \
   -p 5778:5778 \
   -p 16686:16686 \
   -p 4317:4317 \
   -p 4318:4318 \
   -p 14250:14250 \
   -p 14268:14268 \
   -p 14269:14269 \
   -p 9411:9411 \
   jaegertracing/all-in-one:1.55
   ```

## 使用说明

1. 启动服务端：
   ```bash
   cd server
   go run server.go
   ```

2. 运行客户端：
   ```bash
   cd client
   go run client.go
   ```

3. 查看跟踪：
   在浏览器中打开 http://localhost:16686 查看Jaeger UI中的跟踪信息。

## API接口

服务器提供以下HTTP接口：

1. Hello World
   ```
   GET /hello
   响应: "Hello World"
   ```

## 实现说明

1. 服务端实现
   ```go
   // 初始化跟踪
   shutdown, err := otlphttp.Init(serviceName, endpoint, path)

   // 创建HTTP服务器
   s := g.Server()
   s.Group("/", func(group *ghttp.RouterGroup) {
       group.GET("/hello", HelloHandler)
   })

   // 处理请求并跟踪
   func HelloHandler(r *ghttp.Request) {
       ctx, span := gtrace.NewSpan(r.Context(), "HelloHandler")
       defer span.End()

       value := gtrace.GetBaggageVar(ctx, "name").String()
       r.Response.Write("hello:", value)
   }
   ```

2. 客户端实现
   ```go
   // 创建新的跟踪span
   ctx, span := gtrace.NewSpan(gctx.New(), "StartRequests")
   defer span.End()

   // 设置跟踪的baggage值
   ctx = gtrace.SetBaggageValue(ctx, "name", "GoFrame")

   // 发送HTTP请求
   response, err := g.Client().Get(ctx, "http://127.0.0.1:8199/hello")
   ```

