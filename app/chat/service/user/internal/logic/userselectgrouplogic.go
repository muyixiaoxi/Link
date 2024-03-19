package logic

import (
	"context"

	"chat/service/user/internal/svc"
	"chat/service/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserSelectGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserSelectGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserSelectGroupLogic {
	return &UserSelectGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserSelectGroupLogic) UserSelectGroup(in *user.UserSelectGroupsRequest) (*user.UserSelectGroupsResponse, error) {
	// todo: add your logic here and delete this line

	return &user.UserSelectGroupsResponse{}, nil
}
