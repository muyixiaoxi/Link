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
	response = &user.Empty{}
	model := types.ApplyFor{
		UserID:  in.Id,
		BeId:    in.BeId,
		Message: in.Message,
		Type:    in.Type,
	}
	// 如果添加过message
	tmp := types.ApplyFor{}
	err = l.svcCtx.DB.Where("user_id = ? and be_id = ?", model.UserID, model.BeId).First(&tmp).Error
	if err != nil {
		err = l.svcCtx.DB.Create(&model).Error
		return
	}
	return response, l.svcCtx.DB.Where("user_id = ? and be_id = ?", model.UserID, model.BeId).Updates(&model).Error
}
