// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package server

import (
	"context"

	"Link/service/user/internal/logic"
	"Link/service/user/internal/svc"
	"Link/service/user/user/user"
)

type UserServer struct {
	svcCtx *svc.ServiceContext
	user.UnimplementedUserServer
}

func NewUserServer(svcCtx *svc.ServiceContext) *UserServer {
	return &UserServer{
		svcCtx: svcCtx,
	}
}

func (s *UserServer) Ping(ctx context.Context, in *user.Request) (*user.Response, error) {
	l := logic.NewPingLogic(ctx, s.svcCtx)
	return l.Ping(in)
}

func (s *UserServer) UserCreate(ctx context.Context, in *user.UserCreateRequest) (*user.UserCreateResponse, error) {
	l := logic.NewUserCreateLogic(ctx, s.svcCtx)
	return l.UserCreate(in)
}

func (s *UserServer) UserLogin(ctx context.Context, in *user.UserLoginRequest) (*user.UserLoginResponse, error) {
	l := logic.NewUserLoginLogic(ctx, s.svcCtx)
	return l.UserLogin(in)
}
