package logic

import (
	"context"

	"chat/service/user/internal/svc"
	"chat/service/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserGetApplyForLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserGetApplyForLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserGetApplyForLogic {
	return &UserGetApplyForLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserGetApplyForLogic) UserGetApplyFor(in *user.UserGetApplyForRequest) (*user.UserGetApplyForResponse, error) {
	// todo: add your logic here and delete this line

	return &user.UserGetApplyForResponse{}, nil
}
