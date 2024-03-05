package userGroup

import (
	"context"
	"encoding/json"
	"user/service/user"

	"user/restful/internal/svc"
	"user/restful/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SelectStrangerGroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSelectStrangerGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SelectStrangerGroupLogic {
	return &SelectStrangerGroupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SelectStrangerGroupLogic) SelectStrangerGroup(req *types.SelectStrangerRequest) (resp *types.RecommendGroupByTagResponse, err error) {
	// 查询陌生群聊
	//获取当前登录用户的id
	jid := l.ctx.Value("user_id").(json.Number)
	userId, _ := jid.Int64()

	respGroup, err := l.svcCtx.UserRpc.SearchStrangerGroup(l.ctx, &user.SearchStrangerGroupRequest{
		UserId:   uint64(userId),
		PageNo:   req.PageNo,
		PageSize: req.PageSize,
		Name:     req.Name,
	})
	if err != nil {
		return nil, err
	}
	var groupList []types.GroupList
	for _, groupInfo := range respGroup.GroupList {
		temp := types.GroupList{
			Id:              groupInfo.Id,
			Type:            2, //2表示群聊
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
