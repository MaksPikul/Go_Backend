package rds

import (
	"context"
	"database/sql"
	"fmt"
	"go_backend/internal/config"

	"github.com/aws-sdk-go-v2/feature/rds/auth"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// either connects to docker
// or RDB proxy pool
func GetDB(
	RDB *config.DBConfig,
	// AWS_RDB *config.AwsRdbConfig,
) (*gorm.DB, error) {

	authToken, err := buildAuthToken(RDB /*AWS_RDB*/)
	if err != nil {
		return nil, err
	}

	dsn := createConString(authToken, RDB /*AWS_RDB*/)

	sqlDB, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	return gorm.Open(postgres.New(postgres.Config{Conn: sqlDB}), &gorm.Config{})
}

func buildAuthToken(
	RDB *config.DBConfig,
	// AWS_RDB *config.AwsRdbConfig,
) (string, error) {

	return auth.BuildAuthToken(context.TODO(), fmt.Sprintf("%s:%d", endpoint, port), region, user, cfg.Credentials)
}

func createConString(
	authToken string,
	RDB *config.DBConfig,
	// AWS_RDB *config.AwsRdbConfig,
) string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=require",
		RDB.Endpoint, RDB.Port, RDB.Username, authToken, RDB.DbName,
	)
}
