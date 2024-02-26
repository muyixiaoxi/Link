package logic

import (
	"context"

	"user/service/internal/svc"
	"user/service/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddUserIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddUserIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserIdLogic {
	return &AddUserIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddUserIdLogic) AddUserId(in *user.Empty) (*user.Empty, error) {
	//reids 实现自增主键
	key := "next_userId"
	_, err := l.svcCtx.RDB.Incr(key)
	if err != nil {
		return nil, err
	}
	return &user.Empty{}, nil
}
