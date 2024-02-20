package handler

import (
	"chat/common/response"
	"chat/restful/internal/svc"
	"chat/restful/internal/types"
	"context"
	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/logc"
	"net/http"
	"sync"
)

type Client struct {
	Id       uint64 `json:"id"`
	Username string `json:"username"`
	conn     *websocket.Conn
}

var (
	clients     = make(map[uint64]*Client)
	clientsLock sync.Mutex
)

func ChatHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
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
		client := &Client{}
		err = conn.ReadJSON(client)
		if err != nil {
			logc.Error(context.Background(), "websocket init failed: ", err)
			return
		}
		client.conn = conn
		clientsLock.Lock()
		clients[client.Id] = client
		clientsLock.Unlock()
		conn.WriteJSON(response.InitBody(nil, response.CodeSuccess))
		defer func() {
			conn.Close()
			delete(clients, client.Id)
		}()

		message := types.Message{}
		for {
			err := client.conn.ReadJSON(&message)
			if err != nil {
				logc.Error(context.Background(), "client.conn.ReadJSON(&message) failed: ", err)
				continue
			}
			if message.Type == 1 {
				SingleChat(message)
			}
		}

	}
}

// SingleChat 单聊
func SingleChat(message types.Message) {
	// 向redis存储消息
	if c, has := clients[message.To]; has {
		c.conn.WriteJSON(message)
	}

}

// GroupChat 群聊
func GroupChat(message types.Message) {

}
