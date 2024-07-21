package admin

import (
	"code/internal/repository/models"
)

type RepositoryMenu interface {
	CreateMenu(string) error
	GetMenuType(int) (*models.Menu, error)
	GetAllMenuTypes() (*[]models.Menu, error)
	UpdateMenu(int, string) error
	DeleteMenu(int) error
}

type RepositoryCategory interface {
	GetAllCategories() ([]models.Category, error)
	CreateCategory(string, string) error
	GetCategory(int) (*models.Category, error)
	UpdateCategory(int, string) error
	DeleteCategory(int) error
}
