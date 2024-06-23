package site

import "log/slog"

type MainHandlers struct {
	log *slog.Logger
}

func NewMainHandlers(log *slog.Logger) *MainHandlers {
	return &MainHandlers{log: log}
}
