package svc

import (
	initGorm "Link/internal/gorm"
	redis2 "Link/internal/redis"
	"Link/service/user/internal/config"
	"Link/service/user/internal/types"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
	RDB    *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := initGorm.InitGorm(c.Mysql.DataSource)
	db.AutoMigrate(&types.User{}, &types.GroupChat{}, &types.Flowed{})
	db.AutoMigrate(&types.ApplyFor{})
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
