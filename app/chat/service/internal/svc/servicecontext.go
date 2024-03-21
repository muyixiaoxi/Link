package svc

import (
	initGorm "chat/common/gorm"
	redis2 "chat/common/redis"
	"chat/service/internal/config"
	"chat/service/internal/types"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
	RDB    *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	rc := redis.RedisConf{
		Host: c.RedisConf.Host,
		Type: c.RedisConf.Type,
		Pass: c.RedisConf.Pass,
	}
	db := initGorm.InitGorm(c.Mysql.DataSource)
	db.AutoMigrate(&types.Message{})
	db.AutoMigrate(&types.GroupMessage{})
	red := redis2.InitRedis(rc)
	return &ServiceContext{
		Config: c,
		DB:     db,
		RDB:    red,
	}
}
