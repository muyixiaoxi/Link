package tagloginlogic

import (
	"context"

	"user/service/tag/service/internal/svc"
	"user/service/tag/service/tag"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChooseTagsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewChooseTagsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChooseTagsLogic {
	return &ChooseTagsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ChooseTagsLogic) ChooseTags(in *tag.ChooseTagsRequest) (*tag.ChooseTagsResponse, error) {
	// todo: add your logic here and delete this line

	return &tag.ChooseTagsResponse{}, nil
}
