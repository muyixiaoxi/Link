package logic

import (
	"chat/service/chat"
	"chat/service/internal/svc"
	"context"
	"strconv"

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
	_, err := l.svcCtx.RDB.Hdel("link:chat:online", strconv.Itoa(int(in.UserId)))
	return &chat.Empty{}, err
}
