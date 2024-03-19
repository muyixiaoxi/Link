package logic

import (
	"context"

	"chat/service/user/internal/svc"
	"chat/service/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryGroupHostLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryGroupHostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryGroupHostLogic {
	return &QueryGroupHostLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *QueryGroupHostLogic) QueryGroupHost(in *user.QueryGroupHostRequest) (*user.QueryGroupHostResponse, error) {
	// todo: add your logic here and delete this line

	return &user.QueryGroupHostResponse{}, nil
}
