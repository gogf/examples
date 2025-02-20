---
title: Rate Limit
slug: /examples/httpserver/rate-limit
keywords: [http, server, rate limit, middleware, goframe]
description: Rate limiting in a HTTP server using GoFrame framework
hide_title: true
---

# HTTP Server Rate Limit


## Description

This example demonstrates how to implement rate limiting in a HTTP server using `GoFrame`. It showcases how to protect your API endpoints from being overwhelmed by too many requests using the token bucket algorithm implemented by `golang.org/x/time/rate` package.

The example implements:
- A simple HTTP endpoint `/hello` that returns a greeting message
- A rate limiting middleware that restricts requests to 10 per second
- Proper error handling when rate limit is exceeded (HTTP 429 Too Many Requests)

## Requirements

- [Go](https://golang.org/dl/) `1.22` or higher
- [Git](https://git-scm.com/downloads)
- [GoFrame](https://goframe.org)

## Structure

- `go.mod`: The Go module file.
- `main.go`: The main application entry point.

## Features

- Token bucket based rate limiting
- Configurable request rate and burst size
- Global middleware implementation
- Clean API endpoint implementation using GoFrame's binding feature
- Request validation and documentation using metadata
- Proper error handling and status codes

## Usage

1. Run the example:
   ```bash
   go run main.go
   ```

2. The server will start at http://127.0.0.1:8080

3. Test the rate limiting:
   ```bash
   # Normal request
   curl "http://127.0.0.1:8080/hello?name=world"
   
   # To test rate limiting, send multiple requests quickly:
   for i in {1..20}; do curl "http://127.0.0.1:8080/hello?name=world"; done
   ```

## Implementation Details

The rate limiting is implemented using:
1. `golang.org/x/time/rate.Limiter` for token bucket algorithm
2. GoFrame's middleware system for request interception
3. Clean request/response structs with validation and documentation

Key components:
- Rate limit is set to 10 requests per second
- Burst size is set to 1 (no bursting allowed)
- Requests exceeding the limit receive HTTP 429 status code
- Request validation ensures required parameters are provided

## Notes

- The rate limit is global (applies to all clients)
- No persistence of rate limit state (resets on server restart)
- Rate limit can be configured by modifying the limiter parameters
- Consider using distributed rate limiting for production environments
