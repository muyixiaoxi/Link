package logic

import (
	"context"
	"user/service/user"

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
	_, err := l.svcCtx.UserRpc.UserDisposeFlowed(context.Background(), &user.DisposeFlowedRequest{
		From:   req.From,
		To:     req.To,
		Type:   req.Type,
		Remark: req.Remark,
		Result: req.Res,
	})
	if err != nil {
		logx.Error("l.svcCtx.UserRpc.UserDisposeFlowed failed: ", err)
		return err
	}

	return err
}
