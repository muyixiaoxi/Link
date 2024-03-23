package svc

import (
	"chat/restfulB/internal/config"
	"chat/service/chatservice"

	"chat/service/user/userservice"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc userservice.UserService
	ChatRpc chatservice.ChatService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: userservice.NewUserService(zrpc.MustNewClient(c.UserRpc)),
		ChatRpc: chatservice.NewChatService(zrpc.MustNewClient(c.ChatRpc)),
	}
}
