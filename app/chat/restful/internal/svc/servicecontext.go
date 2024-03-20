package svc

import (
	"chat/restful/internal/config"
	"chat/service/chat"
	"chat/service/user/userservice"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc userservice.UserService
	ChatRpc chat.ChatServer
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: userservice.NewUserService(zrpc.MustNewClient(c.UserRpc)),
	}
}
