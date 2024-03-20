package logic

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"user/common/jwt"
	"user/restful/internal/svc"
	"user/restful/internal/types"
	"user/service/user"

	"github.com/zeromicro/go-zero/core/logx"
)

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
		Phone:    req.Phone,
		Password: req.Password,
	})
	if err != nil {
		err = errors.New("用户名或密码错误")
		logc.Error(context.Background(), "l.svcCtx.UserRpc.UserLogin failed: ", err)
		return
	}
	jwts := jwt.InitNewJWTUtils()
	claims := jwt.UserClaims{
		UserId:   response.Id,
		Username: response.Username,
	}
	//返回双token
	accessTokenString, refreshTokenString, expireAt := jwts.GetToken(claims)
	if accessTokenString == "" || refreshTokenString == "" {
		return nil, status.Errorf(codes.FailedPrecondition, "token签发失败")
	}
	resp = &types.UserLoginResponse{
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
		Expires:      expireAt,
		Id:           response.Id,
		Username:     response.Username,
		Avatar:       response.Avatar,
	}
	return
}
