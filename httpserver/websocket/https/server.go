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
			CheckOrigin: func(r *http.Request) bool {
				// In production, you should implement proper origin checking
				return true
			},
			Error: func(w http.ResponseWriter, r *http.Request, status int, reason error) {
				// Error callback function.
			},
		}
	)

	s.BindHandler("/ws", func(r *ghttp.Request) {
		ws, err := wsUpGrader.Upgrade(r.Response.Writer, r.Request, nil)
		if err != nil {
			r.Response.Write(err.Error())
			return
		}
		defer ws.Close()

		var ctx = r.Context()
		for {
			msgType, msg, err := ws.ReadMessage()
			if err != nil {
				break
			}
			logger.Infof(ctx, "received message: %s", msg)
			if err = ws.WriteMessage(msgType, msg); err != nil {
				break
			}
		}
		logger.Info(ctx, "websocket connection closed")
	})
	s.EnableHTTPS("certs/server.crt", "certs/server.key")
	s.SetServerRoot("static")
	s.SetPort(8000)
	s.Run()
}
