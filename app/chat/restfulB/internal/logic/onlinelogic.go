package logic

import (
	"chat/service/chat"
)

func (l *ChatWSLogic) GetConnector(userId uint64) (ip string) {
	resp, _ := l.svcCtx.ChatRpc.GetConnectorId(l.ctx, &chat.UserId{
		UserId: userId,
	})
	return resp.ConnectorId
}

func (l *ChatWSLogic) Offline(userId uint64) {
	l.svcCtx.ChatRpc.Offline(l.ctx, &chat.UserId{
		UserId: userId,
	})
}

func (l *ChatWSLogic) Online(userId uint64) {
	l.svcCtx.ChatRpc.Online(l.ctx, &chat.OnlineRequest{
		UserId:      userId,
		ConnectorId: Conn.LocalAddr().String(),
	})
}
