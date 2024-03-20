package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"user/restful/internal/config"
	"user/restful/internal/middleware"
	"user/service/tag/service/client/taglogin"
	"user/service/tag/service/client/tagsign"
	"user/service/userservice"
)

type ServiceContext struct {
	Config      config.Config
	JWT         rest.Middleware
	UserRpc     userservice.UserService
	TagSignRpc  tagsign.TagSign
	TagLoginRpc taglogin.TagLogin
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		JWT:         middleware.NewJWTMiddleware().Handle,
		UserRpc:     userservice.NewUserService(zrpc.MustNewClient(c.UserRpc)),
		TagSignRpc:  tagsign.NewTagSign(zrpc.MustNewClient(c.TagRpc)),
		TagLoginRpc: taglogin.NewTagLogin(zrpc.MustNewClient(c.TagRpc)),
	}
}
