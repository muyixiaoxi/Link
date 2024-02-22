package svc

import (
	"github.com/zeromicro/go-zero/zrpc"

	"user/restful/internal/config"
	"user/service/tag/service/client/taglogin"
	"user/service/tag/service/client/tagsign"
	"user/service/userservice"
)

type ServiceContext struct {
	Config      config.Config
	UserRpc     userservice.UserService
	TagSignRpc  tagsign.TagSign
	TagLoginRpc taglogin.TagLogin
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		UserRpc:     userservice.NewUserService(zrpc.MustNewClient(c.UserRpc)),
		TagSignRpc:  tagsign.NewTagSign(zrpc.MustNewClient(c.TagRpc)),
		TagLoginRpc: taglogin.NewTagLogin(zrpc.MustNewClient(c.TagRpc)),
	}
}
