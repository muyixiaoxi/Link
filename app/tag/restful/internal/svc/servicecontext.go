package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"tag/restful/internal/config"

	"tag/service/client/tagsign"

	"tag/service/client/taglogin"
)

type ServiceContext struct {
	Config    config.Config
	TagLogin  taglogin.TagLogin
	SystemTag tagsign.TagSign
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		TagLogin:  taglogin.NewTagLogin(zrpc.MustNewClient(c.TagRpc)),
		SystemTag: tagsign.NewTagSign(zrpc.MustNewClient(c.TagRpc)),
	}
}
