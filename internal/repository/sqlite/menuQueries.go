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

func (r *SQLiteRepository) ReadMenu() ([]string, error) {
	stmt := `SELECT menu_type FROM menu`

	rows, err := r.db.Query(stmt)
	if err != nil {
		r.log.Error(err.Error())
		return nil, err
	}
	defer rows.Close()

	var menuTypes []string
	for rows.Next() {
		var menuType string
		if err := rows.Scan(&menuType); err != nil {
			r.log.Error(err.Error())
			return nil, err
		}
		
		menuTypes = append(menuTypes, menuType)
	}

	if err := rows.Err(); err != nil {
		r.log.Error(err.Error())
		return nil, err
	}

	return menuTypes, nil
}

func (r *SQLiteRepository) UpdateMenu(editedMenuType string, id int) error {
	stmt := `UPDATE menu SET menu_type = ? WHERE id = ?`

	_, err := r.db.Exec(stmt, editedMenuType, id)
	if err != nil {
		r.log.Error(err.Error())
		return err
	}

	return nil
}

func (r *SQLiteRepository) DeleteMenu(id int) error {
	stmt := `DELETE FROM menu WHERE id = ?`

	_, err := r.db.Exec(stmt, id)
	if err != nil {
		r.log.Error(err.Error())
		return err
	}

	return nil
}
