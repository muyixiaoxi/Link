package logic

import (
	"context"
	"user/service/internal/svc"
	"user/service/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserUserSelfGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserUserSelfGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserUserSelfGroupLogic {
	return &UserUserSelfGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserUserSelfGroupLogic) UserUserSelfGroup(in *user.UserSelfGroupRequest) (resp *user.UserSelectGroupsResponse, err error) {
	//page := int(in.PageNow)
	//pageSize := int(in.PageSize)
	//offset := (page - 1) * pageSize
	// 查询和自己相关的群聊列表 1. 自己加入的群聊 2.自己创建的群聊
	//var (
	//	groupInformations []*user.UserSelectGroupsResponse
	//	total             int64
	//)
	return &user.UserSelectGroupsResponse{}, nil
}
