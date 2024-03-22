package logic

import (
	"context"
	"fmt"

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
	response.ConnectorId, err = l.svcCtx.RDB.Get(fmt.Sprintf("link:chat:online:%d", in.UserId))
	return response, err
}
