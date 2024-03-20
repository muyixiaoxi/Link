package userGroup

import (
	"context"
	"user/common/jwt"
	"user/service/user"

	"user/restful/internal/svc"
	"user/restful/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type KickOutUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewKickOutUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *KickOutUserLogic {
	return &KickOutUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *KickOutUserLogic) KickOutUser(req *types.KickOutGroupRequest) error {
	// 踢出群聊
	// 获取当前登录用户的id
	userID := l.ctx.Value(jwt.UserId).(uint64)
	_, err := l.svcCtx.UserRpc.KickOutUserGroup(l.ctx, &user.KickOutUserGroupRequest{
		UserId:         userID,
		GroupId:        req.GroupId,
		KickOutUsersId: req.UserIDs,
	})
	return err
}
