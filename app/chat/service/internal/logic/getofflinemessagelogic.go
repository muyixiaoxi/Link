package logic

import (
	"context"
	"encoding/json"
	"fmt"

	"chat/service/chat"
	"chat/service/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOfflineMessageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOfflineMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOfflineMessageLogic {
	return &GetOfflineMessageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetOfflineMessageLogic) GetOfflineMessage(in *chat.GetOfflineMessageRequest) (response *chat.Messages, err error) {
	response = &chat.Messages{}
	// 首先获取自己redis，如果redis没有获取Mysql
	messages, err := l.svcCtx.RDB.Zrange(fmt.Sprintf("link:chat:user:%d", in.UserId), 0, -1)
	if err != nil {
		return
	}
	l.svcCtx.RDB.Zrem(fmt.Sprintf("link:chat:user:%d", in.UserId), messages)

	// 如果不存在离线消息去 mysql查询 ,后面再写

	response.Messages = make([]*chat.Message, len(messages))
	for i, message := range messages {
		tmp := &chat.Message{}
		json.Unmarshal([]byte(message), tmp)
		response.Messages[i] = tmp
	}
	return response, err
}
