package logic

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
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
	jId := l.ctx.Value("user_id").(json.Number)
	id, _ := jId.Int64()
	response, err := l.svcCtx.UserRpc.UserFriendList(context.Background(), &user.UserFriendRequest{
		Id: uint64(id),
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
