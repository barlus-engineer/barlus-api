package cache

import (
	"context"

	"github.com/barlus-engineer/barlus-api/config"
	"github.com/redis/go-redis/v9"
)

var Redis *redis.Client

func RedisConnect() error {
	var (
		cfg = config.GetConfig()
		ctx = context.Background()
	)
	opts, err := redis.ParseURL(cfg.Cache.RedisURL)
	if err != nil {
		return err
	}
	if err := redis.NewClient(opts).Ping(ctx).Err(); err != nil {
		return err
	}
	Redis = redis.NewClient(opts)

	return nil
}

func GetRedis() *redis.Client {
	return Redis
}