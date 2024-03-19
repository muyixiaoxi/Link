package logic

import (
	"context"

	"chat/service/user/internal/svc"
	"chat/service/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserQueryFriendLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserQueryFriendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserQueryFriendLogic {
	return &UserQueryFriendLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserQueryFriendLogic) UserQueryFriend(in *user.UserQueryFriendRequest) (*user.UserFriendResponse, error) {
	// todo: add your logic here and delete this line

	return &user.UserFriendResponse{}, nil
}
