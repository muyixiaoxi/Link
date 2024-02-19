package tagLogin

import (
	"context"
	"tag/restful/internal/svc"
	"tag/restful/internal/types"

	"tag/service/tag"

	"github.com/zeromicro/go-zero/core/logx"
)

type SelectUserTagByGroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSelectUserTagByGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SelectUserTagByGroupLogic {
	return &SelectUserTagByGroupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SelectUserTagByGroupLogic) SelectUserTagByGroup(req *types.SelectUserTagByGroupRequest) (resp *types.SelectUserTagByGroupResponse, err error) {
	// todo: add your logic here and delete this line
	respBody, err := l.svcCtx.TagLogin.SelectAllTagsByGroup(l.ctx, &tag.SelectAllTagsByGroupName{GroupAme: req.GroupName})
	if err != nil {
		return nil, err
	}
	var tags []types.Tag
	for _, value := range respBody.LowTags {
		tags = append(tags, types.Tag{
			TagId:     value.Id,
			CreatorId: value.CreatorId,
			TagName:   value.TagName,
		})
	}
	resp = &types.SelectUserTagByGroupResponse{
		GroupName: req.GroupName,
		Tags:      tags,
	}
	return
}
