package config

import (
	"go_backend/internal/util"
	"os"
)

type Config struct {
	DB    *DBConfig
	Cache *CacheConfig
	Cloud *CloudConfig
}

type DBConfig struct {
	Endpoint string
	Port     string
	Username string
	DbName   string
}

func loadDBConfig() *DBConfig {

	if util.IsProd() {

		// AWS
		return &DBConfig{
			Endpoint: os.Getenv("AWS_RDS_PROXY_ENDPOINT"),
			Port:     os.Getenv("AWS_RDS_PROXY_PORT"),
			Username: os.Getenv("AWS_RDS_PROXY_USERNAME"),
			DbName:   os.Getenv("AWS_RDS_PROXY_DBNAME"),
		}
	}

	// DOCKER
	return &DBConfig{
		Endpoint: os.Getenv("DOCKER_RDS_ENDPOINT"),
		Port:     os.Getenv("DOCKER_RDS_PORT"),
		Username: os.Getenv("DOCKER_RDS_USERNAME"),
		DbName:   os.Getenv("DOCKER_RDS_DBNAME"),
	}
}

type CacheConfig struct {
}

func loadCacheConfig() *CacheConfig {

}

type CloudConfig struct {
	Provider string
	Region   string
}

func loadAWSConfig() *CloudConfig {
	return &CloudConfig{
		Provider: os.Getenv("CLOUD_PROVIDER"),
		Region:   os.Getenv("CLOUD_REGION"),
	}
}

func LoadConfig() (*Config, error) {

	dbCfg := loadDBConfig()

	cloudCfg := loadAWSConfig()

	cacheCfg := loadCacheConfig()

	return &Config{
		DB:    dbCfg,
		Cache: cacheCfg,
		Cloud: cloudCfg,
	}, nil
}

/*

	// prehaps a handle config errors function ?
	if dbCfg.Port == "" {
		slog.Warn("PORT not set in environment. Using default :8080")
		cfg.Port = "8080"
	}
	if dbCfg.Endpoint == "" {
		return nil, errors.New("DBConfig.Endpoint is required")
	}
*/
