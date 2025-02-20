---
title: WebSocket
slug: /examples/httpserver/websocket
keywords: [websocket, server, client, goframe, https]
description: A WebSocket server and client implementation using GoFrame framework
hide_title: true
---

# `WebSocket` Server and Client

## Description

This example demonstrates how to implement `WebSocket` communication using `GoFrame`. The example includes both `HTTP` and `HTTPS` implementations with the following components:

1. `WebSocket` server that handles both secure and non-secure connections
2. `WebSocket` client implementations for both `HTTP` and `HTTPS` connections
3. Example of handling `WebSocket` message exchange

The implementation demonstrates:
- Basic `WebSocket` server setup
- Secure `WebSocket` (`WSS`) implementation
- Client connection handling
- Message exchange between client and server
- Proper connection cleanup

## Requirements

- [Go](https://golang.org/dl/) `1.22` or higher
- [Git](https://git-scm.com/downloads)
- [`GoFrame`](https://goframe.org)
- [`Gorilla WebSocket`](https://github.com/gorilla/websocket)

## Structure

```text
websocket/
├── http/
│   ├── client.go      # HTTP WebSocket client
│   ├── server.go      # HTTP WebSocket server
│   └── static/        # Static files directory
├── https/
│   ├── client.go      # HTTPS WebSocket client
│   ├── server.go      # HTTPS WebSocket server
│   ├── static/        # Static files directory
│   └── certs/         # SSL certificates
└── README.md
```

## Features

- `WebSocket` server implementation
- Secure `WebSocket` (`WSS`) support
- Client connection handling
- Message echo functionality
- Connection lifecycle management
- Origin checking (configurable)
- Error handling

## Setup

1. Clone the repository:
    ```bash
    git clone https://github.com/gogf/examples.git
    cd examples/httpserver/websocket
    ```

2. Install the dependencies:
    ```bash
    go mod tidy
    ```

3. For HTTPS/WSS support, generate self-signed certificates (optional):
    ```bash
    mkdir -p https/certs
    openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout https/certs/server.key -out https/certs/server.crt
    ```

## Usage

### HTTP WebSocket

1. Start the server:
   ```bash
   cd http
   go run server.go
   ```

2. Run the client:
   ```bash
   cd http
   go run client.go
   ```

3. Access the demo page in browser:
   Open `http://127.0.0.1:8000` in your browser to see the `WebSocket` demo in action. The page provides a simple chat interface where you can send and receive messages in real-time.

### HTTPS WebSocket

1. Start the secure server:
   ```bash
   cd https
   go run server.go
   ```

2. Run the secure client:
   ```bash
   cd https
   go run client.go
   ```

3. Access the demo page in browser:
   Open `https://127.0.0.1:8000` in your browser to see the `WebSocket` demo in action. Note: Since we're using a self-signed certificate, your browser may show a security warning, which is normal.

## Implementation Details

### Server Features
- `WebSocket` upgrade handling
- Message echo functionality
- Connection lifecycle management
- Configurable origin checking
- Error handling and logging

### Client Features
- Connection establishment
- Message sending and receiving
- `TLS` configuration for secure connections
- Clean connection closure

## Notes

- The `WebSocket` server echoes back any message it receives
- For `HTTPS`/`WSS`, self-signed certificates are used in the example
- In production, implement proper origin checking
- Handle connection closure properly to avoid resource leaks

## More Information

For more details about WebSocket implementation, please refer to:
- [GoFrame WebSocket Guide](https://goframe.org/docs/web/senior-websocket)
- [Gorilla WebSocket Documentation](https://github.com/gorilla/websocket)

