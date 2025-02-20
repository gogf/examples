---
title: Dependency Injection
slug: /examples/practices/injection
keywords: [dependency injection, testing, goframe, di]
description: A dependency injection example using GoFrame framework and do package
hide_title: true
---

# Dependency Injection Example

## Description

This example demonstrates how to implement dependency injection in a `GoFrame` application. The example showcases:

1. Basic dependency injection setup
2. Service layer implementation with DI
3. Controller layer with gRPC integration
4. Unit testing with mocked dependencies

The implementation focuses on making the core business logic easily testable by:
- Separating concerns using dependency injection
- Providing a clean interface for service dependencies
- Making dependencies easily mockable for testing

## Requirements

- [Go](https://golang.org/dl/) `1.22` or higher
- [Git](https://git-scm.com/downloads)
- [`GoFrame`](https://goframe.org)
- [`do`](https://github.com/samber/do)
- [MongoDB](https://www.mongodb.com)
- [Redis](https://redis.io)

## Structure

```text
injection/
├── app/
│   ├── gateway/          # API Gateway Service
│   │   ├── api/         # API definitions
│   │   ├── internal/    # Internal implementation
│   │   │   ├── cmd/    # Command line tools
│   │   │   ├── controller/  # Controllers
│   │   │   ├── model/  # Data models
│   │   │   └── service/  # Business logic
│   │   └── manifest/   # Configuration files
│   └── user/           # User Service
│       ├── api/        # API definitions
│       │   ├── entity/ # Entity definitions
│       │   └── user/   # User API proto
│       ├── internal/   # Internal implementation
│       │   ├── cmd/    # Command line tools
│       │   ├── controller/  # Controllers with DI
│       │   ├── dao/    # Data Access Objects
│       │   ├── model/  # Data models
│       │   └── service/  # Business logic with DI
│       └── manifest/   # Configuration files
├── hack/               # Development tools
└── utility/           # Common utilities
    ├── injection/     # DI utilities
    └── mongohelper/   # MongoDB helpers
```

## Features

- Dependency injection using `do` package
- MongoDB and Redis integration
- gRPC service implementation
- Comprehensive unit testing
- Clean shutdown handling
- Named dependency support

## Setup

1. Clone the repository:
    ```bash
    git clone https://github.com/gogf/examples.git
    cd examples/practices/injection
    ```

2. Install the dependencies:
    ```bash
    go mod tidy
    ```

3. Start required services using Docker:
    ```bash
    # Start MongoDB
    docker run -d --name mongo -p 27017:27017 mongo:latest

    # Start Redis
    docker run -d --name redis -p 6379:6379 redis:latest
    ```

## Usage

1. Run the application:
   ```bash
   go run main.go
   ```

2. Run the tests:
   ```bash
   go test ./...
   ```

## Implementation Details

### Dependency Injection Setup
- Uses `do` package for dependency management
- Supports both named and unnamed dependencies
- Provides helper functions for common operations

### Service Layer
- Clean separation of concerns
- Interface-based design
- Easy to test with mock implementations

### Controller Layer
- gRPC integration with DI support
- Clean error handling
- Proper resource management

### Testing
- Comprehensive unit tests
- Mock dependencies
- Clean test setup and teardown

## Notes

- Always use interfaces for dependencies to make them mockable
- Properly handle resource cleanup in `Shutdown` methods
- Use named dependencies when multiple instances of the same type are needed
- Consider using mock implementations for external services in tests


