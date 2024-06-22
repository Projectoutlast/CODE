package site

import "log/slog"

type Handlers struct {
	log *slog.Logger
}

func NewHandlers(log *slog.Logger) *Handlers {
	return &Handlers{log: log}
}
