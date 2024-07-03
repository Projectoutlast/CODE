package admin

import (
	"log/slog"
)

const baseHTMLLayout string = "./ui/html/admin/admin_base.layout.html"

type AdminHandlers struct {
	category RepositoryCategory
	log      *slog.Logger
	menu     RepositoryMenu
}

func NewAdminHandlers(category RepositoryCategory, log *slog.Logger, menu RepositoryMenu) *AdminHandlers {
	return &AdminHandlers{
		category: category,
		log:      log,
		menu:     menu,
	}
}
