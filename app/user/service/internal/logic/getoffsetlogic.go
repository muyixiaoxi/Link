package logic

import (
	"context"
	"fmt"
	"strconv"

	"user/service/internal/svc"
	"user/service/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOffsetLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOffsetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOffsetLogic {
	return &GetOffsetLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetOffsetLogic) GetOffset(in *user.GetOffsetRequest) (resp *user.GetOffsetResponse, err error) {
	cache, err := l.svcCtx.RDB.Get(fmt.Sprintf("link:user:offset:%d", in.UserId))
	id, _ := strconv.Atoi(cache)
	resp.Offset = uint64(id)
	return
}
