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
	cache, _ := l.svcCtx.RDB.Get(fmt.Sprintf("user:%d", in.Id))
	if cache == "" {
		err = l.svcCtx.DB.Where("id = ?", in.Id).First(model).Error
		if err != nil {
			return
		}
		rp = &user.UserInfoResponse{
			Id:       uint64(model.ID),
			Username: model.Username,
			Avatar:   model.Avatar,
			Age:      uint64(model.Age),
			Gender:   gender[model.Gender],
			Address:  model.Address,
			Phone:    model.Phone,
		}
		js, _ := json.Marshal(rp)
		l.svcCtx.RDB.Set(fmt.Sprintf("user:%d", model.ID), string(js))
		return
	}
	json.Unmarshal([]byte(cache), rp)
	return
}
