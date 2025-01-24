---
title: 回调处理
slug: /examples/observability/metric/callback
keywords: [指标, 回调, prometheus, opentelemetry, goframe]
description: GoFrame 中基于回调的指标收集实现
hide_title: true
sidebar_position: 1
---

# 指标收集 - 回调处理

## 简介

本示例演示了如何在 GoFrame 中使用 OpenTelemetry 和 Prometheus 集成来实现基于回调的指标收集。主要展示：
- 创建带有回调函数的指标
- 通过回调自动更新指标值
- 配置指标属性
- 导出 Prometheus 格式的指标

## 目录结构

```
.
├── go.mod            # Go 模块文件
├── go.sum            # Go 模块校验和
└── main.go           # 主程序，演示回调式指标使用
```

## 功能特性

本示例展示了以下功能：

1. 基于回调的指标
   - 自动值更新
   - 自定义回调函数
   - 无需手动更新
   - 适用于动态值监控

2. 常规指标
   - 手动值更新
   - 直接控制值
   - 与回调方式对比
   - 适用于事件计数

3. 指标配置
   - 自定义属性
   - 帮助文本
   - 单位设置
   - 常量标签

## 环境要求

- Go 1.22 或更高版本
- GoFrame 框架
- GoFrame OpenTelemetry Metric 扩展

## 使用说明

1. 运行示例：
   ```bash
   go run main.go
   ```

2. 访问指标：
   ```bash
   # 使用 curl
   curl http://localhost:8000/metrics
   
   # 或在浏览器中打开
   http://localhost:8000/metrics
   ```

3. 指标输出示例：
   ```
   # HELP goframe_metric_demo_counter 这是 Counter 使用的简单演示
   goframe_metric_demo_counter{const_attr_1="1"} 11
   
   # HELP goframe_metric_demo_observable_counter 这是 ObservableCounter 使用的简单演示
   goframe_metric_demo_observable_counter{const_attr_3="3"} 10
   ```

## 实现说明

1. 回调式指标实现
   ```go
   observableCounter := meter.MustObservableCounter(
       "goframe.metric.demo.observable_counter",
       gmetric.MetricOption{
           Help: "回调计数器示例",
           Unit: "%",
           Callback: func(ctx context.Context, obs gmetric.MetricObserver) error {
               obs.Observe(10)  // 自动设置值为 10
               return nil
           },
       },
   )
   ```

2. 常规指标实现
   ```go
   counter := meter.MustCounter(
       "goframe.metric.demo.counter",
       gmetric.MetricOption{
           Help: "普通计数器示例",
           Unit: "%",
       },
   )
   counter.Inc(ctx)      // 手动增加 1
   counter.Add(ctx, 10)  // 手动增加指定值
   ```
