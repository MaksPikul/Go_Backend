package bucket

import (
	"fmt"
	"go_backend/internal/config"
	bucket "go_backend/internal/data/Bucket"
	"log/slog"
	"sync"
)

// Singleton Pattern

type IBucketUploader interface {
}

var (
	once             sync.Once
	uploaderInstance *IBucketUploader
)

func GetBucketUploader(
	cfg *config.CloudConfig, // Rename to Cloud Config
) (*IBucketUploader, error) {

	var err error

	if uploaderInstance == nil {

		once.Do(func() {
			switch cfg.Provider {
			case "aws":
				uploaderInstance, err = bucket.NewAWSUploader(cfg)

			case "gcp":
				// gcp when time come

			case "azure":
				// azure

			default:
				err = fmt.Errorf("unsupported provider: %s", cfg.Provider)
			}
		})
	} else {
		slog.Warn("Attempt was made to Create Another Singleton Uploader Instance.")
	}

	return uploaderInstance, err
}
