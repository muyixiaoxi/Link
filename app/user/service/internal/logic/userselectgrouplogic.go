package logic

import (
	"context"
	"user/service/internal/svc"
	"user/service/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserSelectGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserSelectGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserSelectGroupLogic {
	return &UserSelectGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserSelectGroupLogic) UserSelectGroup(in *user.UserSelectGroupsRequest) (*user.UserSelectGroupsResponse, error) {
	// 查询群列表 , 如果两个标签id为空 , 则查询的是当前登录用户所选小标签
	//db := l.svcCtx.DB

	return &user.UserSelectGroupsResponse{}, nil
}
