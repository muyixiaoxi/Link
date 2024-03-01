package tagloginlogic

import (
	"context"

	"user/service/tag/service/internal/svc"
	"user/service/tag/service/tag"

	"github.com/zeromicro/go-zero/core/logx"
)

type SelectMyTagsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSelectMyTagsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SelectMyTagsLogic {
	return &SelectMyTagsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SelectMyTagsLogic) SelectMyTags(in *tag.SelectMyTagsRequest) (*tag.AllTagsByGroupNameResponse, error) {
	// todo: add your logic here and delete this line

	return &tag.AllTagsByGroupNameResponse{}, nil
}
