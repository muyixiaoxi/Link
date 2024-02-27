package logic

import (
	"context"
	"user/service/internal/svc"
	"user/service/internal/types"
	"user/service/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserCreateGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserCreateGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserCreateGroupLogic {
	return &UserCreateGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserCreateGroupLogic) UserCreateGroup(in *user.UserCreateGroupRequest) (*user.Empty, error) {
	tx := l.svcCtx.DB.Begin()
	// 用户创建群聊
	groupChat := types.GroupChat{
		GroupBossID:   in.GroupBossId,
		Name:          in.Name,
		SystemTagId:   in.SystemTagId,
		UserSelfTagId: in.UserSelfTagId,
		Avatar:        in.Avatar,
	}
	//插入群聊信息
	err := tx.Create(&groupChat).Error
	//向中间表中插入一条信息
	userGroup := types.UserGroupChat{
		GroupChatID: uint64(groupChat.ID),
		UserID:      in.GroupBossId,
	}
	err = tx.Create(&userGroup).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	return &user.Empty{}, tx.Commit().Error
}
