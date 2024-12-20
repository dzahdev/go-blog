package logger

import (
	"log/slog"
	"os"
)

const (
	envDevelopment = "development"
	envLocal       = "local"
	envProduction  = "production"
)

func SetupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envProduction:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	default:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	}

	return log
}
