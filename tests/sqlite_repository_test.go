package tests_test

import (
	"code/internal/logger"
	"code/internal/repository/sqlite"
	"database/sql"
	"testing"

	"code/tests/suits"

	"github.com/stretchr/testify/require"

	_ "github.com/mattn/go-sqlite3"
)

func TestMenu(t *testing.T) {
	t.Parallel()

	logger := logger.NewLogger("dev")

	db, err := sql.Open("sqlite3", ":memory:")

	require.NoError(t, err)

	require.NoError(t, createTables(db))

	sqliteRepository := sqlite.NewSQLiteRepository(logger, db)

	require.NotNil(t, sqliteRepository)

	for _, menu := range suits.MenuTypes {
		err := sqliteRepository.CreateMenu(menu.MenuType)
		require.NoError(t, err)
	}

	newMenuType := "Безлактозное меню"
	err = sqliteRepository.UpdateMenu(4, newMenuType)
	require.NoError(t, err)

	menuType, err := sqliteRepository.GetMenuType(4)
	require.NoError(t, err)

	require.NotEqual(t, menuType, suits.MenuTypes[3].MenuType)

	require.NoError(t, sqliteRepository.DeleteMenu(4))

}

func createTables(db *sql.DB) error {
	stmts := []string{
		`CREATE TABLE IF NOT EXISTS menu (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		menu_type VARCHAR NOT NULL UNIQUE
		);`,
		`CREATE TABLE IF NOT EXISTS category_dish (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		menu_type_id INTEGER NOT NULL,
		category_name VARCHAR UNIQUE NOT NULL,
		FOREIGN KEY (menu_type_id) REFERENCES menu(id)
		);`,
		`CREATE TABLE IF NOT EXISTS dishes (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		dish_name VARCHAR UNIQUE NOT NULL,
		category_dish_id INTEGER NOT NULL,
		composition_of_the_dish VARCHAR NOT NULL,
		dish_description VARCHAR,
		price NUMERIC(4, 2) NOT NULL,
		dish_weight INTEGER NOT NULL,
		dish_image BLOB,
		tags VARCHAR ARRAY,
		FOREIGN KEY (category_dish_id) REFERENCES category_dish(id)
		);`,
		`CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		event_name VARCHAR NOT NULL,
		event_description VARCHAR,
		event_date TIMESTAMP DEFAULT current_timestamp,
		event_time TIMESTAMP DEFAULT current_timestamp
		);`,
		`CREATE TABLE IF NOT EXISTS admin_panel_users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_login VARCHAR NOT NULL UNIQUE,
		user_password VARCHAR NOT NULL,
		email VARCHAR NOT NULL UNIQUE,
		firstname VARCHAR NOT NULL,
		lastname VARCHAR,
		user_role VARCHAR NOT NULL DEFAULT 'менеджер',
		create_date TIMESTAMP DEFAULT current_timestamp,
		update_date TIMESTAMP DEFAULT current_timestamp,
		CHECK (user_role IN ('менеджер', 'управляющий'))
		);`,
	}

	for _, stmt := range stmts {
		_, err := db.Exec(stmt)
		if err != nil {
			return err
		}
	}

	return nil
}
