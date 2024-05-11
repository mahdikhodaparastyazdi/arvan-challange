package redis

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func NewClient(ctx context.Context, cfg Config) (*redis.Client, error) {
	if cfg.Port == "" {
		cfg.Port = "6379"
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       cfg.Database,
	})
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}

	return rdb, err
}
