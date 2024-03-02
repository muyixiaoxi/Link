package tagloginlogic

import (
	"context"

	"user/service/tag/service/internal/svc"
	"user/service/tag/service/tag"

	"github.com/zeromicro/go-zero/core/logx"
)

type CancelUserTagLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCancelUserTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CancelUserTagLogic {
	return &CancelUserTagLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CancelUserTagLogic) CancelUserTag(in *tag.CancelRequest) (*tag.Empty, error) {
	// todo: add your logic here and delete this line

	return &tag.Empty{}, nil
}
