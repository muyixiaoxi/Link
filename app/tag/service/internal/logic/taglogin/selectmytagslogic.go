package tagloginlogic

import (
	"context"
	"tag/service/internal/svc"
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
	err := l.svcCtx.DB.Table("tb_tag").
		Joins("LEFT JOIN tb_user_tag ON tb_tag.id = tb_user_tag.tag_id AND tb_tag.type != 'OFFICIAL' AND tb_tag.creator_id = ? WHERE tb_user_tag.user_id = ?", in.UserId, in.UserId).
		Select("tb_tag.*").
		Scan(&lowTags).Error
	return &tag.AllTagsByGroupNameResponse{
		LowTags: lowTags}, err
}
