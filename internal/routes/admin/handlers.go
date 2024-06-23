package admin

import "log/slog"

type AdminHandlers struct {
	log *slog.Logger
}

func NewAdminHandlers(log *slog.Logger) *AdminHandlers {
	return &AdminHandlers{log: log}
}