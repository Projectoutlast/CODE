package sqlite

import (
	"code/internal/models"
)

func (r *SQLiteRepository) GetAllUsers() ([]models.User, error) {
	stmt := `
	SELECT
		id, user_login, email, firstname, lastname, user_role, create_date, update_date
	FROM
		admin_panel_users`

	rows, err := r.db.Query(stmt)
	if err != nil {
		r.log.Error(err.Error())
		return nil, err
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User
		if err := rows.Scan(
			&user.ID,
			&user.Login,
			&user.Email,
			&user.Firstname,
			&user.Lastname,
			&user.UserRole,
			&user.Create,
			&user.Update,
		); err != nil {
			r.log.Error(err.Error())
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (r *SQLiteRepository) RegisterUser(user *models.User) error {
	stmt := `INSERT INTO admin_panel_users (
		user_login, user_password, email, firstname, lastname) VALUES (
		?, ?, ?, ?, ?)`

	_, err := r.db.Exec(stmt, user.Login, user.PasswordHash, user.Email, user.Firstname, user.Lastname)
	if err != nil {
		r.log.Error(err.Error())
		return err
	}

	return nil
}

func (r *SQLiteRepository) ViewUser(userID int) (*models.User, error) {
	stmt := `SELECT
		id, user_login, email, firstname, lastname, user_role, create_date, update_date
	FROM
		admin_panel_users
	WHERE id = ?`

	row := r.db.QueryRow(stmt, userID)
	if err := row.Err(); err != nil {
		r.log.Error(err.Error())
		return nil, err
	}

	var user models.User
	if err := row.Scan(
		&user.ID,
		&user.Login,
		&user.Email,
		&user.Firstname,
		&user.Lastname,
		&user.UserRole,
		&user.Create,
		&user.Update,
	); err != nil {
		r.log.Error(err.Error())
		return nil, err
	}

	return &user, nil
}

func (r *SQLiteRepository) UpdateUser(user *models.User) error {
	stmt := `UPDATE admin_panel_users
	SET user_login = ?, email = ?, firstname = ?, lastname = ?
	WHERE id = ?`

	_, err := r.db.Exec(stmt, user.Login, user.Email, user.Firstname, user.Lastname, user.ID)
	if err != nil {
		r.log.Error(err.Error())
		return err
	}

	return nil
}

func (r *SQLiteRepository) DeleteUser(userID int) error {
	stmt := `DELETE admin_panel_users WHERE id = ?`

	_, err := r.db.Exec(stmt, userID)
	if err != nil {
		r.log.Error(err.Error())
		return err
	}

	return nil
}
