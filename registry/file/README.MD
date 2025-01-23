---
title: File Integration
slug: /examples/registry/file
keywords: [registry, file, service discovery, goframe]
description: An example demonstrating File-based service registry integration in GoFrame
hide_title: true
---

# Registry - File Integration

## Description

This example demonstrates how to use File-based service registry with GoFrame applications. It shows how to:
- Register services using file system
- Discover services from files
- Implement simple service discovery
- Build basic distributed systems

## Structure

```
.
├── client/            # Client example
│   └── client.go      # Client implementation with file-based discovery
├── server/            # Server example
│   └── server.go      # Server implementation with file registration
├── go.mod             # Go module file
└── go.sum             # Go module checksums
```

## Features

The example showcases the following features:
1. Service Registration
   - File-based service registration
   - Simple metadata storage
   - Local file system integration
   - Automatic file management

2. Service Discovery
   - File-based service discovery
   - Local service lookup
   - Basic load balancing
   - Simple failover support

3. Protocol Support
   - HTTP services
   - Extensible for other protocols

## Requirements

- [Go](https://golang.org/dl/) 1.22 or higher
- [Git](https://git-scm.com/downloads)
- [GoFrame](https://goframe.org)
- [GoFrame File Registry](https://github.com/gogf/gf/tree/master/contrib/registry/file)

## Prerequisites

No additional infrastructure required. The example uses the local file system for service registration and discovery.

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
1. File-based registry configuration
2. Service registration process
3. Service discovery mechanism
4. Basic health check implementation
5. Simple load balancing strategy

## Notes

- Uses local file system
- Suitable for development and testing
- Simple to set up and use
- No external dependencies
- File-based persistence
- Basic service discovery
- Limited scalability
- Consider security in production
- Handle file permissions properly
- Implement proper error handling
- Monitor file system usage
- Clean up old registrations
