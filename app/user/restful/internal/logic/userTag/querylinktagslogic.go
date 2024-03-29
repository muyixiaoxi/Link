package userTag

import (
	"context"
	"user/common/jwt"
	"user/service/tag/service/tag"

	"user/restful/internal/svc"
	"user/restful/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryLinkTagsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryLinkTagsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryLinkTagsLogic {
	return &QueryLinkTagsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryLinkTagsLogic) QueryLinkTags(req *types.UserInfoRequest) (resp []types.QueryLinkTagsResponse, err error) {
	var userId uint64
	var isSelf bool
	if req.Id == 0 {
		// 查询和当前登录用户相关的标签
		userId = l.ctx.Value(jwt.UserId).(uint64)
		isSelf = true //查询的是自己的标签
	} else {
		userId = req.Id
		isSelf = false
	}
	//查询
	rpcResp, err := l.svcCtx.TagLoginRpc.SelectLinkTags(l.ctx, &tag.SelectLinkTagsRequest{Id: userId, IsSelf: isSelf})
	if rpcResp == nil {
		return
	}
	for _, value := range rpcResp.SelectLinkTags {
		var linkTags []types.QueryLink
		for _, linkTag := range value.GetLinkTags() {
			mid := types.QueryLink{
				Id:        linkTag.Id,
				CreatorId: linkTag.CreatorId,
				TagName:   linkTag.TagName,
			}
			linkTags = append(linkTags, mid)
		}
		temp := types.QueryLinkTagsResponse{
			SystemTagName: value.GroupName,
			SystemTagId:   value.GroupId,
			LinkTags:      linkTags,
		}
		resp = append(resp, temp)
	}
	return
}
