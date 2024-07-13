package sqlite

import (
	"code/internal/repository/models"
)

func (r *SQLiteRepository) GetAllCategories() ([]models.Category, error) {
	stmt := `SELECT id, menu_type_id, category_name FROM category_dish`

	rows, err := r.db.Query(stmt)
	if err != nil {
		r.log.Error(err.Error())
		return nil, err
	}
	defer rows.Close()

	var categories []models.Category

	for rows.Next() {
		var category models.Category
		if err := rows.Scan(&category.ID, &category.MenuTypeID, &category.CategoryName); err != nil {
			r.log.Error(err.Error())
			return nil, err
		}

		categories = append(categories, category)
	}

	return categories, nil
}

func (r *SQLiteRepository) CreateCategory(menuTypeID, category string) error {
	stmt := `INSERT INTO category_dish (menu_type_id, category_name) VALUES (?, ?)`

	_, err := r.db.Exec(stmt, menuTypeID, category)

	if err != nil {
		r.log.Error(err.Error())
		return err
	}

	return nil
}
