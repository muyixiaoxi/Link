package logic

import (
	"context"
	"fmt"
	"strconv"

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
	err := l.svcCtx.RDB.Hset(fmt.Sprintf("link:chat:online"), strconv.Itoa(int(in.UserId)), in.GetConnectorId())
	return &chat.Empty{}, err
}
