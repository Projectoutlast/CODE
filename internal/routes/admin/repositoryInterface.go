package admin

import (
	"code/internal/repository/models"
)

type RepositoryMenu interface {
	CreateMenu(string) error
	GetAllMenuTypes() (*[]models.Menu, error)
	UpdateMenu(int, string) error
	DeleteMenu(int) error
}

type RepositoryCategory interface {
	CreateCategory(string, string) error
	UpdateCategory(string, string) error
	DeleteCategory(string) error
}
