package svc

import (
	initGorm "Link/internal/gorm"
	"Link/service/chat/internal/config"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := initGorm.InitGorm(c.Mysql.DataSource)
	return &ServiceContext{
		Config: c,
		DB:     db,
	}
}
