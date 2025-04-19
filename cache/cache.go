package cache

import (
	"github.com/redis/go-redis/v9"
	"memoryDataBase/config"
)

var RedisClient *redis.Client

// InitRedis 初始化redis
func InitRedis(config config.RedisConfig) *redis.Client {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     config.Addr,
		Password: config.Password,
		DB:       config.DB,
	})
	return RedisClient
}
