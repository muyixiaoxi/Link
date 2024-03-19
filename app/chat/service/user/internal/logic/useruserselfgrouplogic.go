package logic

import (
	"context"

	"chat/service/user/internal/svc"
	"chat/service/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserUserSelfGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserUserSelfGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserUserSelfGroupLogic {
	return &UserUserSelfGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserUserSelfGroupLogic) UserUserSelfGroup(in *user.UserSelfGroupRequest) (*user.UserSelectGroupsResponse, error) {
	// todo: add your logic here and delete this line

	return &user.UserSelectGroupsResponse{}, nil
}
