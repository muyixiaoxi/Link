package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"user/common/jwt"
	"user/restful/internal/svc"
	"user/restful/internal/types"
	"user/service/user"
)

type GetFriendsListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetFriendsListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFriendsListLogic {
	return &GetFriendsListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetFriendsListLogic) GetFriendsList() (resp *types.UserFriendsResponse, err error) {
	userID := l.ctx.Value(jwt.UserId).(uint64)
	response, err := l.svcCtx.UserRpc.UserFriendList(context.Background(), &user.UserFriendRequest{
		Id: userID,
	})
	resp = &types.UserFriendsResponse{}
	for _, u := range response.List {
		resp.Friends = append(resp.Friends, types.UserFriend{
			Id:        u.Id,
			Username:  u.Username,
			Avatar:    u.Avatar,
			Remark:    u.Remark,
			Signature: u.Signature,
			TagName:   u.TagName,
			Type:      1, //好友
		})
	}

	return resp, err
}
