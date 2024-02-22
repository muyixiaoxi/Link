package logic

import (
	"context"

	"user/restful/internal/svc"
	"user/restful/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DisposeFlowedLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDisposeFlowedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DisposeFlowedLogic {
	return &DisposeFlowedLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DisposeFlowedLogic) DisposeFlowed(req *types.DisposeFlowedRequest) error {
	// todo: add your logic here and delete this line

	return nil
}
