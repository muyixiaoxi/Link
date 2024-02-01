package tagLogin

import (
	"Link/restful/tag/internal/svc"
	"Link/restful/tag/internal/types"
	"Link/service/tag/tag"
	"context"

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
	//封装请求参数
	createTagParams := &tag.CreateTagRequest{
		CreatorId: req.CreatorId,
		TagName:   req.TagName,
		GroupName: req.GroupName,
		Type:      req.Type,
	}
	responseBody, err := l.svcCtx.TagLogin.CreateTag(l.ctx, createTagParams)
	if err != nil {
		return nil, err
	}
	resp = &types.CreateTagResponse{
		GroupName: responseBody.GroupName,
		TagNames:  responseBody.TagName,
	}
	return
}
