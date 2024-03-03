package userGroup

import (
	"context"
	"encoding/json"
	"user/service/user"

	"user/restful/internal/svc"
	"user/restful/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SelectMyGroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSelectMyGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SelectMyGroupLogic {
	return &SelectMyGroupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SelectMyGroupLogic) SelectMyGroup(req *types.SelectMyGroupRequest) (resp []types.GroupList, err error) {
	// 查询我加入的群聊
	//获取当前登录用户的id
	jid := l.ctx.Value("user_id").(json.Number)
	userId, _ := jid.Int64()
	//查询相关群聊
	respGroup, err := l.svcCtx.UserRpc.SearchMyGroupByName(l.ctx, &user.SearchMyGroupByNameRequest{
		UserId:       uint64(userId),
		RemarkOrName: req.Name,
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
			Remark:          groupInfo.Remark,
		}
		groupList = append(groupList, temp)
	}
	return groupList, err
}
