package sqlite

import (
	"code/internal/models"

	_ "github.com/mattn/go-sqlite3"
)

func (r *SQLiteRepository) GetAllDishes() ([]models.Dish, error) {
	stmt := `SELECT * FROM dishes`

	rows, err := r.db.Query(stmt)
	if err != nil {
		r.log.Error(err.Error())
		return nil, err
	}
	defer rows.Close()

	var dishes []models.Dish

	for rows.Next() {
		var dish models.Dish
		var tags string
		if err := rows.Scan(
			&dish.ID,
			&dish.Name,
			&dish.MenuTypeID,
			&dish.CategoryDishID,
			&dish.Composition,
			&dish.Description,
			&dish.Price,
			&dish.Weight,
			&dish.Image,
			&tags,
		); err != nil {
			r.log.Error(err.Error())
			return nil, err
		}

		dish.Tags = tags
		dishes = append(dishes, dish)
	}

	return dishes, nil
}

func (r *SQLiteRepository) CreateNewDish(dish *models.Dish) error {
	stmt := `INSERT INTO dishes (
	dish_name,
	menu_type_id,
	category_dish_id,
	composition_of_the_dish,
	dish_description,
	price, 
	dish_weight,
	tags) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`

	_, err := r.db.Exec(
		stmt,
		dish.Name,
		dish.MenuTypeID,
		dish.CategoryDishID,
		dish.Composition,
		dish.Description,
		dish.Price,
		dish.Weight,
		dish.Tags,
	)

	if err != nil {
		r.log.Error(err.Error())
		return err
	}

	return err
}
