package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"user/restful/internal/svc"
)

type ChatWSLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChatWSLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChatWSLogic {
	return &ChatWSLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChatWSLogic) ChatWS() error {
	// todo: add your logic here and delete this line

	return nil
}