package sqlite

import (
	"code/internal/repository/errors"
)

func (r *SQLiteRepository) CreateCategory(menuType, categoryName string) error {
	if !r.isMenuTypeExist(menuType) {
		return errors.ErrNoMenuType
	}

	stmt := `INSERT INTO category (menu_type, category_dish) VALUES (?, ?)`

	_, err := r.db.Exec(stmt, menuType, categoryName)
	if err != nil {
		r.log.Error(err.Error())
		return err
	}

	return nil
}

func (r *SQLiteRepository) UpdateCategory(editedCategoryName, categoryName string) error {
	stmt := `UPDATE category SET category_dish = ? WHERE category_dish = ?`

	_, err := r.db.Exec(stmt, editedCategoryName, categoryName)
	if err != nil {
		r.log.Error(err.Error())
		return err
	}

	return nil
}

func (r *SQLiteRepository) DeleteCategory(categoryName string) error {
	stmt := `DELETE FROM category WHERE category_dish = ?`

	_, err := r.db.Exec(stmt, categoryName)
	if err != nil {
		r.log.Error(err.Error())
		return err
	}

	return nil
}
