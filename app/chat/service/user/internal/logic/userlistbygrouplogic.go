package logic

import (
	"context"

	"chat/service/user/internal/svc"
	"chat/service/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserListByGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserListByGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserListByGroupLogic {
	return &UserListByGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserListByGroupLogic) UserListByGroup(in *user.SelectUserListByGroupRequest) (*user.SelectUserListByGroupResponse, error) {
	// todo: add your logic here and delete this line

	return &user.SelectUserListByGroupResponse{}, nil
}
