package site

type RepositoryMenu interface {
	CreateMenu(string) error
	ReadMenu() ([]string, error)
	UpdateMenu(string, int) error
	DeleteMenu(int)
}