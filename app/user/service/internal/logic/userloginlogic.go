package logic

import (
	"context"
	"errors"
	"user/common/bcrypt"
	"user/service/internal/svc"
	"user/service/internal/types"
	"user/service/user"

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
	// 查找用户密码
	err = l.svcCtx.DB.Where("username = ? ", in.Username).Find(model).Error
	if err != nil {
		return
	}
	if ok := bcrypt.ComparePwd(model.Password, in.Password); !ok {
		err = errors.New("用户名或密码错误")
		return
	}
	pd = &user.UserLoginResponse{
		Id:       uint64(model.ID),
		Username: model.Username,
		Avatar:   model.Avatar,
	}
	return
}
