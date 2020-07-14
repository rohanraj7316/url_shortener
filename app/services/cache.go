package services

import (
	"context"
	"fmt"
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

// SetValue set cache
func SetValue(ctx context.Context, key string, value interface{}, exp time.Duration) error {
	err := CacheConnection.Set(ctx, key, value, exp).Err()
	if err != nil {
		return err
	}
	return nil
}

// GetValue return value for the key
func GetValue(ctx context.Context, key string) (string, error) {
	val, err := CacheConnection.Get(ctx, key).Result()
	if err == redis.Nil {
		// TODO: figure out logging
		return "", fmt.Errorf("Redis: no value")
	} else if err != nil {
		return "", err
	} else {
		return val, nil
	}
}
