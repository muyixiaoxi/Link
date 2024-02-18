package userservicelogic

import (
	"Link/internal/bcrypt"
	"Link/service/user/internal/types"
	"context"
	"errors"
	"gorm.io/gorm"

	"Link/service/user/internal/svc"
	"Link/service/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserUpdateInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserUpdateInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserUpdateInfoLogic {
	return &UserUpdateInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserUpdateInfoLogic) UserUpdateInfo(in *user.UserUpdateInfoRequest) (rp *user.Empty, err error) {
	rp = &user.Empty{}
	// 判断用户是否存在
	m := &types.User{}
	l.svcCtx.DB.Where("username = ?", in.Username).First(m)
	if m.ID != 0 {
		err = errors.New("用户已存在")
		return nil, err
	}
	pwd, _ := bcrypt.GetPwd(in.Password)
	model := &types.User{
		Model:    gorm.Model{ID: uint(in.Id)},
		Username: in.Username,
		Password: pwd,
		Age:      uint(in.Age),
		Gender:   uint(in.Gender),
		Address:  in.Address,
		Phone:    in.Phone,
		Avatar:   in.Avatar,
	}
	err = l.svcCtx.DB.Updates(model).Error
	return
}
