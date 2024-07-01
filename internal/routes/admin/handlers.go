package admin

import (
	"log/slog"
)

const baseHTMLLayout string = "./ui/html/admin/admin_base.layout.html"

type AdminHandlers struct {
	log  *slog.Logger
	menu RepositoryMenu
}

func NewAdminHandlers(log *slog.Logger, menu RepositoryMenu) *AdminHandlers {
	return &AdminHandlers{
		log:  log,
		menu: menu,
	}
}
