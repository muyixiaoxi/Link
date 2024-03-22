package logic

import (
	"context"
	"fmt"

	"chat/service/chat"
	"chat/service/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type OnlineLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOnlineLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OnlineLogic {
	return &OnlineLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OnlineLogic) Online(in *chat.OnlineRequest) (*chat.Empty, error) {
	err := l.svcCtx.RDB.Set(fmt.Sprintf("link:chat:online:%d", in.UserId), in.GetConnectorId())
	return &chat.Empty{}, err
}
