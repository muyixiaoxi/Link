package logic

import (
	"context"

	"user/service/internal/svc"
	"user/service/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserDisposeFlowedLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserDisposeFlowedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserDisposeFlowedLogic {
	return &UserDisposeFlowedLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserDisposeFlowedLogic) UserDisposeFlowed(in *user.DisposeFlowedRequest) (*user.Empty, error) {
	// todo: add your logic here and delete this line

	return &user.Empty{}, nil
}
