package logic

import (
	"chat/restfulB/internal/types"
	"chat/service/chat"
)

// SingleChat 单聊
func (l *ChatWSLogic) SingleChat(userId uint64, message types.Message) (err error) {
	// 在线存储+转发
	if c, has := Clients.Load(userId); has {
		// 如果存储失败返回
		if err = l.SaveMessage(message, true); err != nil {
			return
		}
		// 如果转发失败，当作离线存储
		if err = c.(*Client).Conn.WriteJSON(message); err != nil {
			l.Offline(userId)
			Clients.Delete(userId)
			// 不管离线存储成功与否都返回
			err = l.SaveMessage(message, false)
			return
		}
		return
	}
	// 离线判断用户是否在其他服务端在线
	resp, _ := l.svcCtx.ChatRpc.GetConnectorId(l.ctx, &chat.UserId{
		UserId: userId,
	})
	if resp != nil && resp.ConnectorId != "" {
		if err = Producer(map[string][]uint64{resp.ConnectorId: []uint64{userId}}, message); err != nil {
			l.Offline(userId)
			l.SaveMessage(message, false)
		}
		return
	}
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
