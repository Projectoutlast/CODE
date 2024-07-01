package site

type RepositoryMenu interface {
	CreateMenu(string) error
	UpdateMenu(int, string) error
	DeleteMenu(int) error
}

type RepositoryCategory interface {
	CreateCategory(string, string) error
	UpdateCategory(string, string) error
	DeleteCategory(string) error
}