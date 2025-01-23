---
title: Metric - Direct Prometheus Integration Example
slug: /examples/observability/metric/prometheus
keywords: [metrics, prometheus, direct integration, goframe]
description: An example demonstrating direct Prometheus integration in GoFrame without OpenTelemetry
---

# Metric - Direct Prometheus Integration Example

## Description

This example demonstrates how to directly integrate Prometheus metrics in GoFrame without using OpenTelemetry. It shows how to:
- Create Prometheus metrics directly
- Register metrics with Prometheus registry
- Expose metrics via HTTP endpoint
- Update metric values dynamically

## Structure

- `go.mod`: The Go module file for dependency management
- `go.sum`: The Go module checksums file
- `prometheus.go`: The main application demonstrating direct Prometheus integration

## Features

The example showcases the following features:
1. Metric Types
   - Counter: Monotonically increasing value
   - Gauge: Value that can go up and down

2. Metric Operations
   - Metric registration
   - Value updates
   - HTTP exposure
   - Random value generation

3. HTTP Endpoints
   - `/`: Triggers metric updates
   - `/metrics`: Exposes Prometheus metrics

## Requirements

- [Go](https://golang.org/dl/) 1.22 or higher
- [Git](https://git-scm.com/downloads)
- [GoFrame](https://goframe.org)
- [GoFrame OpenTelemetry Metric](https://github.com/gogf/gf/tree/master/contrib/metric/otelmetric)

## Usage

1. Run the example:
   ```bash
   go run prometheus.go
   ```

2. Test the application:
   ```bash
   # Generate some metric values
   curl http://localhost:8000/
   
   # View metrics
   curl http://localhost:8000/metrics
   ```

3. Example metrics output:
   ```
   # HELP demo_counter A demo counter.
   # TYPE demo_counter counter
   demo_counter 1
   
   # HELP demo_gauge A demo gauge.
   # TYPE demo_gauge gauge
   demo_gauge 42
   ```

## Implementation Details

The example demonstrates:
1. Direct Prometheus metric creation
2. Manual metric registration
3. HTTP server configuration
4. Dynamic value updates
5. Metric exposure

## Notes

- Uses native Prometheus client
- No OpenTelemetry dependency
- Simpler but less feature-rich
- Manual metric management
- Good for basic use cases
- Default port is 8000
- Metrics endpoint at /metrics
- Consider security when exposing metrics
- Random values for demonstration
- Counter only increases
- Gauge can increase or decrease
