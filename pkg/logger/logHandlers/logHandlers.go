package logHandlers

import (
	"fmt"
	"go_backend/internal/util"
	"log/slog"
	"os"

	slogcloud "github.com/melkeydev/slog-cloud"
)

type HandlerProvider func(config *HandlerConfig) (slog.Handler, error)

type HandlerOption func(*HandlerConfig)

type HandlerConfig struct {
	FilePath string
}

func ConsoleProvider() HandlerProvider {
	return func(config *HandlerConfig) (slog.Handler, error) {
		return slog.NewTextHandler(os.Stdout, nil), nil
	}
}

func FileProvider() HandlerProvider {
	return func(config *HandlerConfig) (slog.Handler, error) {
		path := "app.log"
		if config.FilePath != "" {
			path = config.FilePath
		}

		file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			return nil, fmt.Errorf("failed to open log file: %w", err)
		}

		return slog.NewJSONHandler(file, nil), nil
	}
}

func CloudWatchProvider() HandlerProvider {
	return func(config *HandlerConfig) (slog.Handler, error) {
		if !util.IsProd() {
			return nil, nil
		}

		client, err := slogcloud.NewCloudwatchClient(
			os.Getenv("CLOUDWATCH_ACCESS_KEY"),
			os.Getenv("CLOUDWATCH_SECRET_KEY"),
			"prod",
			os.Getenv("CLOUDWATCH_REGION"),
		)
		if err != nil {
			return nil, err
		}

		return slogcloud.NewCloudWatchLogHandler(client), nil
	}
}
