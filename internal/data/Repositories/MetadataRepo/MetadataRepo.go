package repositories

import (
	"gorm.io/gorm"
)

type MetadataRepository struct {
	db *gorm.DB
}

func NewMetadataRepository(db *gorm.DB) *MetadataRepository {
	return &MetadataRepository{
		db: db,
	}
}

// Only this needs to be wrapped
func (br *MetadataRepository) GetContent() string {

	// User

	// Get List

	// return list

	return "s"
}
