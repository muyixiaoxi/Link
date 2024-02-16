package tagloginlogic

import (
	"Link/service/tag/internal/types"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"Link/service/tag/internal/svc"
	"Link/service/tag/tag"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateTagLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateTagLogic {
	return &UpdateTagLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateTagLogic) UpdateTag(in *tag.CreateTagRequest) (*tag.CreateTagResponse, error) {
	// todo: add your logic here and delete this line
	var newTag types.Tag
	err := l.svcCtx.DB.Where("creator_id = ? and tag_name = ?", in.CreatorId, in.OldTagName).First(&newTag).Error
	if err != nil {
		return nil, status.Error(codes.NotFound, "标签不存在")
	}
	newTag.TagName = in.TagName
	err = l.svcCtx.DB.Where("id = ?", newTag.ID).Updates(&newTag).Error
	if err != nil {
		return nil, err
	}
	return &tag.CreateTagResponse{}, nil
}
