package redis2

import "github.com/go-redis/redis/v8"

func InitRedis(options *redis.Options) *redis.Client {
	return redis.NewClient(options)
}
