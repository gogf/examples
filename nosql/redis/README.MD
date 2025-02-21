---
title: Redis
slug: /examples/nosql/redis
keywords: [nosql, redis, cache, database, goframe]
description: An example demonstrating Redis integration in GoFrame
hide_title: true
sidebar_position: 1
---


# GoFrame Redis Example

This example demonstrates how to use `Redis` with `GoFrame` framework.

## Overview

This example shows:
1. How to configure `Redis` connection using `YAML` configuration
2. How to create a `Redis` client
3. Basic `Redis` operations (`SET`/`GET`)

## Requirements

- `Go 1.15` or higher
- `Redis` server
- `GoFrame v2`

## Configuration

The `Redis` configuration is stored in `config.yaml`:

```yaml
redis:
  address: "127.0.0.1:6379"
  password:
```

You can modify these settings according to your `Redis` server configuration.

## Running Redis with Docker

If you don't have `Redis` installed locally, you can quickly start a `Redis` instance using `Docker`:

```bash
# Run Redis container
docker run --name redis-test -p 6379:6379 -d redis:latest

# Verify the container is running
docker ps

# If you need to stop the container later
docker stop redis-test

# If you need to remove the container
docker rm redis-test
```

For `Redis` with password authentication:

```bash
# Run Redis with password
docker run --name redis-test -p 6379:6379 -d redis:latest redis-server --requirepass your_password

# Remember to update config.yaml accordingly:
# redis:
#   address: "127.0.0.1:6379"
#   password: "your_password"
```

## Running the Example

1. Make sure your `Redis` server is running
2. Update the `config.yaml` if needed
3. Run the example:

```bash
go run main.go
```

## Code Structure

- `main.go`: Contains the main logic and `Redis` client initialization
- `config.yaml`: `Redis` configuration file

## Features

1. `Redis` client initialization with error handling
2. Configuration management using `GoFrame`'s configuration system
3. Basic `Redis` operations demonstration
4. Proper error handling and logging

## Further Reading

For more advanced `Redis` usage, please refer to the third-party package [github.com/redis/go-redis](https://github.com/redis/go-redis).

## Notes

- The example uses `go-redis/v9` client
- All operations are performed with context for proper cancellation and timeout handling
- The code includes proper error handling and logging
