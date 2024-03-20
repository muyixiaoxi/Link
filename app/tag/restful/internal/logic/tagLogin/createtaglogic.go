package tagLogin

import (
	"context"
	"google.golang.org/grpc/status"
	"tag/common/jwt"
	"tag/restful/internal/svc"
	"tag/restful/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"tag/service/tag"
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
	// 用户创建标签
	userID := l.ctx.Value(jwt.UserId).(uint64)
	// 限制每个用户最多有9个标签
	tagCount, err := l.svcCtx.TagLogin.CheckTagCount(l.ctx, &tag.CheckTagCountRequest{UserId: userID})
	if tagCount.Count > 9 {
		return nil, status.Error(899, "标签数量超过最大数量限制")
	}
	//封装请求参数
	createTagParams := &tag.CreateTagRequest{
		CreatorId: userID,
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
