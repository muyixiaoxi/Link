package logic

import (
	"context"
	"user/service/internal/svc"
	"user/service/internal/types"
	"user/service/user"

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

func (l *UserQueryFriendLogic) UserQueryFriend(in *user.UserQueryFriendRequest) (resp *user.UserFriendResponse, err error) {
	resp = &user.UserFriendResponse{}
	err = l.svcCtx.DB.Model(&types.User{}).Select("users.id,users.username,friends.remark,users.avatar,users.signature").Joins("join friends on users.id = friends.friend_id").Where("friends.user_id = ? and (users.username like ? or friends.remark like ?)",
		in.UserId, "%"+in.Param+"%", "%"+in.Param+"%").Find(&resp.List).Error

	// 获取每个用户的标签
	for i := range resp.List {
		l.svcCtx.DB.Table("tb_tag").Select("IFNULL(tag_name,group_name) tag_name").Joins("join tb_user_tag on tb_tag.id = tb_user_tag.tag_id").Where("tb_user_tag.user_id = ?", resp.List[i].Id).Scan(&resp.List[i].TagName)
	}
	return resp, err
}
