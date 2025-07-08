package repositories

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type WrappedMetadataRepo struct {
	repo  *MetadataRepository
	cache *redis.Client
}

func NewWrappedMetadataRepo(r *MetadataRepository, c *redis.Client) *WrappedMetadataRepo {
	return &WrappedMetadataRepo{
		repo:  r,
		cache: c,
	}
}

// This will only be used by SQL
func (d *WrappedMetadataRepo) GetContent() string {

	cacheClient := d.cache

	content, err := cacheClient.Get(context.Background(), "user:1").Result()

	if content == string(redis.Nil) {

		// hit db, update cache

		content = d.repo.GetContent()

		// turn content into a string

		cacheClient.Set(context.Background(), "user:1", content, 0).Err()

	}

	return content
}
