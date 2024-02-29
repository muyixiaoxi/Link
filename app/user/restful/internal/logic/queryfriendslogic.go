package logic

import (
	"context"
	"encoding/json"
	"user/service/user"

	"user/restful/internal/svc"
	"user/restful/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryFriendsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryFriendsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryFriendsLogic {
	return &QueryFriendsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryFriendsLogic) QueryFriends(req *types.UserQueryFriendRequest) (resp *types.UserFriendsResponse, err error) {
	resp = &types.UserFriendsResponse{}
	jid := l.ctx.Value("user_id").(json.Number)
	userId, _ := jid.Int64()
	response, err := l.svcCtx.UserRpc.UserQueryFriend(context.Background(), &user.UserQueryFriendRequest{UserId: uint64(userId), Param: req.Param})

	for _, u := range response.List {
		resp.Friends = append(resp.Friends, types.UserFriend{
			Id:        u.Id,
			Username:  u.Username,
			Avatar:    u.Avatar,
			Remark:    u.Remark,
			Signature: u.Signature,
			TagName:   u.TagName,
		})
	}
	return
}
