package tagloginlogic

import (
	"context"

	"user/service/tag/service/internal/svc"
	"user/service/tag/service/tag"

	"github.com/zeromicro/go-zero/core/logx"
)

type SelectLinkTagsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSelectLinkTagsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SelectLinkTagsLogic {
	return &SelectLinkTagsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SelectLinkTagsLogic) SelectLinkTags(in *tag.SelectLinkTagsRequest) (*tag.SelectLinkTagsResponse, error) {
	// todo: add your logic here and delete this line

	return &tag.SelectLinkTagsResponse{}, nil
}
