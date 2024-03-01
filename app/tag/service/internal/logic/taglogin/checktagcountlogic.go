package tagloginlogic

import (
	"context"
	"tag/service/internal/types"

	"tag/service/internal/svc"
	"tag/service/tag"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckTagCountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckTagCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckTagCountLogic {
	return &CheckTagCountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckTagCountLogic) CheckTagCount(in *tag.CheckTagCountRequest) (*tag.CheckTagCountResponse, error) {
	var totalCount int64
	// 检查当前用户选择或者创建的标签数量
	if err := l.svcCtx.DB.Model(&types.UserTagFollow{}).Where("user_id = ?", in.UserId).Count(&totalCount).Error; err != nil {
		return nil, err
	}
	return &tag.CheckTagCountResponse{
		Count: uint64(totalCount),
	}, nil
}
