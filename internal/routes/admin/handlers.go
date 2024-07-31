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
	employee RepositoryEmployee
}

func NewAdminHandlers(
	category RepositoryCategory,
	dish RepositoryDish,
	log *slog.Logger,
	menu RepositoryMenu,
	employee RepositoryEmployee,
) *AdminHandlers {
	return &AdminHandlers{
		category: category,
		dish:     dish,
		log:      log,
		menu:     menu,
		employee: employee,
	}
}
