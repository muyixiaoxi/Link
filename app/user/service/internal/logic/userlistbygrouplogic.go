package logic

import (
	"context"
	"user/service/internal/svc"
	"user/service/internal/types"
	"user/service/user"

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

// UserListByGroup 查询群聊中的人员信息
func (l *UserListByGroupLogic) UserListByGroup(in *user.SelectUserListByGroupRequest) (*user.SelectUserListByGroupResponse, error) {
	var userList []*user.SelectUserListByGroup
	err := l.svcCtx.DB.
		Table("users").
		Joins("LEFT JOIN user_group_chats ON users.id = user_group_chats.user_id").
		Where("user_group_chats.group_chat_id = ?", in.GroupId).
		Find(&userList).Error
	if err != nil {
		return nil, err
	}
	//查询此群的群主id
	var bossId uint64
	err = l.svcCtx.DB.Model(&types.GroupChat{}).Select("group_boss_id").Where("id = ?", in.GroupId).Find(&bossId).Error
	return &user.SelectUserListByGroupResponse{
		GroupBossId: bossId,
		UserList:    userList}, nil
}
