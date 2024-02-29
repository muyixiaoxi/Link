package logic

import (
	"user/restful/internal/types"
)

// SingleChat 单聊
func SingleChat(message types.Message) {
	// 在线直接转发
	if c, has := Clients[message.To]; has {
		c.Conn.WriteJSON(message)
	}
	// 离线存储消息队列

}
