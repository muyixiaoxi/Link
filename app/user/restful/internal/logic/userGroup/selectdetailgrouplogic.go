package userGroup

import (
	"context"
	"encoding/json"
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
	//获取当前登录用户的id
	jid := l.ctx.Value("user_id").(json.Number)
	userId, _ := jid.Int64()
	detailGroupRpc, err := l.svcCtx.UserRpc.UserSelectDetailGroup(l.ctx, &user.DetailGroupRequest{Id: req.ID, UserId: uint64(userId)})
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
		SystemTagName:   detailGroupRpc.SystemTagName,
		AddressCount:    detailGroupRpc.AddressCount,
		IsExist:         detailGroupRpc.IsExit,
	}
	return
}
