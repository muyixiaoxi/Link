package svc

import (
	"Link/restful/tag/internal/config"
	"Link/service/tag/client/taglogin"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config   config.Config
	TagLogin taglogin.TagLogin
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		TagLogin: taglogin.NewTagLogin(zrpc.MustNewClient(c.TagRpc)),
	}
}
