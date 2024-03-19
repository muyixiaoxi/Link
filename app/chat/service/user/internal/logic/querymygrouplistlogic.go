package logic

import (
	"context"

	"chat/service/user/internal/svc"
	"chat/service/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryMyGroupListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryMyGroupListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryMyGroupListLogic {
	return &QueryMyGroupListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *QueryMyGroupListLogic) QueryMyGroupList(in *user.QueryMyGroupListRequest) (*user.UserSelectGroupsResponse, error) {
	// todo: add your logic here and delete this line

	return &user.UserSelectGroupsResponse{}, nil
}
