package repositories

import (
	bucket "go_backend/internal/data/Bucket"
	bRepo "go_backend/internal/data/Repositories/BucketRepo"
	mdRepo "go_backend/internal/data/Repositories/MetadataRepo"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Repositories struct {
	mdRepository     *mdRepo.MetadataRepository
	bucketRepository *bRepo.BucketRepository
}

func NewRepositories(
	db *gorm.DB,
	uploader *bucket.IBucketUploader,
) *Repositories {
	return &Repositories{
		mdRepository:     mdRepo.NewMetadataRepository(db),
		bucketRepository: bRepo.NewBucketRepository(uploader),
	}
}

type WrappedRepos struct {
	mdRepository     *mdRepo.WrappedMetadataRepo
	bucketRepository *bRepo.BucketRepository
}

func NewWrappedRepos(
	base *Repositories,
	redisClient *redis.Client,
) *WrappedRepos {
	return &WrappedRepos{
		mdRepository:     mdRepo.NewWrappedMetadataRepo(base.mdRepository, redisClient),
		bucketRepository: base.bucketRepository,
	}
}
