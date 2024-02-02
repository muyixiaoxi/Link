package tagLogin

import (
	"context"

	"Link/restful/tag/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type SelectGourpTagLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSelectGourpTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SelectGourpTagLogic {
	return &SelectGourpTagLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SelectGourpTagLogic) SelectGourpTag() error {
	// todo: add your logic here and delete this line

	return nil
}
