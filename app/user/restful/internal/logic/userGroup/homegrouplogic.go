package userGroup

import (
	"context"
	"encoding/json"
	"user/restful/internal/svc"
	"user/restful/internal/types"
	"user/service/tag/service/tag"
	"user/service/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type HomeGroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomeGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomeGroupLogic {
	return &HomeGroupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomeGroupLogic) HomeGroup(req *types.RecommendGroupByTagRequest) (resp *types.RecommendGroupByTagResponse, err error) {
	// 首页推荐群聊
	//获取当前登录用户的id
	jid := l.ctx.Value("user_id").(json.Number)
	userId, _ := jid.Int64()
	var (
		selfIds   []uint64
		systemIds []uint64
	)
	//如果用户没有选择标签
	if len(req.UserSelfTagId) == 0 && len(req.SystemTagId) == 0 {
		//查询出用户自己相关的标签
		respRpc, _ := l.svcCtx.TagLoginRpc.SelectLinkTags(l.ctx, &tag.SelectLinkTagsRequest{Id: uint64(userId)})
		for _, value := range respRpc.LinkTags {
			if value.CreatorId != 0 {
				//我加入或者创建的标签
				selfIds = append(selfIds, value.Id)
			} else {
				systemIds = append(systemIds, value.Id)
			}
		}
	} else {
		systemIds = append(systemIds, req.SystemTagId...)
		selfIds = append(selfIds, req.UserSelfTagId...)
	}
	//查询相关群聊
	respGroup, err := l.svcCtx.UserRpc.UserSelectGroup(l.ctx, &user.UserSelectGroupsRequest{
		SystemTagId:   systemIds,
		UserSelfTagId: selfIds,
		PageNo:        req.PageNo,
		PageSize:      req.PageSize,
	})
	if err != nil {
		return nil, err
	}
	var groupList []types.GroupList
	for _, groupInfo := range respGroup.GroupList {
		temp := types.GroupList{
			Id:              groupInfo.Id,
			Name:            groupInfo.Name,
			SystemTagName:   groupInfo.SystemTagName,
			UserSelfTagName: groupInfo.UserSelfTagName,
			Avatar:          groupInfo.Avatar,
		}
		groupList = append(groupList, temp)
	}
	return &types.RecommendGroupByTagResponse{
		GroupList: groupList,
		Total:     int64(respGroup.Total),
	}, nil
}
