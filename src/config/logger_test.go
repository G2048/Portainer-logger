package config

import (
	"log/slog"
	"testing"
)

func TestLogger(t *testing.T) {
	var logger *slog.Logger
	logger = InitLogger(slog.LevelDebug)
	logger.Debug("Hello World!")

	logger = InitLogger(slog.LevelInfo)
	logger.Info("Hello Json World!")
}
