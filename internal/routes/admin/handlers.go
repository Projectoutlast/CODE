package admin

import (
	"log/slog"
)

const baseHTMLLayout string = "./ui/html/admin/admin_base.layout.html"

type AdminHandlers struct {
	category RepositoryCategory
	dish     RepositoryDish
	log      *slog.Logger
	menu     RepositoryMenu
}

func NewAdminHandlers(
	category RepositoryCategory,
	dish RepositoryDish,
	log *slog.Logger,
	menu RepositoryMenu,
) *AdminHandlers {
	return &AdminHandlers{
		category: category,
		dish:     dish,
		log:      log,
		menu:     menu,
	}
}
