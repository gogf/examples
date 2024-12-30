---
title: gRPC - Basic Usage
keywords: [grpc, basic, goframe]
description: An example demonstrating basic gRPC usage in GoFrame
---

# gRPC - Basic Usage

## Description

This example demonstrates the basic usage of gRPC in GoFrame applications. It shows how to:
- Create gRPC servers and clients
- Define and use protocol buffers
- Implement service handlers
- Make RPC calls

## Structure

```
.
├── client/         # Client example
│   └── client.go   # Client implementation
├── controller/     # Service controllers
│   └── hello.go    # Hello service implementation
├── protobuf/       # Protocol buffer definitions
├── server/         # Server example
│   ├── config.yaml # Server configuration
│   └── server.go   # Server implementation
├── go.mod          # Go module file
└── go.sum          # Go module checksums
```

## Features

The example showcases the following features:
1. Service Implementation
   - Protocol buffer definitions
   - Service handlers
   - Request/response handling
   - Error management

2. Client Usage
   - Connection management
   - RPC calls
   - Error handling
   - Context usage

3. Protocol Support
   - gRPC services
   - Protocol buffer messages
   - Custom metadata
   - Stream processing

## Prerequisites

1. Protocol buffer compiler installed:
   ```bash
   # For macOS
   brew install protobuf
   
   # Install protoc-gen-go and protoc-gen-go-grpc
   go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
   go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
   ```

## Usage

1. Generate protocol buffer code:
   ```bash
   cd protobuf
   protoc --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. *.proto
   ```

2. Start the server:
   ```bash
   cd server
   go run server.go
   ```

3. Run the client:
   ```bash
   cd client
   go run client.go
   ```

## Implementation Details

The example demonstrates:
1. Service definition
2. Client/server implementation
