---
title: Polaris
slug: /examples/registry/polaris
keywords: [registry, polaris, service discovery, goframe]
description: Polaris service registry integration in GoFrame
hide_title: true
---

# Registry - Polaris Integration

## Description

This example demonstrates how to integrate Polaris service registry with GoFrame applications. It shows how to:
- Register services with Polaris
- Discover services using Polaris
- Implement service health checks
- Build distributed systems

## Structure

```
.
├── client/           # Client example
│   └── client.go     # Client implementation with Polaris discovery
├── server/           # Server example
│   └── server.go     # Server implementation with Polaris registration
├── go.mod            # Go module file
└── go.sum            # Go module checksums
```

## Features

The example showcases the following features:
1. Service Registration
   - Automatic service registration
   - Health check configuration
   - Metadata management
   - Service routing
   - Rate limiting

2. Service Discovery
   - Dynamic service discovery
   - Load balancing
   - Failover support
   - Circuit breaking
   - Fault tolerance

3. Protocol Support
   - HTTP services
   - Custom protocols
   - Protocol extension

## Requirements

- [Go](https://golang.org/dl/) 1.22 or higher
- [Git](https://git-scm.com/downloads)
- [GoFrame](https://goframe.org)
- [GoFrame Polaris Registry](https://github.com/gogf/gf/tree/master/contrib/registry/polaris)
- [Polaris](https://polarismesh.cn/docs/)

## Prerequisites

1. Running Polaris server:
   ```bash
   # Using docker
   docker run -d \
     --name polaris \
     -p 8090:8090 \
     -p 8091:8091 \
     -p 8093:8093 \
     -p 8761:8761 \
     polarismesh/polaris-server
   ```

## Usage

1. Start the server:
   ```bash
   cd server
   go run server.go
   ```

2. Run the client:
   ```bash
   cd client
   go run client.go
   ```

## Implementation Details

The example demonstrates:
1. Polaris client configuration
2. Service registration process
3. Service discovery mechanism
4. Health check implementation
5. Load balancing strategies
6. Circuit breaker configuration
7. Rate limiting setup

## Notes

- Polaris server must be running
- Default Polaris address: 127.0.0.1:8091
- Services are registered automatically
- Health checks are configured by default
- Load balancing is handled automatically
- Consider security in production
- Handle failover scenarios
- Monitor service health
- Implement proper error handling
- Configure rate limits
- Set up circuit breakers
- Manage service routing
