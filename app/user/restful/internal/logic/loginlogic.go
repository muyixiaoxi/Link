package logic

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logc"
	"user/common/jwt"
	"user/restful/internal/svc"
	"user/restful/internal/types"
	"user/service/user"

	"github.com/zeromicro/go-zero/core/logx"
)

var jwts jwt.JWT

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// Login 用户登录
func (l *LoginLogic) Login(req *types.UserLoginRequest) (resp *types.UserLoginResponse, err error) {
	response, err := l.svcCtx.UserRpc.UserLogin(l.ctx, &user.UserLoginRequest{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		err = errors.New("用户名或密码错误")
		logc.Error(context.Background(), "l.svcCtx.UserRpc.UserLogin failed: ", err)
		return
	}
	claims := jwt.UserClaims{
		UserID:   response.Id,
		Username: response.Username,
	}
	auth := l.svcCtx.Config.Auth
	token, _ := jwts.GenToken(claims, auth.AccessSecret, int64(response.Id))
	resp = &types.UserLoginResponse{
		Token:    token,
		Username: response.Username,
		Avatar:   response.Avatar,
	}
	return
}
