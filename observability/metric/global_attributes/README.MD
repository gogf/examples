---
title: Metric - Global Attributes Example
slug: /examples/observability/metric/global_attributes
keywords: [metrics, global attributes, prometheus, opentelemetry, goframe]
description: An example demonstrating global metric attributes in GoFrame
---

# Metric - Global Attributes Example

## Description

This example demonstrates how to work with global metric attributes in GoFrame using OpenTelemetry and Prometheus integration. It shows how to:
- Set and manage global attributes
- Apply global attributes across multiple metrics
- Configure attribute scope and patterns
- Combine global attributes with local ones

## Structure

- `go.mod`: The Go module file for dependency management
- `go.sum`: The Go module checksums file
- `main.go`: The main application demonstrating global metric attributes

## Features

The example showcases the following features:
1. Global Attributes
   - Global attribute configuration
   - Attribute inheritance
   - Pattern-based application

2. Metric Types with Global Attributes
   - Counter with global attributes
   - Observable Counter with global attributes
   - Local attribute combination

3. Attribute Management
   - Global attribute scoping
   - Version-based filtering
   - Pattern-based filtering

## Requirements

- [Go](https://golang.org/dl/) 1.22 or higher
- [Git](https://git-scm.com/downloads)
- [GoFrame](https://goframe.org)
- [GoFrame OpenTelemetry Metric](https://github.com/gogf/gf/tree/master/contrib/metric/otelmetric)

## Usage

1. Run the example:
   ```bash
   go run main.go
   ```

2. Access the metrics:
   ```bash
   # Using curl
   curl http://localhost:8000/metrics
   
   # Or open in browser
   http://localhost:8000/metrics
   ```

3. Example metrics output:
   ```
   # HELP goframe_metric_demo_counter This is a simple demo for Counter usage
   goframe_metric_demo_counter{const_attr_1="1",global_attr_1="1"} 11
   
   # HELP goframe_metric_demo_observable_counter This is a simple demo for ObservableCounter usage
   goframe_metric_demo_observable_counter{const_attr_2="2",global_attr_1="1"} 10
   ```

## Implementation Details

The example demonstrates:
1. Setting up global attributes
2. Configuring attribute patterns
3. Version-based attribute filtering
4. Attribute inheritance and precedence
5. Global attribute scoping

## Notes

- Global attributes apply to all matching metrics
- Local attributes take precedence over global ones
- Pattern matching controls attribute application
- Version filtering enables targeted attribute sets
- Consider attribute cardinality impact
- Default port is 8000
- Metrics endpoint is at /metrics
