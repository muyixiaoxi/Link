package svc

import (
	"Link/restful/tag/internal/config"
	"Link/service/tag/client/taglogin"
	"Link/service/tag/client/tagsign"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config   config.Config
	TagLogin taglogin.TagLogin
	TagTest  tagsign.TagSign
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		TagLogin: taglogin.NewTagLogin(zrpc.MustNewClient(c.TagRpc)),
		TagTest:  tagsign.NewTagSign(zrpc.MustNewClient(c.TagRpc)),
	}
}
