package logger

import (
	"log/slog"
	"os"
)

func NewLogger(logLevel string) *slog.Logger {
	var log *slog.Logger

	switch logLevel {
	case "dev":
		log = slog.New(slog.NewTextHandler(os.Stdin, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case "test":
		log = slog.New(slog.NewTextHandler(os.Stdin, &slog.HandlerOptions{Level: slog.LevelInfo}))
	case "prod":
		log = slog.New(slog.NewJSONHandler(os.Stdin, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return log
}
