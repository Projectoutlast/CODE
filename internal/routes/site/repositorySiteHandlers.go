package site

type RepositoryMenu interface {
	CreateMenu(string) error
	UpdateMenu(string, string) error
	DeleteMenu(string) error
}

type RepositoryCategory interface {
	CreateCategory(string, string) error
	UpdateCategory(string, string) error
	DeleteCategory(string) error
}