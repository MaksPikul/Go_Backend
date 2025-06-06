package logger

import (
	"log/slog"
	"os"

	"go_backend/pkg/logger/logHandlers"
)

// uses composite pattern i believe

func InitLogger() *os.File {

	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic("fatal error: failed to init logger:" + err.Error())
	}

	providers := getHandlerProviders()

	config := getHandlerConfig()

	handlers := getHandlers(providers, config)

	slog.SetDefault(slog.New(&MultiHandler{handlers: handlers}))

	return file
}

// To add another log handler,
// implement a HandlerProvider, and then add to list
func getHandlerProviders() []logHandlers.HandlerProvider {
	return []logHandlers.HandlerProvider{
		logHandlers.ConsoleProvider(),
		logHandlers.FileProvider(),
		logHandlers.CloudWatchProvider(),
	}
}

func getHandlerConfig() *logHandlers.HandlerConfig {
	return &logHandlers.HandlerConfig{
		FilePath: "app.log",
	}
}

func getHandlers(
	providers []logHandlers.HandlerProvider,
	config *logHandlers.HandlerConfig,
) []slog.Handler {

	var handlers []slog.Handler
	for _, provider := range providers {
		h, err := provider(config)
		if err != nil {
			panic("fatal error: failed to init logger:" + err.Error())
		}
		if h != nil {
			handlers = append(handlers, h)
		}
	}

	return handlers
}
