package sqlite

import (
	"code/internal/repository/models"

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
			&dish.DishName,
			&dish.CategoryDishID,
			&dish.CompositionOnTheDish,
			&dish.DishDescription,
			&dish.Price,
			&dish.DishWeight,
			&dish.DishImage,
			&tags,
		); err != nil {
			r.log.Error(err.Error())
			return nil, err
		}

		dish.Tags = []string{tags}
		dishes = append(dishes, dish)
	}

	return dishes, nil
}
