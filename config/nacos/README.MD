---
title: Nacos
slug: /examples/config/nacos
keywords: [config, nacos, goframe]
description: Example demonstrating Nacos configuration center integration with GoFrame
hide_title: true
---

## Nacos Configuration Center Example

## Description

This directory contains an example demonstrating how to integrate Nacos configuration center with GoFrame applications. It shows:

1. Nacos Client Configuration
   - Nacos client setup and initialization
   - Configuration adapter implementation
   - Error handling and logging

2. Configuration Management
   - Configuration loading and parsing
   - Dynamic configuration updates
   - Configuration value retrieval

## Directory Structure

```
.
├── boot/           # Bootstrap configuration
│   └── boot.go     # Nacos client initialization
├── main.go         # Main application entry
├── go.mod          # Go module file
└── go.sum          # Go module checksums
```

## Requirements

- [Go](https://golang.org/dl/) 1.22 or higher
- [Git](https://git-scm.com/downloads)
- [GoFrame](https://goframe.org)
- [GoFrame Nacos Config](https://github.com/gogf/gf/tree/master/contrib/config/nacos)
- [Nacos Server](https://nacos.io/)

## Features

The example showcases the following features:

1. Nacos Integration
   - Client configuration
   - Server connection management
   - Configuration namespace
   - Error handling

2. Configuration Management
   - Configuration loading
   - Value retrieval
   - Type conversion
   - Default values

3. Dynamic Updates
   - Configuration watching
   - Change notification
   - Hot reload support

## Configuration

### Nacos Server
1. Server Configuration:
   - Default port: 8848
   - Default address: localhost
   - Default protocol: HTTP

2. Authentication:
   - Username/Password authentication
   - Namespace isolation
   - Group level access control

### Client Configuration
1. Basic Settings:
   - Server address and port
   - Cache directory
   - Log directory
   - Data ID and Group

2. Advanced Options:
   - Namespace customization
   - Cache management
   - Logging configuration
   - Retry settings

## Usage

1. Start Nacos Server:
   ```bash
   # Start Nacos server using Docker
   docker run -d --name nacos \
      -p 8848:8848 -p 9848:9848 \
      -e MODE=standalone \
      nacos/nacos-server:latest
   ```

2. Configure Nacos:
   - Access Nacos console at http://localhost:8848/nacos
   - Create configuration with Data ID: config.toml
   - Add configuration content in TOML format
   - Publish configuration

3. Run Example:
   ```bash
   go run main.go
   ```

## Implementation Details

1. Client Setup (`boot/boot.go`):
   - Server connection configuration
   - Client settings
   - Cache and log directory setup
   - Configuration parameter setup

2. Configuration Access (`main.go`):
   - Configuration availability check
   - Bulk configuration retrieval
   - Single value access
   - Error handling

## Notes

- The example uses Nacos's configuration center features
- Configuration changes in Nacos are automatically reflected
- Local cache is used for better performance
- Proper logging is configured for debugging
- Configuration can be organized by Data ID and Group
- Supports dynamic configuration updates
