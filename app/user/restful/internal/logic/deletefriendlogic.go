package logic

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logc"
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
	jId, _ := l.ctx.Value("user_id").(json.Number)
	id, _ := jId.Int64()
	_, err = l.svcCtx.UserRpc.UserDeleteFriend(context.Background(), &user.UserDeleteFriendRequest{
		UserId:   uint64(id),
		FriendId: req.FriendId,
	})
	if err != nil {
		logc.Error(context.Background(), "l.svcCtx.UserRpc.UserDeleteFriend failed: ", err)
	}
	return
}
