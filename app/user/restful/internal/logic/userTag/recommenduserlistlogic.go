package userTag

import (
	"context"
	"encoding/json"
	"user/service/tag/service/tag"
	"user/service/user"

	"user/restful/internal/svc"
	"user/restful/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RecommendUserListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRecommendUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RecommendUserListLogic {
	return &RecommendUserListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RecommendUserListLogic) RecommendUserList(req *types.RecommendGroupByTagRequest) (resp *types.RecommendUserListResponse, err error) {
	// 根据标签推荐人员列表
	// 获取用户id
	//获取当前登录用户的id
	jid := l.ctx.Value("user_id").(json.Number)
	userId, _ := jid.Int64()
	var tagIds []uint64

	//如果用户没有选择标签
	if len(req.TagIds) == 0 {
		//查询出用户自己相关的标签
		respRpc, _ := l.svcCtx.TagLoginRpc.SelectLinkTags(l.ctx, &tag.SelectLinkTagsRequest{Id: uint64(userId)})
		for _, value := range respRpc.SelectLinkTags {
			//tagIds = append(tagIds, value.Id)
			for _, userSelf := range value.LinkTags {
				tagIds = append(tagIds, userSelf.Id)
			}
		}
	} else {
		tagIds = append(tagIds, req.TagIds...)
	}
	//查询人员列列表
	rpcUserList, err := l.svcCtx.UserRpc.RecommendUsers(l.ctx, &user.RecommendUsersRequest{
		PageNo:   int64(req.PageNo),
		PageSize: int64(req.PageSize),
		TagId:    tagIds,
		UserId:   uint64(userId),
	})
	var userList []types.RecommendUser
	for _, userRpc := range rpcUserList.RecommendUserList {
		temp := types.RecommendUser{
			Id:        userRpc.Id,
			Username:  userRpc.Username,
			Avatar:    userRpc.Avatar,
			Signature: userRpc.Signature,
			Type:      1, //1表示人员
		}
		userList = append(userList, temp)
	}
	resp = &types.RecommendUserListResponse{
		RecommendUserList: userList,
		Total:             rpcUserList.Total,
	}
	return
}
