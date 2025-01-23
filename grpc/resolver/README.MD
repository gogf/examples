---
title: gRPC - Service Resolver
slug: /examples/grpc/resolver
keywords: [grpc, resolver, etcd, goframe]
description: An example demonstrating gRPC service resolver with etcd in GoFrame
---

# gRPC - Service Resolver

## Description

This example demonstrates how to use service resolver with etcd in gRPC services using GoFrame. It shows how to:
- Configure service resolver with etcd
- Register services to etcd
- Discover services using resolver
- Handle service updates

## Structure

```
.
├── client/              # Client example
│   └── client.go        # Client with service resolver
├── controller/          # Service controllers
│   └── helloworld.go    # Hello service implementation
├── protobuf/            # Protocol buffer definitions
│   └── helloworld.proto # Service and message definitions
├── server/              # Server example
│   └── server.go        # Server with service registration
├── go.mod               # Go module file
└── go.sum               # Go module checksums
```

## Features

The example showcases the following features:
1. Service Resolution
   - etcd integration
   - Service registration
   - Service discovery
   - Update handling

2. Client Usage
   - Resolver configuration
   - Service discovery
   - Connection management
   - Error handling

3. Server Features
   - Service registration
   - etcd integration
   - Health checking
   - Metadata management

## Requirements

- [Go](https://golang.org/dl/) 1.22 or higher
- [Git](https://git-scm.com/downloads)
- [GoFrame](https://goframe.org)
- [gRPC](https://grpc.io/docs/languages/go/quickstart/)

## Prerequisites

1. Running etcd instance:
   ```bash
   # Using Docker
   docker run -d --name etcd -p 2379:2379 -e ALLOW_NONE_AUTHENTICATION=yes bitnami/etcd:3.4.24
   ```

2. Protocol buffer compiler:
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
1. Service registration with etcd
2. Service discovery using resolver
