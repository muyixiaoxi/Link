package logic

import (
	"chat/restfulB/internal/types"
	"chat/service/chat"
	"github.com/zeromicro/go-zero/core/logc"
)

func (l *ChatWSLogic) SaveGroupMessage(message types.Message, userIds []uint64) (err error) {
	groupMessage := &chat.SaveMessageRequest{
		Id:          message.Id,
		From:        message.From,
		To:          message.To,
		Type:        message.Type,
		ContentType: message.ContentType,
		Time:        message.Time,
		Content:     message.Content,
	}
	_, err = l.svcCtx.ChatRpc.SaveGroupMessage(l.ctx, &chat.SaveGroupMessageRequest{
		UserIds: userIds,
		Message: groupMessage,
	})
	if err != nil {
		logc.Error(l.ctx, "l.svcCtx.ChatRpc.SaveMessage failed:", err)
	}
	return
}
