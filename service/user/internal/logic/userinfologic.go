package logic

import (
	"Link/service/user/internal/types"
	"context"

	"Link/service/user/internal/svc"
	"Link/service/user/user"

	"github.com/zeromicro/go-zero/core/logx"
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

func (l *UserInfoLogic) UserInfo(in *user.UserInfoRequest) (rp *user.UserInfoResponse, err error) {
	model := &types.User{}
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

	return
}
