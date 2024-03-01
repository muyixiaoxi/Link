package logic

import (
	"context"
	"fmt"

	"user/service/internal/svc"
	"user/service/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetOffsetLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSetOffsetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetOffsetLogic {
	return &SetOffsetLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SetOffsetLogic) SetOffset(in *user.SetOffsetRequest) (resp *user.Empty, err error) {
	resp = &user.Empty{}
	err = l.svcCtx.RDB.Set(fmt.Sprintf("link:user:offset:%d", in.UserId), string(in.Offset))
	return
}
