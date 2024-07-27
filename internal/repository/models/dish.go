package models

type Dish struct {
	ID                   int
	DishName             string
	CategoryDishID       int
	CompositionOnTheDish string
	DishDescription      string
	Price                float64
	DishWeight           int
	DishImage            []byte
	Tags                 []string
}

type CreateDish struct {
	MenuTypes  []Menu
	Categories []Category
}
