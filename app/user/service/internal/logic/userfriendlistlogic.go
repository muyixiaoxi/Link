package logic

import (
	"context"
	"user/service/internal/types"

	"user/service/internal/svc"
	"user/service/user"

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

func (l *UserFriendListLogic) UserFriendList(in *user.UserFriendRequest) (res *user.UserFriendResponse, err error) {
	res = &user.UserFriendResponse{}
	err = l.svcCtx.DB.Model(&types.User{}).Select("users.id id,users.avatar avatar,users.username username,friends.remark remark").Joins("join friends on users.id = friends.friend_id").Where("friends.user_id = ? and friends.is_friend = 1", in.Id).Scan(&res.List).Error
	return res, err
}
