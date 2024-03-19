package logic

import (
	"context"

	"chat/service/user/internal/svc"
	"chat/service/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchStrangerGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchStrangerGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchStrangerGroupLogic {
	return &SearchStrangerGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchStrangerGroupLogic) SearchStrangerGroup(in *user.SearchStrangerGroupRequest) (*user.UserSelectGroupsResponse, error) {
	// todo: add your logic here and delete this line

	return &user.UserSelectGroupsResponse{}, nil
}
