package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"gorm.io/gorm"
	initGorm "tag/common/gorm"
	redis2 "tag/common/redis"
	"tag/service/internal/config"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
	RDB    *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := initGorm.InitGorm(c.Mysql.DataSource)
	//db.AutoMigrate(&types.Tag{})
	//db.AutoMigrate(&types.UserTagFollow{})
	rc := redis.RedisConf{
		Host: c.RedisConf.Host,
		Type: c.RedisConf.Type,
		Pass: c.RedisConf.Pass,
	}
	red := redis2.InitRedis(rc)
	return &ServiceContext{
		Config: c,
		DB:     db,
		RDB:    red,
	}
}
