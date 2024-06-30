package tests_test

import (
	"code/internal/logger"
	"code/internal/repository/sqlite"
	"database/sql"
	"log"
	"testing"

	"code/tests/suits"

	"github.com/stretchr/testify/require"

	_ "github.com/mattn/go-sqlite3"
)

func TestMenu(t *testing.T) {
	logger := logger.NewLogger("dev")
	sqliteRepository := sqlite.NewSQLiteRepository(logger, testDBInit())

	require.NotNil(t, sqliteRepository)

	for _, menu := range suits.MenuTypes {
		err := sqliteRepository.CreateMenu(menu.MenuType)
		require.NoError(t, err)
	}

	err := sqliteRepository.DeleteMenu(suits.MenuTypes[3].MenuType)
	require.NoError(t, err)

	newMenuType := "Безглютеновое меню"
	err = sqliteRepository.UpdateMenu(newMenuType, suits.MenuTypes[4].MenuType)
	require.NoError(t, err)
	require.NotEqual(t, suits.MenuTypes[1].MenuType, newMenuType)

}

func testDBInit() *sql.DB {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}

	stmt := `
	CREATE TABLE IF NOT EXISTS menu (
    	id INTEGER PRIMARY KEY AUTOINCREMENT,
    	menu_type VARCHAR NOT NULL
	);
	`
	_, err = db.Exec(stmt)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
