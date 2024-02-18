package handler

import (
	"Link/internal/response"
	"fmt"
	"net/http"

	"Link/restful/chat/internal/svc"
	"github.com/gorilla/websocket"
)

func WebSocketHandler(svcCtx *svc.ServiceContext) func(w http.ResponseWriter, r *http.Request) {
	// Upgrader 用于将Http连接升级为 WebSocket 连接
	upgader := websocket.Upgrader{
		// 解决跨域问题
		CheckOrigin: func(r *http.Request) bool {
			// 允许所有来源的连接
			return true
		},
	}

	return func(w http.ResponseWriter, r *http.Request) {
		// 将http 连接升级为WebSocket 连接
		conn, err := upgader.Upgrade(w, r, nil)
		if err != nil {
			response.Response(w, nil, response.CodeWebSocketLinkFail)
			return
		}

		// 处理 WebSocket 连接的逻辑
		go handleWebSocket(conn, svcCtx)
	}
}
func handleWebSocket(conn *websocket.Conn, svcCtx *svc.ServiceContext) {
	defer conn.Close()
	for {
		// 读取客户端发送的消息
		_, msg, err := conn.ReadMessage()
		if err != nil {
			// 处理连接关闭等错误
			break
		}
		fmt.Println(string(msg))

		// 在这里编写处理客户端消息的逻辑
		// 例如，调用适当的业务逻辑组件来处理消息
		// resp, err := l.HandleWebSocketMessage(msg)

		// 处理完逻辑后，可以选择向客户端发送响应消息
		// conn.WriteMessage(websocket.TextMessage, resp)
	}
}
