package logic

import (
	"chat/restful/internal/types"
	"chat/service/chat"
	"github.com/zeromicro/go-zero/core/logc"
)

func (l *ChatWSLogic) SaveMessage(message types.Message, online bool) (err error) {
	_, err = l.svcCtx.ChatRpc.SaveOfflineMessage(l.ctx, &chat.SaveMessageRequest{
		Id:          message.Id,
		From:        message.From,
		To:          message.To,
		Type:        message.Type,
		ContentType: message.ContentType,
		Time:        message.Time,
		Content:     message.Content,
		Online:      online,
	})
	if err != nil {
		logc.Error(l.ctx, "l.svcCtx.ChatRpc.SaveMessage failed:", err)
	}
	return
}
