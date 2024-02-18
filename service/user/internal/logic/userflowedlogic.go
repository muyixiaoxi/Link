package logic

import (
	"Link/service/user/internal/types"
	"context"

	"Link/service/user/internal/svc"
	"Link/service/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFlowedLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserFlowedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFlowedLogic {
	return &UserFlowedLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserFlowedLogic) UserFlowed(in *user.UserAddRequest) (response *user.Empty, err error) {
	model := types.ApplyFor{
		UserID:  in.Id,
		BeId:    in.BeId,
		Message: in.Message,
		Type:    in.Type,
	}
	err = l.svcCtx.DB.Create(model).Error
	return
}
