package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logc"
	"user/common/jwt"
	"user/service/user"

	"user/restful/internal/svc"
	"user/restful/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteFriendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteFriendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteFriendLogic {
	return &DeleteFriendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteFriendLogic) DeleteFriend(req *types.UserDeleteFriendRequest) (err error) {
	userID := l.ctx.Value(jwt.UserId).(uint64)
	_, err = l.svcCtx.UserRpc.UserDeleteFriend(context.Background(), &user.UserDeleteFriendRequest{
		UserId:   userID,
		FriendId: req.FriendId,
	})
	if err != nil {
		logc.Error(context.Background(), "l.svcCtx.UserRpc.UserDeleteFriend failed: ", err)
	}
	return
}
