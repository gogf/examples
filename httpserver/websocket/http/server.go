package main

import (
	"net/http"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gorilla/websocket"
)

func main() {
	var (
		s          = g.Server()
		logger     = g.Log()
		wsUpGrader = websocket.Upgrader{
			// CheckOrigin allows any origin in development
			// In production, implement proper origin checking for security
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
			// Error handler for upgrade failures
			Error: func(w http.ResponseWriter, r *http.Request, status int, reason error) {
				// Implement error handling logic here
			},
		}
	)

	// Bind WebSocket handler to /ws endpoint
	s.BindHandler("/ws", func(r *ghttp.Request) {
		// Upgrade HTTP connection to WebSocket
		ws, err := wsUpGrader.Upgrade(r.Response.Writer, r.Request, nil)
		if err != nil {
			r.Response.Write(err.Error())
			return
		}
		defer ws.Close()

		// Get request context for logging
		var ctx = r.Context()

		// Message handling loop
		for {
			// Read incoming WebSocket message
			msgType, msg, err := ws.ReadMessage()
			if err != nil {
				break // Connection closed or error occurred
			}
			// Log received message
			logger.Infof(ctx, "received message: %s", msg)
			// Echo the message back to client
			if err = ws.WriteMessage(msgType, msg); err != nil {
				break // Error writing message
			}
		}
		// Log connection closure
		logger.Info(ctx, "websocket connection closed")
	})

	// Configure static file serving
	s.SetServerRoot("static")
	// Set server port
	s.SetPort(8000)
	// Start the server
	s.Run()
}
