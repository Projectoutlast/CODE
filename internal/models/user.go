package models

type User struct {
	ID           int    `db:"id"`
	Login        string `db:"user_login"`
	PasswordHash string `db:"user_password"`
	Email        string `db:"email"`
	Firstname    string `db:"firstname"`
	Lastname     string `db:"lastname"`
	UserRole     string `db:"user_role"`
	Create       string `db:"create_date"`
	Update       string `db:"update_date"`
}
