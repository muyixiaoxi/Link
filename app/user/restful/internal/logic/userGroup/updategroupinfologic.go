package userGroup

import (
	"context"
	"encoding/json"
	"user/service/user"

	"user/restful/internal/svc"
	"user/restful/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateGroupInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateGroupInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateGroupInfoLogic {
	return &UpdateGroupInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateGroupInfoLogic) UpdateGroupInfo(req *types.UpdateGroupInfoRequest) (resp *types.SelectGroupDeatilResponse, err error) {
	// 获取当前登录用户的id
	jUserId := l.ctx.Value("user_id").(json.Number)
	userId, _ := jUserId.Int64()
	// 修改群聊信息
	_, err = l.svcCtx.UserRpc.UpdateGroupInformation(l.ctx, &user.UpdateGroupInfoRequest{
		Id:            req.GroupId,
		UserId:        uint64(userId),
		SystemTagId:   req.SystemTagId,
		UserSelfTagId: req.UserSelfTagId,
		Name:          req.Name,
		Avatar:        req.Avatar,
	})
	if err != nil {
		return nil, err
	}
	//查询修改完的群聊信息
	detailGroupRpc, err := l.svcCtx.UserRpc.UserSelectDetailGroup(l.ctx, &user.DetailGroupRequest{Id: req.GroupId})
	if err != nil {
		return nil, err
	}
	createdAt := detailGroupRpc.CreatedAt.String()
	resp = &types.SelectGroupDeatilResponse{
		ID:              detailGroupRpc.Id,
		Count:           detailGroupRpc.Count,
		GroupBossId:     detailGroupRpc.GroupBossId,
		Name:            detailGroupRpc.Name,
		UserSelfTagName: detailGroupRpc.UserSelfTagName,
		Avatar:          detailGroupRpc.Avatar,
		CreatedAt:       createdAt,
		Man:             detailGroupRpc.Man,
		WoMan:           detailGroupRpc.Woman,
		Address:         detailGroupRpc.Address,
	}
	if err != nil {
		return nil, err
	}
	return resp, nil
}
