package userGroup

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"user/restful/internal/svc"
)

type HomeGroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomeGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomeGroupLogic {
	return &HomeGroupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomeGroupLogic) HomeGroup() error {
	// todo: add your logic here and delete this line

	return nil
}
