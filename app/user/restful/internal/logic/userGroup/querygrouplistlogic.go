package userGroup

import (
	"context"
	"user/common/jwt"
	"user/restful/internal/svc"
	"user/restful/internal/types"
	"user/service/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryGroupListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryGroupListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryGroupListLogic {
	return &QueryGroupListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryGroupListLogic) QueryGroupList(req *types.QueryGroupListRequest) (resp *types.RecommendGroupByTagResponse, err error) {
	// 查询当前登录用户的群列表
	// 获取当前登录用户的id
	userID := l.ctx.Value(jwt.UserId).(uint64)
	//查询相关群聊
	respGroup, err := l.svcCtx.UserRpc.QueryMyGroupList(l.ctx, &user.QueryMyGroupListRequest{
		UserId:   userID,
		PageNo:   req.PageNo,
		PageSize: req.PageSize,
	})
	if err != nil {
		return nil, err
	}
	var groupList []types.GroupList
	for _, groupInfo := range respGroup.GroupList {
		temp := types.GroupList{
			Id:              groupInfo.Id,
			Type:            2,
			Name:            groupInfo.Name,
			SystemTagName:   groupInfo.SystemTagName,
			UserSelfTagName: groupInfo.UserSelfTagName,
			Avatar:          groupInfo.Avatar,
			Remark:          groupInfo.Remark,
		}
		groupList = append(groupList, temp)
	}
	return &types.RecommendGroupByTagResponse{
		GroupList: groupList,
		Total:     int64(respGroup.Total),
	}, nil

}
