// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: chat.proto

package chat

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
	ChatService_SaveOfflineMessage_FullMethodName = "/chat.ChatService/SaveOfflineMessage"
	ChatService_GetOfflineMessage_FullMethodName  = "/chat.ChatService/GetOfflineMessage"
	ChatService_Online_FullMethodName             = "/chat.ChatService/Online"
	ChatService_Offline_FullMethodName            = "/chat.ChatService/Offline"
	ChatService_GetConnectorId_FullMethodName     = "/chat.ChatService/GetConnectorId"
)

// ChatServiceClient is the client API for ChatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChatServiceClient interface {
	SaveOfflineMessage(ctx context.Context, in *SaveMessageRequest, opts ...grpc.CallOption) (*Empty, error)
	GetOfflineMessage(ctx context.Context, in *GetOfflineMessageRequest, opts ...grpc.CallOption) (*Messages, error)
	Online(ctx context.Context, in *OnlineRequest, opts ...grpc.CallOption) (*Empty, error)
	Offline(ctx context.Context, in *UserId, opts ...grpc.CallOption) (*Empty, error)
	GetConnectorId(ctx context.Context, in *UserId, opts ...grpc.CallOption) (*ConnectorId, error)
}

type chatServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChatServiceClient(cc grpc.ClientConnInterface) ChatServiceClient {
	return &chatServiceClient{cc}
}

func (c *chatServiceClient) SaveOfflineMessage(ctx context.Context, in *SaveMessageRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, ChatService_SaveOfflineMessage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) GetOfflineMessage(ctx context.Context, in *GetOfflineMessageRequest, opts ...grpc.CallOption) (*Messages, error) {
	out := new(Messages)
	err := c.cc.Invoke(ctx, ChatService_GetOfflineMessage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) Online(ctx context.Context, in *OnlineRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, ChatService_Online_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) Offline(ctx context.Context, in *UserId, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, ChatService_Offline_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) GetConnectorId(ctx context.Context, in *UserId, opts ...grpc.CallOption) (*ConnectorId, error) {
	out := new(ConnectorId)
	err := c.cc.Invoke(ctx, ChatService_GetConnectorId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatServiceServer is the server API for ChatService service.
// All implementations must embed UnimplementedChatServiceServer
// for forward compatibility
type ChatServiceServer interface {
	SaveOfflineMessage(context.Context, *SaveMessageRequest) (*Empty, error)
	GetOfflineMessage(context.Context, *GetOfflineMessageRequest) (*Messages, error)
	Online(context.Context, *OnlineRequest) (*Empty, error)
	Offline(context.Context, *UserId) (*Empty, error)
	GetConnectorId(context.Context, *UserId) (*ConnectorId, error)
	mustEmbedUnimplementedChatServiceServer()
}

// UnimplementedChatServiceServer must be embedded to have forward compatible implementations.
type UnimplementedChatServiceServer struct {
}

func (UnimplementedChatServiceServer) SaveOfflineMessage(context.Context, *SaveMessageRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveOfflineMessage not implemented")
}
func (UnimplementedChatServiceServer) GetOfflineMessage(context.Context, *GetOfflineMessageRequest) (*Messages, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOfflineMessage not implemented")
}
func (UnimplementedChatServiceServer) Online(context.Context, *OnlineRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Online not implemented")
}
func (UnimplementedChatServiceServer) Offline(context.Context, *UserId) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Offline not implemented")
}
func (UnimplementedChatServiceServer) GetConnectorId(context.Context, *UserId) (*ConnectorId, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConnectorId not implemented")
}
func (UnimplementedChatServiceServer) mustEmbedUnimplementedChatServiceServer() {}

// UnsafeChatServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChatServiceServer will
// result in compilation errors.
type UnsafeChatServiceServer interface {
	mustEmbedUnimplementedChatServiceServer()
}

func RegisterChatServiceServer(s grpc.ServiceRegistrar, srv ChatServiceServer) {
	s.RegisterService(&ChatService_ServiceDesc, srv)
}

func _ChatService_SaveOfflineMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).SaveOfflineMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_SaveOfflineMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).SaveOfflineMessage(ctx, req.(*SaveMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_GetOfflineMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOfflineMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).GetOfflineMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_GetOfflineMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).GetOfflineMessage(ctx, req.(*GetOfflineMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_Online_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OnlineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).Online(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_Online_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).Online(ctx, req.(*OnlineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_Offline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).Offline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_Offline_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).Offline(ctx, req.(*UserId))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_GetConnectorId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).GetConnectorId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_GetConnectorId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).GetConnectorId(ctx, req.(*UserId))
	}
	return interceptor(ctx, in, info, handler)
}

// ChatService_ServiceDesc is the grpc.ServiceDesc for ChatService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChatService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chat.ChatService",
	HandlerType: (*ChatServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SaveOfflineMessage",
			Handler:    _ChatService_SaveOfflineMessage_Handler,
		},
		{
			MethodName: "GetOfflineMessage",
			Handler:    _ChatService_GetOfflineMessage_Handler,
		},
		{
			MethodName: "Online",
			Handler:    _ChatService_Online_Handler,
		},
		{
			MethodName: "Offline",
			Handler:    _ChatService_Offline_Handler,
		},
		{
			MethodName: "GetConnectorId",
			Handler:    _ChatService_GetConnectorId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chat.proto",
}
