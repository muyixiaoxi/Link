package tagLogin

import (
	"Link/restful/tag/internal/svc"
	"Link/restful/tag/internal/types"
	"Link/service/tag/tag"
	"context"
	"encoding/json"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateTagLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateTagLogic {
	return &CreateTagLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateTagLogic) CreateTag(req *types.CreateTagRequest) (resp *types.CreateTagResponse, err error) {
	// todo: add your logic here and delete this line
	userId := l.ctx.Value("user_id").(json.Number)
	id, _ := userId.Int64()
	//封装请求参数
	createTagParams := &tag.CreateTagRequest{
		CreatorId: uint64(id),
		TagName:   req.TagName,
		GroupName: req.GroupName,
		Type:      req.Type,
	}
	responseBody, err := l.svcCtx.TagLogin.CreateTag(l.ctx, createTagParams)
	if err != nil {
		return nil, err
	}
	// 转换 LowTags 切片为 TagGroupName 切片
	var tagGroupNames []types.TagGroupName
	for _, lowTag := range responseBody.LowTags {
		tagGroupNames = append(tagGroupNames, types.TagGroupName{
			GroupNameId: int(lowTag.Id),
			GroupName:   lowTag.TagName,
			Creator:     int(lowTag.CreatorId),
		})
	}

	resp = &types.CreateTagResponse{
		GroupName:    responseBody.GroupName,
		TagGroupName: tagGroupNames,
	}
	return
}
