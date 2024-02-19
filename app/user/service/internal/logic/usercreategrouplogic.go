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
	// 用户创建群聊
	groupChat := types.GroupChat{
		GroupBossID:   in.GroupBossId,
		Name:          in.Name,
		SystemTagId:   in.SystemTagId,
		UserSelfTagId: in.UserSelfTagId,
		Avatar:        in.Avatar,
	}
	err := l.svcCtx.DB.Create(&groupChat).Error
	return &user.Empty{}, err
}
