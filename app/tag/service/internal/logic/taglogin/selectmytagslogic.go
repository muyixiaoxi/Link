package tagloginlogic

import (
	"context"
	"tag/service/internal/svc"
	"tag/service/internal/types"
	"tag/service/tag"

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
	// 查询我创建的标签
	var lowTags []*tag.AllTags
	err := l.svcCtx.DB.Model(&types.Tag{}).Where("creator_id = ?", in.UserId).Find(&lowTags).Error
	return &tag.AllTagsByGroupNameResponse{
		LowTags: lowTags}, err
}
