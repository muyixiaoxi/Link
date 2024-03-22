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
	userID := l.ctx.Value(jwt.UserId).(uint64)
	response, err := l.svcCtx.UserRpc.UserQueryFriend(context.Background(), &user.UserQueryFriendRequest{UserId: userID, Param: req.Param})
	if err != nil {
		logc.Error(context.Background(), "l.svcCtx.UserRpc.UserQueryFriend failed: ", err)
		return
	}
	for _, u := range response.List {
		resp.Friends = append(resp.Friends, types.UserFriend{
			Id:        u.Id,
			Username:  u.Username,
			Avatar:    u.Avatar,
			Remark:    u.Remark,
			Signature: u.Signature,
			TagName:   u.TagName,
			Type:      1, //人
		})
	}
	return
}
