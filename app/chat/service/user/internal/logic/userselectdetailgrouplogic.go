package logic

import (
	"context"

	"chat/service/user/internal/svc"
	"chat/service/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserSelectDetailGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserSelectDetailGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserSelectDetailGroupLogic {
	return &UserSelectDetailGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserSelectDetailGroupLogic) UserSelectDetailGroup(in *user.DetailGroupRequest) (*user.DetailGroupResponse, error) {
	// todo: add your logic here and delete this line

	return &user.DetailGroupResponse{}, nil
}
