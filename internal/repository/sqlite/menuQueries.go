package sqlite

import (
	"code/internal/models"
)

func (r *SQLiteRepository) GetAllMenuTypes() (*[]models.Menu, error) {
	stmt := `SELECT id, menu_type FROM menu`

	rows, err := r.db.Query(stmt)
	if err != nil {
		r.log.Error(err.Error())
		return nil, err
	}
	defer rows.Close()

	var menuTypes []models.Menu

	for rows.Next() {
		var menuType models.Menu
		if err := rows.Scan(&menuType.ID, &menuType.Type); err != nil {
			r.log.Error(err.Error())
			return nil, err
		}
		menuTypes = append(menuTypes, menuType)
	}

	return &menuTypes, nil
}

func (r *SQLiteRepository) CreateMenu(menuType string) error {
	stmt := `INSERT INTO menu (menu_type) VALUES (?)`

	_, err := r.db.Exec(stmt, menuType)
	if err != nil {
		r.log.Error(err.Error())
		return err
	}

	return nil
}

func (r *SQLiteRepository) UpdateMenu(menuTypeID int, editedMenuType string) error {
	stmt := `UPDATE menu SET menu_type = ? WHERE id = ?`

	_, err := r.db.Exec(stmt, editedMenuType, menuTypeID)
	if err != nil {
		r.log.Error(err.Error())
		return err
	}

	return nil
}

func (r *SQLiteRepository) DeleteMenu(menuTypeID int) error {
	stmt := `DELETE FROM menu WHERE id = ?`

	_, err := r.db.Exec(stmt, menuTypeID)
	if err != nil {
		r.log.Error(err.Error())
		return err
	}

	return nil
}

func (r *SQLiteRepository) isMenuTypeExist(menuType string) bool {
	stmt := `SELECT menu_type FROM menu WHERE menu_type = ?`

	row := r.db.QueryRow(stmt, menuType)
	if err := row.Err(); err != nil {
		r.log.Error(err.Error())
		return false
	}

	return true
}

func (r *SQLiteRepository) GetMenuType(menuTypeID int) (*models.Menu, error) {
	stmt := `SELECT id, menu_type FROM menu WHERE id = ?`

	row := r.db.QueryRow(stmt, menuTypeID)
	if err := row.Err(); err != nil {
		r.log.Error(err.Error())
		return nil, err
	}

	var menuType models.Menu
	if err := row.Scan(&menuType.ID, &menuType.Type); err != nil {
		r.log.Error(err.Error())
		return nil, err
	}

	return &menuType, nil
}
