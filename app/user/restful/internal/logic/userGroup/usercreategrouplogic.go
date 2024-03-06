package userGroup

import (
	"context"
	"encoding/json"
	"google.golang.org/grpc/status"
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
	//判断该用户加入 创建的群聊数是否超过了最大限制
	groupCount, err := l.svcCtx.UserRpc.SelectMyGroupCount(l.ctx, &user.SelectMyGroupCountRequest{UserId: uint64(userId)})
	if groupCount.Count > 50 {
		return status.Error(899, "创建或者加入的群聊数据达到了最大限制")
	}
	// 封装请求参数
	createGroupParams := user.UserCreateGroupRequest{
		GroupBossId:   uint64(userId),
		Name:          req.Name,
		SystemTagId:   req.SystemTagId,
		UserSelfTagId: req.UserSelfTagId,
		Avatar:        req.Avatar,
	}
	_, err = l.svcCtx.UserRpc.UserCreateGroup(l.ctx, &createGroupParams)
	return err
}
