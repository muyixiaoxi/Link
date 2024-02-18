package logic

import (
	"context"

	"Link/service/user/internal/svc"
	"Link/service/user/user"

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
