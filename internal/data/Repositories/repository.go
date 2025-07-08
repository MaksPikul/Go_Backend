package repositories

import (
	bucket "go_backend/internal/data/Bucket"
)

type Repositories struct {
	bucketRepository *BucketRepository
}

func NewRepositories(
	uploader *bucket.IBucketUploader,
) *Repositories {
	return &Repositories{
		bucketRepository: NewBucketRepository(uploader),
	}
}
