package config

import (
	"log/slog"
	"os"
)

func InitLogger(level slog.Level) *slog.Logger {
	var handler slog.Handler
	opts := &slog.HandlerOptions{
		AddSource: false,
		Level:     level,
	}

	if level == slog.LevelDebug {
		opts.AddSource = true
		handler = slog.NewTextHandler(os.Stderr, opts)
	} else {
		handler = slog.NewJSONHandler(os.Stderr, opts)
	}

	logger := slog.New(handler)
	slog.SetDefault(logger)
	return logger
}
