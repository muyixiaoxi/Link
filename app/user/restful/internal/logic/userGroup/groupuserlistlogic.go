package userGroup

import (
	"context"
	"user/service/user"

	"user/restful/internal/svc"
	"user/restful/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupUserListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroupUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupUserListLogic {
	return &GroupUserListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupUserListLogic) GroupUserList(req *types.GroupUserListRequest) (resp *types.GroupUserListResponse, err error) {
	// 查询群聊中的人员信息
	groupUserListRpc, err := l.svcCtx.UserRpc.UserListByGroup(l.ctx, &user.SelectUserListByGroupRequest{GroupId: req.GroupID})
	var groupUserList []types.GroupUserList
	for _, groupUser := range groupUserListRpc.UserList {
		temp := types.GroupUserList{
			ID:       groupUser.Id,
			Avatar:   groupUser.Avatar,
			UserName: groupUser.Username,
		}
		groupUserList = append(groupUserList, temp)
	}
	return &types.GroupUserListResponse{
		GroupBossId:   groupUserListRpc.GroupBossId,
		GroupUserList: groupUserList,
	}, err
}
