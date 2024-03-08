package logic

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"user/common/bcrypt"
	"user/service/internal/svc"
	"user/service/internal/types"
	"user/service/user"

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
	id := in.Id
	m := &types.User{}
	l.svcCtx.DB.Where("phone = ?", in.Phone).First(m)
	if m.ID != 0 && m.ID != uint(id) {
		err = errors.New("该手机号已被注册")
		return nil, err
	}
	l.svcCtx.RDB.Del(fmt.Sprintf("link:user:%d", m.ID))
	var pwd string
	if len(in.Password) != 0 {
		pwd, _ = bcrypt.GetPwd(in.Password)
	}
	model := &types.User{
		Model:     gorm.Model{ID: uint(in.Id)},
		Username:  in.Username,
		Password:  pwd,
		Age:       uint(in.Age),
		Gender:    uint(in.Gender),
		Address:   in.Address,
		Phone:     in.Phone,
		Avatar:    in.Avatar,
		Signature: in.Signature,
	}
	err = l.svcCtx.DB.Updates(model).Error
	model.Password = ""
	js, _ := json.Marshal(model)
	if err == nil {
		l.svcCtx.RDB.Set(fmt.Sprintf("link:user:%d", model.ID), string(js))
	}
	return
}
