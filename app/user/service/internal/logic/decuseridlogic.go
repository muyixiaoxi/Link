package logic

import (
	"context"

	"user/service/internal/svc"
	"user/service/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type DecUserIDLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDecUserIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DecUserIDLogic {
	return &DecUserIDLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DecUserIDLogic) DecUserID(in *user.Empty) (*user.Empty, error) {
	// redis自增id 减少1
	key := "next_userId"
	_, err := l.svcCtx.RDB.Decr(key)
	if err != nil {
		return nil, err
	}
	return &user.Empty{}, nil
}
