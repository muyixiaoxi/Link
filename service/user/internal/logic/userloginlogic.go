package logic

import (
	"Link/service/user/internal/types"
	"context"

	"Link/service/user/internal/svc"
	"Link/service/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserLoginLogic) UserLogin(in *user.UserLoginRequest) (pd *user.UserLoginResponse, err error) {
	model := &types.User{}
	err = l.svcCtx.DB.Where("username = ? and password = ?", in.Username, in.Password).Find(model).Error
	if err != nil {
		return
	}
	pd = &user.UserLoginResponse{
		Id:       uint64(model.ID),
		Username: model.Username,
	}
	return
}
