package main

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gorilla/websocket"
)

func main() {
	var (
		ctx    = context.Background()
		logger = g.Log()
	)

	// Connect to WebSocket server using default dialer
	ws, _, err := websocket.DefaultDialer.Dial("ws://127.0.0.1:8000/ws", nil)
	if err != nil {
		logger.Fatalf(ctx, "dial failed: %+v", err)
	}
	// Ensure connection is closed when function returns
	defer ws.Close()

	// Send a test message to the server
	err = ws.WriteMessage(websocket.TextMessage, []byte("hello"))
	if err != nil {
		logger.Fatalf(ctx, "ws.WriteMessage failed: %+v", err)
	}

	// Read the server's response
	_, msg, err := ws.ReadMessage()
	if err != nil {
		logger.Fatalf(ctx, "ws.ReadMessage failed: %+v", err)
		return
	}

	logger.Infof(ctx, `received message: %s`, msg)

	// Cleanly close the connection by sending a close message
	// This is important for proper connection cleanup
	err = ws.WriteMessage(websocket.CloseMessage, []byte("going to close"))
	if err != nil {
		logger.Fatalf(ctx, "ws.WriteMessage failed: %+v", err)
	}
}
