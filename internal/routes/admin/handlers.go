package admin

import "log/slog"


const baseHTMLLayout string = "./ui/html/admin/base.layout.html"

type AdminHandlers struct {
	log *slog.Logger
}

func NewAdminHandlers(log *slog.Logger) *AdminHandlers {
	return &AdminHandlers{log: log}
}