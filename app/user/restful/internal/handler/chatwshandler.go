package handler

import (
	"context"
	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/logc"
	"net/http"
	"user/common/response"
	"user/restful/internal/logic"
	"user/restful/internal/svc"
	"user/restful/internal/types"
)

func chatWSHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		upgrader := websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			// 允许跨域
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		}

		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			logc.Error(context.Background(), "websocket link failed: ", err)
			return
		}
		client := &logic.Client{}
		err = conn.ReadJSON(client)
		if err != nil {
			logc.Error(context.Background(), "websocket init failed: ", err)
			return
		}
		client.Conn = conn
		logic.ClientsLock.Lock()
		logic.Clients[client.Id] = client
		logic.ClientsLock.Unlock()
		conn.WriteJSON(response.InitBody(nil, response.CodeSuccess))
		defer func() {
			conn.Close()
			delete(logic.Clients, client.Id)
		}()

		message := types.Message{}
		for {
			err := client.Conn.ReadJSON(&message)
			if err != nil {
				logc.Error(context.Background(), "client.conn.ReadJSON(&message) failed: ", err)
				continue
			}
			// 心跳检测
			if message.Type == 0 {
				client.Conn.WriteJSON(message)
			}
			if message.Type == 1 {
				logic.SingleChat(message)
			}
		}

	}
}
