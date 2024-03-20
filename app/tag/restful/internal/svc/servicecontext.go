package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"tag/restful/internal/config"
	"tag/restful/internal/middleware"

	"tag/service/client/tagsign"

	"tag/service/client/taglogin"
)

type ServiceContext struct {
	Config    config.Config
	JWT       rest.Middleware
	TagLogin  taglogin.TagLogin
	SystemTag tagsign.TagSign
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		TagLogin:  taglogin.NewTagLogin(zrpc.MustNewClient(c.TagRpc)),
		JWT:       middleware.NewJWTMiddleware().Handle,
		SystemTag: tagsign.NewTagSign(zrpc.MustNewClient(c.TagRpc)),
	}
}
