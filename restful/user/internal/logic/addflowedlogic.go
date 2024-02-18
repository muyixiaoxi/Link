package logic

import (
	"context"

	"Link/restful/user/internal/svc"
	"Link/restful/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddFlowedLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddFlowedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddFlowedLogic {
	return &AddFlowedLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddFlowedLogic) AddFlowed(req *types.UserFlowedRequest) error {

	return nil
}
