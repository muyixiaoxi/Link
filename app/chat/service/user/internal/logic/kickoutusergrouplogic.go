package logic

import (
	"context"

	"chat/service/user/internal/svc"
	"chat/service/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type KickOutUserGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewKickOutUserGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *KickOutUserGroupLogic {
	return &KickOutUserGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *KickOutUserGroupLogic) KickOutUserGroup(in *user.KickOutUserGroupRequest) (*user.KickOutUserGroupResponse, error) {
	// todo: add your logic here and delete this line

	return &user.KickOutUserGroupResponse{}, nil
}
