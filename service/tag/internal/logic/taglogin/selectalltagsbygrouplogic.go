package tagloginlogic

import (
	"Link/service/tag/internal/types"
	"context"

	"Link/service/tag/internal/svc"
	"Link/service/tag/tag"

	"github.com/zeromicro/go-zero/core/logx"
)

type SelectAllTagsByGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSelectAllTagsByGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SelectAllTagsByGroupLogic {
	return &SelectAllTagsByGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SelectAllTagsByGroupLogic) SelectAllTagsByGroup(in *tag.SelectAllTagsByGroupName) (*tag.AllTagsByGroupNameResponse, error) {
	// 根据标签组名称查询标签
	var lowTags []*tag.AllTags
	err := l.svcCtx.DB.Model(&types.Tag{}).Where("group_name = ? and type != 'OFFICIAL'", in.GroupAme).Find(&lowTags).Error
	return &tag.AllTagsByGroupNameResponse{
		LowTags: lowTags}, err
}
