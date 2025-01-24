---
title: 基础用法
slug: /examples/observability/metric/basic
keywords: [指标, 基础, prometheus, opentelemetry, goframe]
description: 在GoFrame中演示各种指标类型及其使用的基础示例
hide_title: true
sidebar_position: 1
---

# 指标收集 - 基础用法

## 简介

本示例演示了如何在GoFrame中使用OpenTelemetry和Prometheus集成来实现基础的指标收集功能。主要展示：
- 创建和使用不同类型的指标
- 配置指标属性
- 导出Prometheus格式的指标
- 设置指标暴露端点

## 目录结构

```
.
├── go.mod            # Go模块文件
├── go.sum            # Go模块校验和
└── main.go           # 主程序，演示指标使用
```

## 功能特性

本示例展示了以下指标类型：

1. Counter (计数器)
   - 累积型测量
   - 只能增加
   - 用于事件计数
   - 示例：请求总数、字节数

2. UpDownCounter (上下计数器)
   - 双向计数器
   - 可增可减
   - 用于测量变化量
   - 示例：队列长度、连接数

3. Histogram (直方图)
   - 测量值分布
   - 可配置的桶
   - 用于延迟/大小测量
   - 示例：请求延迟、响应大小

4. Observable指标
   - ObservableCounter
   - ObservableUpDownCounter
   - ObservableGauge
   - 通过回调更新
   - 示例：CPU使用率、内存使用

## 环境要求

- Go 1.22 或更高版本
- GoFrame框架
- GoFrame OpenTelemetry Metric扩展

## 使用说明

1. 运行示例：
   ```bash
   go run main.go
   ```

2. 访问指标：
   ```bash
   # 使用curl
   curl http://localhost:8000/metrics
   
   # 或在浏览器中打开
   http://localhost:8000/metrics
   ```

3. 指标输出示例：
   ```
   # HELP goframe_metric_demo_counter 这是Counter使用的简单演示
   goframe_metric_demo_counter{const_attr_1="1"} 11
   
   # HELP goframe_metric_demo_histogram 这是histogram使用的简单演示
   goframe_metric_demo_histogram_bucket{const_attr_3="3",le="0"} 0
   goframe_metric_demo_histogram_bucket{const_attr_3="3",le="10"} 1
   ...
   ```

## 实现说明

1. Counter实现
   ```go
   counter := meter.MustCounter(
       "goframe.metric.demo.counter",
       gmetric.MetricOption{
           Help: "计数器使用示例",
           Unit: "bytes",
       },
   )
   counter.Inc(ctx)      // 增加1
   counter.Add(ctx, 10)  // 增加指定值
   ```

2. UpDownCounter实现
   ```go
   upDownCounter := meter.MustUpDownCounter(
       "goframe.metric.demo.updown_counter",
       gmetric.MetricOption{
           Help: "上下计数器使用示例",
           Unit: "%",
       },
   )
   upDownCounter.Inc(ctx)   // 增加1
   upDownCounter.Dec(ctx)   // 减少1
   ```

3. Histogram实现
   ```go
   histogram := meter.MustHistogram(
       "goframe.metric.demo.histogram",
       gmetric.MetricOption{
           Help: "直方图使用示例",
           Unit: "ms",
           Buckets: []float64{0, 10, 20, 50, 100},
       },
   )
   histogram.Record(30)  // 记录测量值
   ```

4. Observable指标实现
   ```go
   observableGauge := meter.MustObservableGauge(
       "goframe.metric.demo.observable_gauge",
       gmetric.MetricOption{
           Help: "可观察量表使用示例",
           Unit: "%",
       },
   )
   meter.MustRegisterCallback(func(ctx context.Context, obs gmetric.Observer) error {
       obs.Observe(observableGauge, 30)
       return nil
   }, observableGauge)
   ```
