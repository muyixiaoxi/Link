package userGroup

import (
	"context"
	"google.golang.org/grpc/status"
	"user/common/jwt"
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
	userID := l.ctx.Value(jwt.UserId).(uint64)
	//判断该用户加入 创建的群聊数是否超过了最大限制
	groupCount, err := l.svcCtx.UserRpc.SelectMyGroupCount(l.ctx, &user.SelectMyGroupCountRequest{UserId: userID})
	if groupCount.Count > 50 {
		return status.Error(899, "创建或者加入的群聊数据达到了最大限制")
	}
	// 封装请求参数
	createGroupParams := user.UserCreateGroupRequest{
		GroupBossId:   userID,
		Name:          req.Name,
		SystemTagId:   req.SystemTagId,
		UserSelfTagId: req.UserSelfTagId,
		Avatar:        req.Avatar,
	}
	_, err = l.svcCtx.UserRpc.UserCreateGroup(l.ctx, &createGroupParams)
	return err
}
