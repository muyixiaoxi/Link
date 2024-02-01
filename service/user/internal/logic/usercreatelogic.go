package logic

import (
	"Link/internal/bcrypt"
	"Link/service/user/internal/types"
	"context"
	"errors"

	"Link/service/user/internal/svc"
	"Link/service/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserCreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserCreateLogic {
	return &UserCreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserCreateLogic) UserCreate(in *user.UserCreateRequest) (pd *user.UserCreateResponse, err error) {
	// 判断用户是否存在
	err = l.svcCtx.DB.Take("username = ?", in.Username).Error
	if err == nil {
		return nil, errors.New("用户存在")
	}
	// 加密
	pwd, _ := bcrypt.GetPwd(in.Password)
	model := types.User{
		Username: in.Username,
		Password: pwd,
	}
	err = l.svcCtx.DB.Create(model).Error
	if err != nil {
		return
	}
	pd = &user.UserCreateResponse{
		Id:       uint64(model.ID),
		Username: model.Username,
	}
	return &user.UserCreateResponse{}, nil
}
