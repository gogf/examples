# GoFrame MCP 示例

本示例展示了如何使用 GoFrame 的 `ghttp` 服务器和 `mark3labs/mcp-go` SDK 实现 [Model Context Protocol (MCP)](https://modelcontextprotocol.io/) 服务器。

## 先决条件

- Go 1.20 或更高版本
- MCP 客户端（例如 Claude Desktop 或 IDE 扩展）

## 安装

1. 初始化模块（如果尚未完成）：
   ```bash
   go mod tidy
   ```

2. 运行服务器：
   ```bash
   go run main.go
   ```

## 用法

服务器暴露了两个端点：
- `GET /sse`: 用于建立连接的 Server-Sent Events 端点。
- `POST /message`: 用于发送 JSON-RPC 消息的端点。

### 连接 Claude Desktop

要在 Claude Desktop 中使用此 MCP 服务器，请将以下配置添加到您的 `claude_desktop_config.json`：

```json
{
  "mcpServers": {
    "goframe-mcp": {
      "command": "go",
      "args": ["run", "path/to/your/project/httpserver/mcp/main.go"]
    }
  }
}
```

*注意：由于此示例运行 HTTP 服务器，您可能需要使用支持 SSE 的 MCP 客户端。上述配置假设您将其作为本地进程运行。如果您想通过 HTTP (SSE) 连接，您的客户端必须支持远程 MCP 连接。*

直接 SSE 测试：
1. 启动服务器。
2. 连接到 `http://localhost:8080/sse`。

## MCP 客户端配置

如果你的客户端或 IDE 支持通过配置文件定义 MCP 服务器，可以在项目根目录下创建 `.vscode/mcp.json`，示例内容如下：

```json
{
  "servers": {
    "add": {
      "url": "http://localhost:8080/sse",
      "type": "http"
    }
  },
  "inputs": []
}
```

说明：
- `servers` 下的键名（例如 `add`）为服务器标识名称。
- `url` 指向示例的 SSE 端点 `http://localhost:8080/sse`。
- `type` 表示连接类型，此处为 `http`（SSE）。
- `inputs` 可选，用于在客户端中预置输入。

保存后，支持该配置的客户端即可使用该文件连接到本示例 MCP 服务。


## 功能

- **工具**: 暴露一个 `add` 工具，可接收两个数字（`a` 和 `b`），返回它们的和。
- **示例**: 你可以通过 MCP 客户端发送如下计算请求：

```
mcp 计算 一加一
```
或英文：
```
mcp add 1 and 1
```
服务器将返回：
```
结果: 1 + 1 = 2
```

## 代码结构

- `main.go`: 包含服务器设置、工具注册以及使用 `ghttp.WrapH` 进行的路由绑定。
