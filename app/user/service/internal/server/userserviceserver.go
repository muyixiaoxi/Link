// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package server

import (
	"context"

	"user/service/internal/logic"
	"user/service/internal/svc"
	"user/service/user"
)

type UserServiceServer struct {
	svcCtx *svc.ServiceContext
	user.UnimplementedUserServiceServer
}

func NewUserServiceServer(svcCtx *svc.ServiceContext) *UserServiceServer {
	return &UserServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *UserServiceServer) UserCreate(ctx context.Context, in *user.UserCreateRequest) (*user.UserCreateResponse, error) {
	l := logic.NewUserCreateLogic(ctx, s.svcCtx)
	return l.UserCreate(in)
}

func (s *UserServiceServer) UserCreateRevertLogin(ctx context.Context, in *user.UserCreateRequest) (*user.UserCreateResponse, error) {
	l := logic.NewUserCreateRevertLoginLogic(ctx, s.svcCtx)
	return l.UserCreateRevertLogin(in)
}

func (s *UserServiceServer) UserIsExists(ctx context.Context, in *user.UserCreateRequest) (*user.Empty, error) {
	l := logic.NewUserIsExistsLogic(ctx, s.svcCtx)
	return l.UserIsExists(in)
}

func (s *UserServiceServer) UserLogin(ctx context.Context, in *user.UserLoginRequest) (*user.UserLoginResponse, error) {
	l := logic.NewUserLoginLogic(ctx, s.svcCtx)
	return l.UserLogin(in)
}

func (s *UserServiceServer) UserInfo(ctx context.Context, in *user.UserInfoRequest) (*user.UserInfoResponse, error) {
	l := logic.NewUserInfoLogic(ctx, s.svcCtx)
	return l.UserInfo(in)
}

func (s *UserServiceServer) UserUpdateInfo(ctx context.Context, in *user.UserUpdateInfoRequest) (*user.Empty, error) {
	l := logic.NewUserUpdateInfoLogic(ctx, s.svcCtx)
	return l.UserUpdateInfo(in)
}

// 好友
func (s *UserServiceServer) UserFlowed(ctx context.Context, in *user.UserAddRequest) (*user.Empty, error) {
	l := logic.NewUserFlowedLogic(ctx, s.svcCtx)
	return l.UserFlowed(in)
}

func (s *UserServiceServer) UserDisposeFlowed(ctx context.Context, in *user.DisposeFlowedRequest) (*user.Empty, error) {
	l := logic.NewUserDisposeFlowedLogic(ctx, s.svcCtx)
	return l.UserDisposeFlowed(in)
}

func (s *UserServiceServer) UserUpdateRemark(ctx context.Context, in *user.UserUpdateRemarkRequest) (*user.Empty, error) {
	l := logic.NewUserUpdateRemarkLogic(ctx, s.svcCtx)
	return l.UserUpdateRemark(in)
}

func (s *UserServiceServer) UserFriendList(ctx context.Context, in *user.UserFriendRequest) (*user.UserFriendResponse, error) {
	l := logic.NewUserFriendListLogic(ctx, s.svcCtx)
	return l.UserFriendList(in)
}

func (s *UserServiceServer) RecommendUsers(ctx context.Context, in *user.RecommendUsersRequest) (*user.RecommendUsersResponse, error) {
	l := logic.NewRecommendUsersLogic(ctx, s.svcCtx)
	return l.RecommendUsers(in)
}

// redis自增id
func (s *UserServiceServer) NextUserID(ctx context.Context, in *user.Empty) (*user.NextUserIDResponse, error) {
	l := logic.NewNextUserIDLogic(ctx, s.svcCtx)
	return l.NextUserID(in)
}

func (s *UserServiceServer) AddUserId(ctx context.Context, in *user.Empty) (*user.Empty, error) {
	l := logic.NewAddUserIdLogic(ctx, s.svcCtx)
	return l.AddUserId(in)
}

func (s *UserServiceServer) DecUserID(ctx context.Context, in *user.Empty) (*user.Empty, error) {
	l := logic.NewDecUserIDLogic(ctx, s.svcCtx)
	return l.DecUserID(in)
}

// 群聊
func (s *UserServiceServer) UserCreateGroup(ctx context.Context, in *user.UserCreateGroupRequest) (*user.Empty, error) {
	l := logic.NewUserCreateGroupLogic(ctx, s.svcCtx)
	return l.UserCreateGroup(in)
}

func (s *UserServiceServer) UserSelectGroup(ctx context.Context, in *user.UserSelectGroupsRequest) (*user.UserSelectGroupsResponse, error) {
	l := logic.NewUserSelectGroupLogic(ctx, s.svcCtx)
	return l.UserSelectGroup(in)
}

func (s *UserServiceServer) UserUserSelfGroup(ctx context.Context, in *user.UserSelfGroupRequest) (*user.UserSelectGroupsResponse, error) {
	l := logic.NewUserUserSelfGroupLogic(ctx, s.svcCtx)
	return l.UserUserSelfGroup(in)
}

func (s *UserServiceServer) UserSelectDetailGroup(ctx context.Context, in *user.DetailGroupRequest) (*user.DetailGroupResponse, error) {
	l := logic.NewUserSelectDetailGroupLogic(ctx, s.svcCtx)
	return l.UserSelectDetailGroup(in)
}
