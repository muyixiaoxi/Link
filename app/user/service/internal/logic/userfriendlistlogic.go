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
	err = l.svcCtx.DB.Model(&types.User{}).Select("users.id,users.avatar,users.username,friends.remark,users.signature").Joins("join friends on users.id = friends.friend_id").Where("friends.user_id = ? and friends.is_friend = 1", in.Id).Scan(&res.List).Error
	// 获取每个用户的标签
	for i := range res.List {
		l.svcCtx.DB.Table("tb_tag").Select("IFNULL(tag_name,group_name) tag_name").Joins("join tb_user_tag on tb_tag.id = tb_user_tag.tag_id").Where("tb_user_tag.user_id = ?", res.List[i].Id).Scan(&res.List[i].TagName)
	}
	return res, err
}
