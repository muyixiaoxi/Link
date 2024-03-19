package logic

import (
	"context"

	"chat/service/user/internal/svc"
	"chat/service/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFriendListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserFriendListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFriendListLogic {
	return &UserFriendListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserFriendListLogic) UserFriendList(in *user.UserFriendRequest) (*user.UserFriendResponse, error) {
	// todo: add your logic here and delete this line

	return &user.UserFriendResponse{}, nil
}
