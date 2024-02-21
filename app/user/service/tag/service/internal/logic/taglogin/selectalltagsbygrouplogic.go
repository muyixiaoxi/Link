package tagloginlogic

import (
	"context"
	"user/service/tag/service/internal/svc"
	"user/service/tag/service/internal/types"
	"user/service/tag/service/tag"

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
