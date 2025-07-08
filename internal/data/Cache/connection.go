package cache

import (
	"context"
	"go_backend/internal/config"

	"github.com/redis/go-redis/v9"
)

func GetCacheClient(
	cfg *config.CloudConfig,
) (*redis.Client, error) {

	var DbIndex int = 0

	client := redis.NewClient(&redis.Options{
		Addr:     "my-redis-endpoint.cache.amazonaws.com:6379",
		Password: "",
		DB:       DbIndex,
	})

	err := client.Ping(context.Background()).Err()

	return client, err
}
