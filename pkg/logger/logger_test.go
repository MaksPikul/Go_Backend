package logger

import (
	"bytes"
	"log/slog"
	"os"
	"strings"
	"testing"
)

func TestMultiHandler(t *testing.T) {
	var buf1, buf2 bytes.Buffer

	textHandler := slog.NewTextHandler(&buf1, nil)
	jsonHandler := slog.NewJSONHandler(&buf2, nil)

	handler := &MultiHandler{
		handlers: []slog.Handler{textHandler, jsonHandler},
	}
	logger := slog.New(handler)

	logger.Info("Test message", slog.String("key", "value"))

	if !strings.Contains(buf1.String(), "Test message") {
		t.Error("TextHandler did not log correctly")
	}
	if !strings.Contains(buf2.String(), `"msg":"Test message"`) {
		t.Error("JSONHandler did not log correctly")
	}
}

func TestLoggerFileOutput(t *testing.T) {
	logFile := "test.log"
	defer os.Remove(logFile)

	f, err := os.Create(logFile)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	fileHandler := slog.NewJSONHandler(f, nil)
	logger := slog.New(fileHandler)

	logger.Error("Test error", slog.String("env", "test"))

	f.Sync()
	content, err := os.ReadFile(logFile)
	if err != nil {
		t.Fatal(err)
	}

	if !strings.Contains(string(content), "Test error") {
		t.Error("Log file does not contain expected message")
	}
}

// Cloud Watch Test
