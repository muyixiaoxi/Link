package logic

import (
	"Link/internal/jwt"
	"Link/service/user/user"
	"context"

	"Link/restful/user/internal/svc"
	"Link/restful/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SignUpLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSignUpLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SignUpLogic {
	return &SignUpLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SignUpLogic) SignUp(req *types.UserCreateRequest) (resp *types.UserCreateResponse, err error) {
	response, err := l.svcCtx.UserRpc.UserCreate(l.ctx, &user.UserCreateRequest{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		logx.Errorf("SignUp failed: ", err)
		return
	}
	auth := l.svcCtx.Config.Auth
	claims := jwt.UserClaims{
		Username: response.Username,
		UserID:   response.Id,
	}
	token, _ := jwts.GenToken(claims, auth.AccessSecret, int64(response.Id))
	resp = &types.UserCreateResponse{
		Token:    token,
		Username: response.Username,
		Avatar:   response.Avatar,
	}
	return
}
