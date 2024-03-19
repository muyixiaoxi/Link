package logic

import (
	"context"

	"chat/service/user/internal/svc"
	"chat/service/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type SelectMyGroupCountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSelectMyGroupCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SelectMyGroupCountLogic {
	return &SelectMyGroupCountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SelectMyGroupCountLogic) SelectMyGroupCount(in *user.SelectMyGroupCountRequest) (*user.SelectMyGroupCountResponse, error) {
	// todo: add your logic here and delete this line

	return &user.SelectMyGroupCountResponse{}, nil
}
