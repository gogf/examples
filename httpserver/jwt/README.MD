---
title: JWT Authentication
slug: /examples/httpserver/jwt
keywords: [http, server, jwt, authentication, goframe]
description: JWT authentication example using GoFrame framework
hide_title: true
sidebar_position: 0
---

# JWT Authentication Example with GoFrame

This example demonstrates how to implement JWT (JSON Web Token) authentication in a GoFrame HTTP server using the `github.com/golang-jwt/jwt` package.

## Features

- User login endpoint that generates JWT tokens
- Protected routes using JWT middleware
- Token validation and parsing
- Example of accessing protected resources
- Standard GoFrame project structure

## Project Structure

```
jwt/
├── api/
│   └── v1/
│       └── auth.go         # API interface definitions
├── internal/
│   ├── controller/
│   │   └── auth.go        # Business logic implementation
│   └── middleware/
│       └── jwt.go         # JWT middleware
└── main.go                # Entry point
```

## API Endpoints

1. Login: `POST /login`
   ```json
   {
       "username": "admin",
       "password": "password"
   }
   ```

2. Protected Resource: `GET /api/protected`
   - Requires Bearer token in Authorization header
   - Example: `Authorization: Bearer your-token-here`

## Running the Example

1. Start the server:
   ```bash
   go run main.go
   ```

2. The server will start on port 8000

## Testing the API

1. Login to get a token:
   ```bash
   curl -X POST http://localhost:8000/login \
   -H "Content-Type: application/json" \
   -d '{"username":"admin","password":"password"}'
   ```

2. Access protected endpoint:
   ```bash
   curl http://localhost:8000/api/protected \
   -H "Authorization: Bearer your-token-here"
   ```

## Security Notes

- In production, replace the hardcoded secret key with a secure value
- Store user credentials in a database
- Implement proper password hashing
- Consider implementing refresh tokens
- Add rate limiting for login attempts

## References

For more detailed information about JWT implementation, please refer to the third-party component documentation:
- [github.com/golang-jwt/jwt](https://github.com/golang-jwt/jwt)
