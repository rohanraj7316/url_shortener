package services

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

// CacheConnection cache connetion
var CacheConnection *redis.Client

// RedisConnection establish connection to redis-server
func RedisConnection() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})

	CacheConnection = rdb
}

// RedisConnectionHealthCheck connection ping
func RedisConnectionHealthCheck() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := CacheConnection.Ping(ctx).Result()
	if err != nil {
		return err
	}

	return nil
}
