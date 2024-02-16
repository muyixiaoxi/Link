package svc

import (
	"Link/restful/user/internal/config"
	"Link/service/tag/client/tagsign"
	"Link/service/user/userservice"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc userservice.UserService
	TagRpc  tagsign.TagSign
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: userservice.NewUserService(zrpc.MustNewClient(c.UserRpc)),
		TagRpc:  tagsign.NewTagSign(zrpc.MustNewClient(c.TagRpc)),
	}
}
