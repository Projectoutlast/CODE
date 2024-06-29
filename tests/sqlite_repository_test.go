package tests_test

import (
	"code/internal/logger"
	"code/internal/repository/sqlite"
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStartSQLiteRepository(t *testing.T) {
	logger := logger.NewLogger("dev")
	sqliteRepository := sqlite.NewSQLiteRepository(logger, testDBInit())

	require.NotNil(t, sqliteRepository)

	newMenu := []string{"menu1", "menu2", "menu3", "menu4", "menu5", "menu6"}

	for _, menu := range newMenu {
		err := sqliteRepository.CreateMenu(menu)
		require.NoError(t, err)
	}

	menuTypes, err := sqliteRepository.ReadMenu()
	require.NoError(t, err)
	require.Equal(t, newMenu, menuTypes)

	err = sqliteRepository.DeleteMenu(1)
	require.NoError(t, err)

	err = sqliteRepository.UpdateMenu("menu100500", 1)
	require.NoError(t, err)
}

func testDBInit() *sql.DB {
	err := os.Chdir("..")
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}

	stmt := `CREATE TABLE IF NOT EXISTS menu (id SERIAL PRIMARY KEY, menu_type VARCHAR NOT NULL);`
	_, err = db.Exec(stmt)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
