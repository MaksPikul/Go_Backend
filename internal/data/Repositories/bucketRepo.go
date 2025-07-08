package repositories

import (
	bucket "go_backend/internal/data/Bucket"
)

type BucketInterface interface {
	// UploadContent (*models.User) error
	// PreSignUrl_UploadContent
	// DeleteContent ()
	// GetContents ()
}

type BucketRepository struct {
	uploader *bucket.IBucketUploader
}

func NewBucketRepository(uploader *bucket.IBucketUploader) *BucketRepository {
	return &BucketRepository{
		uploader: uploader,
	}
}
