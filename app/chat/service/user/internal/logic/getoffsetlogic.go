package logic

import (
	"context"

	"chat/service/user/internal/svc"
	"chat/service/user/user"

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

// 偏移量
func (l *GetOffsetLogic) GetOffset(in *user.GetOffsetRequest) (*user.GetOffsetResponse, error) {
	// todo: add your logic here and delete this line

	return &user.GetOffsetResponse{}, nil
}
