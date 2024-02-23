package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"gorm.io/gorm"
	initGorm "user/common/gorm"
	redis2 "user/common/redis"
	"user/service/internal/config"
	"user/service/internal/types"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
	RDB    *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := initGorm.InitGorm(c.Mysql.DataSource)
	db.AutoMigrate(&types.UserGroupChat{})
	db.AutoMigrate(&types.User{})
	db.AutoMigrate(&types.GroupChat{})
	db.AutoMigrate(&types.ApplyFor{})
	db.AutoMigrate(&types.Friend{})

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
