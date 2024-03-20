package logic

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"user/common/jwt"

	"user/restful/internal/svc"
	"user/restful/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RefreshTokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRefreshTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RefreshTokenLogic {
	return &RefreshTokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RefreshTokenLogic) RefreshToken(req *types.RefreshTokenRequest) (resp *types.RefreshTokenResponse, err error) {
	// 无感刷新token
	// 提取refreshToken
	refreshToken := req.RefreshToken
	if refreshToken == "" {
		return nil, status.Error(codes.NotFound, "RefreshToken不存在")
	}
	jwts := jwt.InitNewJWTUtils()
	parseRefreshToken, isUpd, err := jwts.ParseRefreshToken(refreshToken)
	if err != nil {
		return nil, status.Errorf(codes.FailedPrecondition, "Token刷新失败")
	}
	if isUpd {
		accessTokenString, refreshTokenString, expiresAt := jwts.GetToken(parseRefreshToken.UserClaims)
		resp = &types.RefreshTokenResponse{
			AccessToken:  accessTokenString,
			RefreshToken: refreshTokenString,
			Expires:      expiresAt,
		}
	}
	return
}
