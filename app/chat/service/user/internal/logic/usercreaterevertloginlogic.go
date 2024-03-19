package logic

import (
	"context"

	"chat/service/user/internal/svc"
	"chat/service/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserCreateRevertLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserCreateRevertLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserCreateRevertLoginLogic {
	return &UserCreateRevertLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserCreateRevertLoginLogic) UserCreateRevertLogin(in *user.UserCreateRequest) (*user.UserCreateResponse, error) {
	// todo: add your logic here and delete this line

	return &user.UserCreateResponse{}, nil
}
