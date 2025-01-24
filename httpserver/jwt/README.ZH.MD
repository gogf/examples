---
title: JWT认证
slug: /examples/httpserver/jwt
keywords: [http, server, jwt, authentication, goframe]
description: 使用GoFrame框架实现的JWT认证示例
hide_title: true
sidebar_position: 0
---

# GoFrame JWT认证示例

本示例展示了如何在`GoFrame`HTTP服务器中使用 `github.com/golang-jwt/jwt` 包实现`JWT(JSON Web Token)`认证。

## 功能特性

- 用户登录接口，生成`JWT`令牌
- 使用`JWT`中间件保护路由
- 令牌验证和解析
- 受保护资源访问示例
- 标准的`GoFrame`项目结构

## 项目结构

```text
jwt/
├── api/
│   └── v1/
│       └── auth.go         # API接口定义
├── internal/
│   ├── controller/
│   │   └── auth.go        # 业务逻辑实现
│   └── middleware/
│       └── jwt.go         # JWT中间件
└── main.go                # 入口文件
```

## API接口

1. 登录接口: `POST /login`
   ```json
   {
       "username": "admin",
       "password": "password"
   }
   ```

2. 受保护资源: `GET /api/protected`
   - 需要在Authorization头部携带Bearer令牌
   - 示例: `Authorization: Bearer your-token-here`

## 运行示例

1. 启动服务器:
   ```bash
   go run main.go
   ```

2. 服务器将在8199端口启动

## 测试API

1. 登录获取令牌:
   ```bash
   curl -X POST http://localhost:8199/login \
   -H "Content-Type: application/json" \
   -d '{"username":"admin","password":"password"}'
   ```

2. 访问受保护的接口:
   ```bash
   curl http://localhost:8199/api/protected \
   -H "Authorization: Bearer your-token-here"
   ```

## 安全注意事项

- 在生产环境中，请替换硬编码的密钥为安全的值
- 将用户凭证存储在数据库中
- 实现适当的密码哈希
- 考虑实现刷新令牌机制
- 为登录尝试添加速率限制

## 参考资料

有关`JWT`实现的更多详细信息，请参考第三方组件文档：
- [github.com/golang-jwt/jwt](https://github.com/golang-jwt/jwt)
