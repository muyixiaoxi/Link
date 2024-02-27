package tagLogin

import (
	"context"
	"tag/service/tag"

	"tag/restful/internal/svc"
	"tag/restful/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SelectAllTagsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSelectAllTagsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SelectAllTagsLogic {
	return &SelectAllTagsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SelectAllTagsLogic) SelectAllTags() (resp []types.SelectAllTagsResponse, err error) {
	// 一次直接查询所有标签
	// 首先查询出系统标签
	empty := &tag.Empty{}
	respBody, err := l.svcCtx.TagLogin.SelectGroupTag(l.ctx, empty)
	if err != nil {
		return nil, err
	}
	for _, value := range respBody.Tags {
		//根据系统标签查询自定义标签
		tagsResp, err := l.svcCtx.TagLogin.SelectAllTagsByGroup(l.ctx, &tag.SelectAllTagsByGroupName{GroupAme: value.GroupName})
		if err != nil {
			return nil, err
		}
		var temp []types.Tag
		for _, tags := range tagsResp.LowTags {
			temp = append(temp, types.Tag{
				TagId:     tags.Id,
				CreatorId: tags.CreatorId,
				TagName:   tags.TagName,
			})
		}
		allTag := types.SelectAllTagsResponse{
			GroupNameId: int64(value.Id),
			GroupName:   value.GroupName,
			Tags:        temp,
		}
		resp = append(resp, allTag)
	}
	return
}
