package logic

import (
	"context"
	"fmt"

	"chat/service/chat"
	"chat/service/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type OfflineLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOfflineLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OfflineLogic {
	return &OfflineLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OfflineLogic) Offline(in *chat.UserId) (*chat.Empty, error) {
	_, err := l.svcCtx.RDB.Del(fmt.Sprintf("link:chat:online:%d", in.UserId))
	return &chat.Empty{}, err
}
