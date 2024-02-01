package logic

import (
	"context"

	"Link/restful/tag/internal/svc"
	"Link/restful/tag/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TagLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TagLogic {
	return &TagLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TagLogic) Tag(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
