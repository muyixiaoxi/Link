package logic

import (
	"user/restful/internal/types"
)

// SingleChat 单聊
func SingleChat(message types.Message) {
	// 向redis存储消息
	if c, has := Clients[message.To]; has {
		c.Conn.WriteJSON(message)
	}
}
