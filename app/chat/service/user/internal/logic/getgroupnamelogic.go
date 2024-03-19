package logic

import (
	"context"

	"chat/service/user/internal/svc"
	"chat/service/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGroupNameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetGroupNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGroupNameLogic {
	return &GetGroupNameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetGroupNameLogic) GetGroupName(in *user.GetGroupNameRequest) (*user.GetGroupNameResponse, error) {
	// todo: add your logic here and delete this line

	return &user.GetGroupNameResponse{}, nil
}
