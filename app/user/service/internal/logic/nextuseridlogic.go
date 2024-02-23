package logic

import (
	"context"
	"strconv"

	"user/service/internal/svc"
	"user/service/user"

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

func (l *NextUserIDLogic) NextUserID(in *user.Empty) (*user.NextUserIDResponse, error) {
	// 获取redis自增主键的值
	key := "next_user_id"
	currentIDStr, err := l.svcCtx.RDB.Get(key)
	if err != nil {
		return nil, err
	}
	currentID, _ := strconv.Atoi(currentIDStr)
	return &user.NextUserIDResponse{
		NextUserId: uint64(currentID),
	}, nil
}
