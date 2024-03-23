package logic

import (
	"context"
	"strconv"

	"chat/service/chat"
	"chat/service/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetConnectorIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetConnectorIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetConnectorIdLogic {
	return &GetConnectorIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetConnectorIdLogic) GetConnectorId(in *chat.UserId) (response *chat.ConnectorId, err error) {
	response = &chat.ConnectorId{}
	response.ConnectorId, err = l.svcCtx.RDB.Hget("link:chat:online", strconv.Itoa(int(in.UserId)))
	return response, err
}
