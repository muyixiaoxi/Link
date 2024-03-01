package userGroup

import (
	"context"
	"user/service/user"

	"user/restful/internal/svc"
	"user/restful/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SelectDetailGroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSelectDetailGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SelectDetailGroupLogic {
	return &SelectDetailGroupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SelectDetailGroupLogic) SelectDetailGroup(req *types.SelectGroupDeatilRequest) (resp *types.SelectGroupDeatilResponse, err error) {
	// 查询群聊详情
	detailGroupRpc, err := l.svcCtx.UserRpc.UserSelectDetailGroup(l.ctx, &user.DetailGroupRequest{Id: req.ID})
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
	return
}
