package logic

import (
	"Link/internal/bcrypt"
	"Link/service/user/internal/svc"
	"Link/service/user/internal/types"
	"Link/service/user/user/user"
	"context"
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

func (l *UserCreateLogic) UserCreate(in *user.UserCreateRequest) (pd *user.UserCreateResponse, error) {
	pd = new(user.UserCreateResponse)
	// 判断用户是否存在
	model := types.User{}
	err := l.svcCtx.DB.Take(&model, "username = ?", in.Username).Error
	if err == nil {
		pd.Err = "该用户已存在"
	}
	// 注册
	pwd, err := bcrypt.GetPwd(in.Password)
	if err != nil {
		return nil, err
	}
	model = types.User{
		Username: in.Username,
		Password: pwd,
	}
	l.svcCtx.DB.Create(model)

	return &user.UserCreateResponse{}, nil
}
