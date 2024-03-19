package logic

import (
	"context"

	"chat/service/user/internal/svc"
	"chat/service/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchMyGroupByNameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchMyGroupByNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchMyGroupByNameLogic {
	return &SearchMyGroupByNameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchMyGroupByNameLogic) SearchMyGroupByName(in *user.SearchMyGroupByNameRequest) (*user.MyGroupResponse, error) {
	// todo: add your logic here and delete this line

	return &user.MyGroupResponse{}, nil
}
