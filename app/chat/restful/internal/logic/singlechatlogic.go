package logic

import (
	"chat/restful/internal/types"
	"chat/service/chat"
)

// SingleChat 单聊
func (l *ChatWSLogic) SingleChat(message types.Message) (err error) {
	// 在线存储+转发
	if c, has := Clients.Load(message.To); has {
		// 如果存储失败返回
		if err = l.SaveMessage(message, true); err != nil {
			return
		}
		// 如果转发失败，当作离线存储
		if err = c.(*Client).Conn.WriteJSON(message); err != nil {
			Clients.Delete(message.To)
			// 不管离线存储成功与否都返回
			err = l.SaveMessage(message, false)
			return
		}
		return
	}
	// 离线存储消息队列
	err = l.SaveMessage(message, false)
	return
}

// ReadOffline 读取自己的离线消息
func (l *ChatWSLogic) ReadOffline(id uint64) (messages []*types.Message, err error) {
	response, err := l.svcCtx.ChatRpc.GetOfflineMessage(l.ctx, &chat.GetOfflineMessageRequest{UserId: id})
	if err != nil {
		return
	}
	messages = make([]*types.Message, len(response.Messages))
	for i, resp := range response.Messages {
		message := &types.Message{
			From:        resp.From,
			To:          resp.To,
			Type:        resp.Type,
			ContentType: resp.ContentType,
			Time:        resp.Time,
			Content:     resp.Content,
		}
		messages[i] = message
	}
	return
}
