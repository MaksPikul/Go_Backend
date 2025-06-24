package config

import (
	"go_backend/internal/util"
	"os"
)

type Config struct {
	RDB *RDBConfig
	//AWS *AWSConfig
}

type RDBConfig struct {
	Endpoint string
	Port     string
	Username string
	DbName   string
}

func loadDBConfig() *RDBConfig {

	if util.IsProd() {

		// AWS
		return &RDBConfig{
			Endpoint: os.Getenv("AWS_RDS_PROXY_ENDPOINT"),
			Port:     os.Getenv("AWS_RDS_PROXY_PORT"),
			Username: os.Getenv("AWS_RDS_PROXY_USERNAME"),
			DbName:   os.Getenv("AWS_RDS_PROXY_DBNAME"),
		}
	}

	// DOCKER
	return &RDBConfig{
		Endpoint: os.Getenv("DOCKER_RDS_ENDPOINT"),
		Port:     os.Getenv("DOCKER_RDS_PORT"),
		Username: os.Getenv("DOCKER_RDS_USERNAME"),
		DbName:   os.Getenv("DOCKER_RDS_DBNAME"),
	}
}

type AWSConfig struct {
}

func LoadConfig() (*Config, error) {

	dbCfg := loadDBConfig()

	return &Config{
		RDB: dbCfg,
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
