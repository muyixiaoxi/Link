package logic

import (
	"context"

	"chat/service/user/internal/svc"
	"chat/service/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type NextUserIDLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewNextUserIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *NextUserIDLogic {
	return &NextUserIDLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// redis自增id
func (l *NextUserIDLogic) NextUserID(in *user.Empty) (*user.NextUserIDResponse, error) {
	// todo: add your logic here and delete this line

	return &user.NextUserIDResponse{}, nil
}
