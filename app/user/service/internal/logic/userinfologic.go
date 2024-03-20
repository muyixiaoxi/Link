package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"user/service/internal/svc"
	"user/service/internal/types"
	"user/service/user"
)

var gender = map[uint]string{0: "未知", 1: "男", 2: "女"}

type UserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UserInfo 用户信息
func (l *UserInfoLogic) UserInfo(in *user.UserInfoRequest) (rp *user.UserInfoResponse, err error) {
	rp = &user.UserInfoResponse{}
	model := &types.User{}
	id := in.UserId
	if in.FriendId != 0 {
		id = in.FriendId
	}
	cache, _ := l.svcCtx.RDB.Get(fmt.Sprintf("link:user:%d", id))
	if cache == "" {
		err = l.svcCtx.DB.Where("id = ?", id).First(model).Error
		if err != nil {
			return
		}
		rp = &user.UserInfoResponse{
			Id:        uint64(model.ID),
			Username:  model.Username,
			Avatar:    model.Avatar,
			Age:       uint64(model.Age),
			Gender:    uint64(model.Gender),
			Address:   model.Address,
			Phone:     model.Phone,
			Signature: model.Signature,
			History:   model.History,
		}
		js, _ := json.Marshal(rp)
		l.svcCtx.RDB.Setex(fmt.Sprintf("link:user:%d", model.ID), string(js), 12*60*60)
	} else {
		json.Unmarshal([]byte(cache), rp)
	}
	if in.FriendId != 0 {
		friend := &types.Friend{}
		l.svcCtx.DB.Where("user_id = ? and friend_id = ?", in.UserId, in.FriendId).First(friend)
		rp.IsFriend = friend.IsFriend
		rp.Remark = friend.Remark
	}

	return
}
