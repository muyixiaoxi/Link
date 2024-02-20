package initRedis

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
)

func InitRedis(conf redis.RedisConf) *redis.Redis {
	rds := redis.MustNewRedis(conf)
	return rds
}
