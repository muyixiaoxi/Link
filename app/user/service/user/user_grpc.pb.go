// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.2
// source: app/user/service/user.proto

package user

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	UserService_UserCreate_FullMethodName            = "/user.UserService/UserCreate"
	UserService_UserCreateRevertLogin_FullMethodName = "/user.UserService/UserCreateRevertLogin"
	UserService_UserLogin_FullMethodName             = "/user.UserService/UserLogin"
	UserService_UserInfo_FullMethodName              = "/user.UserService/UserInfo"
	UserService_UserUpdateInfo_FullMethodName        = "/user.UserService/UserUpdateInfo"
	UserService_UserFlowed_FullMethodName            = "/user.UserService/UserFlowed"
	UserService_UserDisposeFlowed_FullMethodName     = "/user.UserService/UserDisposeFlowed"
	UserService_UserUpdateRemark_FullMethodName      = "/user.UserService/UserUpdateRemark"
	UserService_NextUserID_FullMethodName            = "/user.UserService/NextUserID"
	UserService_UserCreateGroup_FullMethodName       = "/user.UserService/UserCreateGroup"
	UserService_UserSelectGroup_FullMethodName       = "/user.UserService/UserSelectGroup"
	UserService_UserSelectDetailGroup_FullMethodName = "/user.UserService/UserSelectDetailGroup"
)

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	UserCreate(ctx context.Context, in *UserCreateRequest, opts ...grpc.CallOption) (*UserCreateResponse, error)
	UserCreateRevertLogin(ctx context.Context, in *UserCreateRequest, opts ...grpc.CallOption) (*UserCreateResponse, error)
	UserLogin(ctx context.Context, in *UserLoginRequest, opts ...grpc.CallOption) (*UserLoginResponse, error)
	UserInfo(ctx context.Context, in *UserInfoRequest, opts ...grpc.CallOption) (*UserInfoResponse, error)
	UserUpdateInfo(ctx context.Context, in *UserUpdateInfoRequest, opts ...grpc.CallOption) (*Empty, error)
	UserFlowed(ctx context.Context, in *UserAddRequest, opts ...grpc.CallOption) (*Empty, error)
	UserDisposeFlowed(ctx context.Context, in *DisposeFlowedRequest, opts ...grpc.CallOption) (*Empty, error)
	UserUpdateRemark(ctx context.Context, in *UserUpdateRemarkRequest, opts ...grpc.CallOption) (*Empty, error)
	NextUserID(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*NextUserIDResponse, error)
	// 群聊
	UserCreateGroup(ctx context.Context, in *UserCreateGroupRequest, opts ...grpc.CallOption) (*Empty, error)
	UserSelectGroup(ctx context.Context, in *UserSelectGroupsRequest, opts ...grpc.CallOption) (*UserSelectGroupsResponse, error)
	UserSelectDetailGroup(ctx context.Context, in *DetailGroupRequest, opts ...grpc.CallOption) (*DetailGroupResponse, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) UserCreate(ctx context.Context, in *UserCreateRequest, opts ...grpc.CallOption) (*UserCreateResponse, error) {
	out := new(UserCreateResponse)
	err := c.cc.Invoke(ctx, UserService_UserCreate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UserCreateRevertLogin(ctx context.Context, in *UserCreateRequest, opts ...grpc.CallOption) (*UserCreateResponse, error) {
	out := new(UserCreateResponse)
	err := c.cc.Invoke(ctx, UserService_UserCreateRevertLogin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UserLogin(ctx context.Context, in *UserLoginRequest, opts ...grpc.CallOption) (*UserLoginResponse, error) {
	out := new(UserLoginResponse)
	err := c.cc.Invoke(ctx, UserService_UserLogin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UserInfo(ctx context.Context, in *UserInfoRequest, opts ...grpc.CallOption) (*UserInfoResponse, error) {
	out := new(UserInfoResponse)
	err := c.cc.Invoke(ctx, UserService_UserInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UserUpdateInfo(ctx context.Context, in *UserUpdateInfoRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, UserService_UserUpdateInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UserFlowed(ctx context.Context, in *UserAddRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, UserService_UserFlowed_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UserDisposeFlowed(ctx context.Context, in *DisposeFlowedRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, UserService_UserDisposeFlowed_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UserUpdateRemark(ctx context.Context, in *UserUpdateRemarkRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, UserService_UserUpdateRemark_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) NextUserID(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*NextUserIDResponse, error) {
	out := new(NextUserIDResponse)
	err := c.cc.Invoke(ctx, UserService_NextUserID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UserCreateGroup(ctx context.Context, in *UserCreateGroupRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, UserService_UserCreateGroup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UserSelectGroup(ctx context.Context, in *UserSelectGroupsRequest, opts ...grpc.CallOption) (*UserSelectGroupsResponse, error) {
	out := new(UserSelectGroupsResponse)
	err := c.cc.Invoke(ctx, UserService_UserSelectGroup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UserSelectDetailGroup(ctx context.Context, in *DetailGroupRequest, opts ...grpc.CallOption) (*DetailGroupResponse, error) {
	out := new(DetailGroupResponse)
	err := c.cc.Invoke(ctx, UserService_UserSelectDetailGroup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	UserCreate(context.Context, *UserCreateRequest) (*UserCreateResponse, error)
	UserCreateRevertLogin(context.Context, *UserCreateRequest) (*UserCreateResponse, error)
	UserLogin(context.Context, *UserLoginRequest) (*UserLoginResponse, error)
	UserInfo(context.Context, *UserInfoRequest) (*UserInfoResponse, error)
	UserUpdateInfo(context.Context, *UserUpdateInfoRequest) (*Empty, error)
	UserFlowed(context.Context, *UserAddRequest) (*Empty, error)
	UserDisposeFlowed(context.Context, *DisposeFlowedRequest) (*Empty, error)
	UserUpdateRemark(context.Context, *UserUpdateRemarkRequest) (*Empty, error)
	NextUserID(context.Context, *Empty) (*NextUserIDResponse, error)
	// 群聊
	UserCreateGroup(context.Context, *UserCreateGroupRequest) (*Empty, error)
	UserSelectGroup(context.Context, *UserSelectGroupsRequest) (*UserSelectGroupsResponse, error)
	UserSelectDetailGroup(context.Context, *DetailGroupRequest) (*DetailGroupResponse, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) UserCreate(context.Context, *UserCreateRequest) (*UserCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserCreate not implemented")
}
func (UnimplementedUserServiceServer) UserCreateRevertLogin(context.Context, *UserCreateRequest) (*UserCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserCreateRevertLogin not implemented")
}
func (UnimplementedUserServiceServer) UserLogin(context.Context, *UserLoginRequest) (*UserLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserLogin not implemented")
}
func (UnimplementedUserServiceServer) UserInfo(context.Context, *UserInfoRequest) (*UserInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserInfo not implemented")
}
func (UnimplementedUserServiceServer) UserUpdateInfo(context.Context, *UserUpdateInfoRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserUpdateInfo not implemented")
}
func (UnimplementedUserServiceServer) UserFlowed(context.Context, *UserAddRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserFlowed not implemented")
}
func (UnimplementedUserServiceServer) UserDisposeFlowed(context.Context, *DisposeFlowedRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserDisposeFlowed not implemented")
}
func (UnimplementedUserServiceServer) UserUpdateRemark(context.Context, *UserUpdateRemarkRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserUpdateRemark not implemented")
}
func (UnimplementedUserServiceServer) NextUserID(context.Context, *Empty) (*NextUserIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NextUserID not implemented")
}
func (UnimplementedUserServiceServer) UserCreateGroup(context.Context, *UserCreateGroupRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserCreateGroup not implemented")
}
func (UnimplementedUserServiceServer) UserSelectGroup(context.Context, *UserSelectGroupsRequest) (*UserSelectGroupsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserSelectGroup not implemented")
}
func (UnimplementedUserServiceServer) UserSelectDetailGroup(context.Context, *DetailGroupRequest) (*DetailGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserSelectDetailGroup not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_UserCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UserCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_UserCreate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UserCreate(ctx, req.(*UserCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UserCreateRevertLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UserCreateRevertLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_UserCreateRevertLogin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UserCreateRevertLogin(ctx, req.(*UserCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UserLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UserLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_UserLogin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UserLogin(ctx, req.(*UserLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_UserInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UserInfo(ctx, req.(*UserInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UserUpdateInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserUpdateInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UserUpdateInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_UserUpdateInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UserUpdateInfo(ctx, req.(*UserUpdateInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UserFlowed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserAddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UserFlowed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_UserFlowed_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UserFlowed(ctx, req.(*UserAddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UserDisposeFlowed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DisposeFlowedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UserDisposeFlowed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_UserDisposeFlowed_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UserDisposeFlowed(ctx, req.(*DisposeFlowedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UserUpdateRemark_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserUpdateRemarkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UserUpdateRemark(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_UserUpdateRemark_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UserUpdateRemark(ctx, req.(*UserUpdateRemarkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_NextUserID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).NextUserID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_NextUserID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).NextUserID(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UserCreateGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserCreateGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UserCreateGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_UserCreateGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UserCreateGroup(ctx, req.(*UserCreateGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UserSelectGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserSelectGroupsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UserSelectGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_UserSelectGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UserSelectGroup(ctx, req.(*UserSelectGroupsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UserSelectDetailGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DetailGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UserSelectDetailGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_UserSelectDetailGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UserSelectDetailGroup(ctx, req.(*DetailGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UserCreate",
			Handler:    _UserService_UserCreate_Handler,
		},
		{
			MethodName: "UserCreateRevertLogin",
			Handler:    _UserService_UserCreateRevertLogin_Handler,
		},
		{
			MethodName: "UserLogin",
			Handler:    _UserService_UserLogin_Handler,
		},
		{
			MethodName: "UserInfo",
			Handler:    _UserService_UserInfo_Handler,
		},
		{
			MethodName: "UserUpdateInfo",
			Handler:    _UserService_UserUpdateInfo_Handler,
		},
		{
			MethodName: "UserFlowed",
			Handler:    _UserService_UserFlowed_Handler,
		},
		{
			MethodName: "UserDisposeFlowed",
			Handler:    _UserService_UserDisposeFlowed_Handler,
		},
		{
			MethodName: "UserUpdateRemark",
			Handler:    _UserService_UserUpdateRemark_Handler,
		},
		{
			MethodName: "NextUserID",
			Handler:    _UserService_NextUserID_Handler,
		},
		{
			MethodName: "UserCreateGroup",
			Handler:    _UserService_UserCreateGroup_Handler,
		},
		{
			MethodName: "UserSelectGroup",
			Handler:    _UserService_UserSelectGroup_Handler,
		},
		{
			MethodName: "UserSelectDetailGroup",
			Handler:    _UserService_UserSelectDetailGroup_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "app/user/service/user.proto",
}
