package userGroup

import (
	"context"
	"encoding/json"
	"user/restful/internal/svc"
	"user/restful/internal/types"
	"user/service/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserCreateGroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserCreateGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserCreateGroupLogic {
	return &UserCreateGroupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserCreateGroupLogic) UserCreateGroup(req *types.UserCreateGroupRequset) error {
	//获取当前登录用户的id
	jid := l.ctx.Value("user_id").(json.Number)
	userId, _ := jid.Int64()
	// 封装请求参数
	createGroupParams := user.UserCreateGroupRequest{
		GroupBossId:   uint64(userId),
		Name:          req.Name,
		SystemTagId:   req.SystemTagId,
		UserSelfTagId: req.UserSelfTagId,
		Avatar:        req.Avatar,
	}
	_, err := l.svcCtx.UserRpc.UserCreateGroup(l.ctx, &createGroupParams)
	return err
}
