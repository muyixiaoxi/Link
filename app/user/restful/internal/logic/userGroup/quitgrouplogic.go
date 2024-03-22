package userGroup

import (
	"context"
	"user/common/jwt"
	"user/service/user"

	"user/restful/internal/svc"
	"user/restful/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QuitGroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQuitGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QuitGroupLogic {
	return &QuitGroupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QuitGroupLogic) QuitGroup(req *types.QuitGroupRequest) error {
	// 解散群聊
	// 获取当前登录用户的id
	userID := l.ctx.Value(jwt.UserId).(uint64)
	//调用rpc服务
	_, err := l.svcCtx.UserRpc.QuitGroup(l.ctx, &user.QuitGroupRequest{
		UserId:  userID,
		GroupId: req.GroupId,
	})
	if err != nil {
		return err
	}
	return nil
}
