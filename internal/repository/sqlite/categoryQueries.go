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

func (r *SQLiteRepository) GetCategory(categoryID int) (*models.Category, error) {
	stmt := `SELECT id, menu_type_id, category_name FROM category_dish WHERE id = ?`

	row := r.db.QueryRow(stmt, categoryID)
	if err := row.Err(); err != nil {
		r.log.Error(err.Error())
		return nil, err
	}

	var category models.Category
	if err := row.Scan(&category.ID, &category.MenuTypeID, &category.CategoryName); err != nil {
		r.log.Error(err.Error())
		return nil, err
	}

	return &category, nil
}

func (r *SQLiteRepository) UpdateCategory(categoryID int, categoryNewName string) error {
	stmt := `UPDATE category_dish SET category_name = ? WHERE id = ?`

	_, err := r.db.Exec(stmt, categoryNewName, categoryID)

	if err != nil {
		r.log.Error(err.Error())
		return err
	}

	return nil
}

func (r *SQLiteRepository) DeleteCategory(categoryID int) error {
	stmt := `DELETE FROM category_dish WHERE id = ?`

	_, err := r.db.Exec(stmt, categoryID)
	if err != nil {
		r.log.Error(err.Error())
		return err
	}

	return nil
}
