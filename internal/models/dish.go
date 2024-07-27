package models

type Dish struct {
	ID             int
	Name           string
	CategoryDishID int
	MenuTypeID     int
	Composition    string
	Description    string
	Price          float64
	Weight         int
	Image          []byte
	Tags           string
}

type CreateDish struct {
	MenuTypes  []Menu
	Categories []Category
}
