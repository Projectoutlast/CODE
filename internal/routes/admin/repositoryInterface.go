package admin

import (
	"code/internal/models"
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

type RepositoryDish interface {
	GetAllDishes() ([]models.Dish, error)
	CreateNewDish(*models.Dish) error
	GetDish(int) (*models.Dish, error)
	UpdateDish(*models.Dish) error
	DeleteDish(int) error
}

type RepositoryEmployee interface {
	GetAllUsers() ([]models.User, error)
	RegisterUser(*models.User) error
	ViewUser(int) (*models.User, error)
	UpdateUser(*models.User) error
	DeleteUser(int) error
}
