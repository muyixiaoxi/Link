package userGroup

import (
	"context"
	"user/restful/internal/svc"
	"user/restful/internal/types"
	"user/service/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserCreateGroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserCreateGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserCreateGroupLogic {
	return &UserCreateGroupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserCreateGroupLogic) UserCreateGroup(req *types.UserCreateGroupRequset) error {
	// 用户创建群组
	// 封装请求参数
	createGroupParams := user.UserCreateGroupRequest{
		GroupBossId:   req.GroupBossId,
		Name:          req.Name,
		SystemTagId:   req.SystemTagId,
		UserSelfTagId: req.UserSelfTagId,
		Avatar:        req.Avatar,
	}
	_, err := l.svcCtx.UserRpc.UserCreateGroup(l.ctx, &createGroupParams)
	return err
}
