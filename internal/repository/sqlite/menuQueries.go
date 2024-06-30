package sqlite

func (r *SQLiteRepository) CreateMenu(menuType string) error {
	stmt := `INSERT INTO menu (menu_type) VALUES (?)`

	_, err := r.db.Exec(stmt, menuType)
	if err != nil {
		r.log.Error(err.Error())
		return err
	}

	return nil
}

func (r *SQLiteRepository) UpdateMenu(editedMenuType, menuType string) error {
	stmt := `UPDATE menu SET menu_type = ? WHERE menu_type = ?`

	_, err := r.db.Exec(stmt, editedMenuType, menuType)
	if err != nil {
		r.log.Error(err.Error())
		return err
	}

	return nil
}

func (r *SQLiteRepository) DeleteMenu(menuType string) error {
	stmt := `DELETE FROM menu WHERE menu_type = ?`

	_, err := r.db.Exec(stmt, menuType)
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
