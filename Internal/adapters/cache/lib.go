package cache

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

var (
	noDataInDatabase = ":0:"
	haveDataInDatabase = ":1:"
)

var (
	ErrCacheMiss = errors.New("cache: miss")
	ErrNotFound = errors.New("cache: not found in database")

	ErrUnableToSetCache = errors.New("cache: unable to set cache")
	ErrUnableToDelCache = errors.New("cache: unable to set cache")
)

func Set(ctx context.Context, key string, data string, cacheTime time.Duration) error {
	data = fmt.Sprint(haveDataInDatabase, data)
	if err := redisClient.Set(ctx, key, data, cacheTime).Err(); err != nil {
		return ErrUnableToSetCache
	}
	return nil
}

func Get(ctx context.Context, key string) (string, error) {
	data, err := redisClient.Get(ctx, key).Result()

	status := data[:3]
	value := data[3:]

	if err == redis.Nil {
		return "", ErrCacheMiss
	}
	if status == noDataInDatabase {
		return "", ErrNotFound
	}
	return value, nil
}

func Del(ctx context.Context, key string) error {
	if err := redisClient.Del(ctx, key).Err(); err != nil {
		return ErrUnableToDelCache
	}
	return nil
}