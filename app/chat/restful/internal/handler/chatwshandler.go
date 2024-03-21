package handler

import (
	"chat/common/response"
	"chat/restful/internal/logic"
	"chat/restful/internal/svc"
	"chat/restful/internal/types"
	"context"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/logc"
	"net/http"
	"time"
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
		logic.Clients.Swap(client.Id, client)
		conn.WriteJSON(response.InitBody(nil, response.CodeSuccess))
		defer func() {
			conn.Close()
			logic.Clients.Delete(client.Id)
		}()

		l := logic.NewChatWSLogic(r.Context(), svcCtx)

		go func() {
			unread, err := l.ReadOffline(client.Id)
			if err != nil {
				conn.WriteJSON(response.InitBody(nil, response.CodeSyncMessageFailed))
			}
			if unread != nil {
				for _, message := range unread {
					client.Conn.WriteJSON(message)
				}
			}
		}()

		message := types.Message{}
		for {
			conn.SetReadDeadline(time.Now().Add(10 * time.Second))
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
				if err = l.SingleChat(message); err == nil {
					conn.WriteJSON(fmt.Sprintf("messageId:%s", message.Id))
				}
			}
			if message.Type == 2 {
				//群聊
				l.GroupChat(message)
			}
		}

	}
}
