// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package userservice

import (
	"context"

	"user/service/user"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	DetailGroupRequest       = user.DetailGroupRequest
	DetailGroupResponse      = user.DetailGroupResponse
	Empty                    = user.Empty
	GroupInformation         = user.GroupInformation
	UserAddRequest           = user.UserAddRequest
	UserCreateGroupRequest   = user.UserCreateGroupRequest
	UserCreateRequest        = user.UserCreateRequest
	UserCreateResponse       = user.UserCreateResponse
	UserInfoRequest          = user.UserInfoRequest
	UserInfoResponse         = user.UserInfoResponse
	UserLoginRequest         = user.UserLoginRequest
	UserLoginResponse        = user.UserLoginResponse
	UserSelectGroupsRequest  = user.UserSelectGroupsRequest
	UserSelectGroupsResponse = user.UserSelectGroupsResponse
	UserUpdateInfoRequest    = user.UserUpdateInfoRequest
	UserUpdateRemarkRequest  = user.UserUpdateRemarkRequest

	UserService interface {
		UserCreate(ctx context.Context, in *UserCreateRequest, opts ...grpc.CallOption) (*UserCreateResponse, error)
		UserCreateRevertLogin(ctx context.Context, in *UserCreateRequest, opts ...grpc.CallOption) (*UserCreateResponse, error)
		UserLogin(ctx context.Context, in *UserLoginRequest, opts ...grpc.CallOption) (*UserLoginResponse, error)
		UserInfo(ctx context.Context, in *UserInfoRequest, opts ...grpc.CallOption) (*UserInfoResponse, error)
		UserUpdateInfo(ctx context.Context, in *UserUpdateInfoRequest, opts ...grpc.CallOption) (*Empty, error)
		UserFlowed(ctx context.Context, in *UserAddRequest, opts ...grpc.CallOption) (*Empty, error)
		UserUpdateRemark(ctx context.Context, in *UserUpdateRemarkRequest, opts ...grpc.CallOption) (*Empty, error)
		// 群聊
		UserCreateGroup(ctx context.Context, in *UserCreateGroupRequest, opts ...grpc.CallOption) (*Empty, error)
		UserSelectGroup(ctx context.Context, in *UserSelectGroupsRequest, opts ...grpc.CallOption) (*UserSelectGroupsResponse, error)
		UserSelectDetailGroup(ctx context.Context, in *DetailGroupRequest, opts ...grpc.CallOption) (*DetailGroupResponse, error)
	}

	defaultUserService struct {
		cli zrpc.Client
	}
)

func NewUserService(cli zrpc.Client) UserService {
	return &defaultUserService{
		cli: cli,
	}
}

func (m *defaultUserService) UserCreate(ctx context.Context, in *UserCreateRequest, opts ...grpc.CallOption) (*UserCreateResponse, error) {
	client := user.NewUserServiceClient(m.cli.Conn())
	return client.UserCreate(ctx, in, opts...)
}

func (m *defaultUserService) UserCreateRevertLogin(ctx context.Context, in *UserCreateRequest, opts ...grpc.CallOption) (*UserCreateResponse, error) {
	client := user.NewUserServiceClient(m.cli.Conn())
	return client.UserCreateRevertLogin(ctx, in, opts...)
}

func (m *defaultUserService) UserLogin(ctx context.Context, in *UserLoginRequest, opts ...grpc.CallOption) (*UserLoginResponse, error) {
	client := user.NewUserServiceClient(m.cli.Conn())
	return client.UserLogin(ctx, in, opts...)
}

func (m *defaultUserService) UserInfo(ctx context.Context, in *UserInfoRequest, opts ...grpc.CallOption) (*UserInfoResponse, error) {
	client := user.NewUserServiceClient(m.cli.Conn())
	return client.UserInfo(ctx, in, opts...)
}

func (m *defaultUserService) UserUpdateInfo(ctx context.Context, in *UserUpdateInfoRequest, opts ...grpc.CallOption) (*Empty, error) {
	client := user.NewUserServiceClient(m.cli.Conn())
	return client.UserUpdateInfo(ctx, in, opts...)
}

func (m *defaultUserService) UserFlowed(ctx context.Context, in *UserAddRequest, opts ...grpc.CallOption) (*Empty, error) {
	client := user.NewUserServiceClient(m.cli.Conn())
	return client.UserFlowed(ctx, in, opts...)
}

func (m *defaultUserService) UserUpdateRemark(ctx context.Context, in *UserUpdateRemarkRequest, opts ...grpc.CallOption) (*Empty, error) {
	client := user.NewUserServiceClient(m.cli.Conn())
	return client.UserUpdateRemark(ctx, in, opts...)
}

// 群聊
func (m *defaultUserService) UserCreateGroup(ctx context.Context, in *UserCreateGroupRequest, opts ...grpc.CallOption) (*Empty, error) {
	client := user.NewUserServiceClient(m.cli.Conn())
	return client.UserCreateGroup(ctx, in, opts...)
}

func (m *defaultUserService) UserSelectGroup(ctx context.Context, in *UserSelectGroupsRequest, opts ...grpc.CallOption) (*UserSelectGroupsResponse, error) {
	client := user.NewUserServiceClient(m.cli.Conn())
	return client.UserSelectGroup(ctx, in, opts...)
}

func (m *defaultUserService) UserSelectDetailGroup(ctx context.Context, in *DetailGroupRequest, opts ...grpc.CallOption) (*DetailGroupResponse, error) {
	client := user.NewUserServiceClient(m.cli.Conn())
	return client.UserSelectDetailGroup(ctx, in, opts...)
}
