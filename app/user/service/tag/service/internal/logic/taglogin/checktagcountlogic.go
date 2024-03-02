package tagloginlogic

import (
	"context"

	"user/service/tag/service/internal/svc"
	"user/service/tag/service/tag"

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
	// todo: add your logic here and delete this line

	return &tag.CheckTagCountResponse{}, nil
}
