package sqlite

import (
	"fmt"
	"log/slog"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestNewSQLiteRepository(t *testing.T) {
	db, _, _ := sqlmock.New()

	newRepository := NewSQLiteRepository(slog.Default(), db)
	assert.NotNil(t, newRepository.db)
	assert.NotNil(t, newRepository.log)
}

func TestGetAllMenuTypes(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewSQLiteRepository(slog.Default(), db)

	rows := sqlmock.NewRows([]string{"id", "menu_type"}).
		AddRow(1, "Main").
		AddRow(2, "Dessert")

	mock.ExpectQuery("SELECT id, menu_type FROM menu").WillReturnRows(rows)

	menuTypes, err := repo.GetAllMenuTypes()
	assert.NoError(t, err)
	assert.NotNil(t, menuTypes)
	assert.Len(t, *menuTypes, 2)
	assert.Equal(t, "Main", (*menuTypes)[0].Type)
	assert.Equal(t, "Dessert", (*menuTypes)[1].Type)

	mock.ExpectQuery("SELECT id, menu_type FROM menu").WillReturnError(fmt.Errorf("no rows in result set"))

	_, err = repo.GetAllMenuTypes()
	assert.Error(t, err)
}

func TestCreateMenu(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewSQLiteRepository(slog.Default(), db)

	mock.ExpectExec("INSERT INTO menu").
		WithArgs("Main").
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.CreateMenu("Main")
	assert.NoError(t, err)

	assert.NoError(t, mock.ExpectationsWereMet())

	mock.ExpectExec("INSERT INTO menu").
		WithArgs("Main").
		WillReturnError(fmt.Errorf("no rows in result set"))

	err = repo.CreateMenu("Main")
	assert.Error(t, err)
}

func TestUpdateMenu(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewSQLiteRepository(slog.Default(), db)

	mock.ExpectExec(`UPDATE menu SET menu_type = \? WHERE id = \?`).
		WithArgs("Appetizer", 1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.UpdateMenu(1, "Appetizer")
	assert.NoError(t, err)

	mock.ExpectExec(`UPDATE menu SET menu_type = \? WHERE id = \?`).
		WithArgs("Appetizer", 1).
		WillReturnError(fmt.Errorf("no rows in result set"))

	err = repo.UpdateMenu(1, "Appetizer")
	assert.Error(t, err)
}

func TestDeleteMenu(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewSQLiteRepository(slog.Default(), db)

	mock.ExpectExec("DELETE FROM menu WHERE id = ?").
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.DeleteMenu(1)
	assert.NoError(t, err)

	mock.ExpectExec("DELETE FROM menu WHERE id = ?").
		WithArgs(1).
		WillReturnError(fmt.Errorf("no rows in result set"))

	err = repo.DeleteMenu(1)
	assert.Error(t, err)
}

func TestIsMenuTypeExtst(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewSQLiteRepository(slog.Default(), db)

	mock.ExpectQuery("SELECT menu_type FROM menu WHERE menu_type = ?").
		WithArgs("Main").
		WillReturnRows(sqlmock.NewRows([]string{"menu_type"}).AddRow("Main"))

	exists := repo.isMenuTypeExist("Main")
	assert.True(t, exists)
	assert.NoError(t, mock.ExpectationsWereMet())

	mock.ExpectQuery("SELECT menu_type FROM menu WHERE menu_type = ?").
		WithArgs("Dessert").
		WillReturnError(fmt.Errorf("no rows in result set"))

	exists = repo.isMenuTypeExist("Dessert")
	assert.False(t, exists)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetMenuType(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewSQLiteRepository(slog.Default(), db)

	mock.ExpectQuery("SELECT id, menu_type FROM menu WHERE id = ?").
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "menu_type"}).AddRow(1, "Main"))

	menuType, err := repo.GetMenuType(1)
	assert.NoError(t, err)
	assert.NotNil(t, menuType)
	assert.Equal(t, "Main", menuType.Type)
	assert.NoError(t, mock.ExpectationsWereMet())

	mock.ExpectQuery("SELECT id, menu_type FROM menu WHERE id = ?").
		WithArgs(1).
		WillReturnError(fmt.Errorf("no rows in result set"))

	_, err = repo.GetMenuType(1)
	assert.Error(t, err)
}

func TestGetAllCategories(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewSQLiteRepository(slog.Default(), db)

	mock.ExpectQuery("SELECT id, menu_type_id, category_name FROM category_dish").
		WillReturnRows(sqlmock.NewRows([]string{"id", "menu_type_id", "category_name"}).
			AddRow(1, 1, "Main").
			AddRow(2, 2, "Dessert"))

	category, err := repo.GetAllCategories()
	assert.NoError(t, err)
	assert.NotNil(t, category)
	assert.Len(t, category, 2)
	assert.Equal(t, "Main", (category)[0].CategoryName)
	assert.Equal(t, "Dessert", (category)[1].CategoryName)
	assert.NoError(t, mock.ExpectationsWereMet())

	mock.ExpectQuery("SELECT id, menu_type_id, category_name FROM category_dish").
		WillReturnError(fmt.Errorf("no rows in result set"))

	category, err = repo.GetAllCategories()
	assert.Error(t, err)
	assert.Nil(t, category)
	assert.Len(t, category, 0)
}

func TestCreateCategory(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewSQLiteRepository(slog.Default(), db)

	mock.ExpectExec(`INSERT INTO category_dish \(menu_type_id, category_name\) VALUES \(\?, \?\)`).
		WithArgs("1", "Appetizers").
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.CreateCategory("1", "Appetizers")

	assert.NoError(t, err)

	mock.ExpectExec(`INSERT INTO category_dish \(menu_type_id, category_name\) VALUES \(\?, \?\)`).
		WithArgs("1", "Appetizers").
		WillReturnError(fmt.Errorf("no rows in result set"))

	err = repo.CreateCategory("1", "Appetizers")

	assert.Error(t, err)
}

func TestGetCategory(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewSQLiteRepository(slog.Default(), db)

	mock.ExpectQuery("SELECT id, menu_type_id, category_name FROM category_dish WHERE id = ?").
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "menu_type_id", "category_name"}).AddRow(1, 1, "Appetizers"))

	category, err := repo.GetCategory(1)
	assert.NoError(t, err)
	assert.NotNil(t, category)
	assert.Equal(t, "Appetizers", category.CategoryName)
	assert.NoError(t, mock.ExpectationsWereMet())

	mock.ExpectQuery("SELECT id, menu_type_id, category_name FROM category_dish WHERE id = ?").
		WithArgs(1).
		WillReturnError(fmt.Errorf("no rows in result set"))

	category, err = repo.GetCategory(1)
	assert.Error(t, err)
	assert.Nil(t, category)
}

func TestUpdateCategory(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewSQLiteRepository(slog.Default(), db)

	mock.ExpectExec(`UPDATE category_dish SET category_name = \? WHERE id = \?`).
		WithArgs("Appetizers", 1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.UpdateCategory(1, "Appetizers")
	assert.NoError(t, err)

	mock.ExpectExec(`UPDATE category_dish SET category_name = \? WHERE id = \?`).
		WithArgs("Appetizers", 1).
		WillReturnError(fmt.Errorf("no rows in result set"))

	err = repo.UpdateCategory(1, "Appetizers")
	assert.Error(t, err)
}

func TestDeleteCategory(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewSQLiteRepository(slog.Default(), db)

	mock.ExpectExec("DELETE FROM category_dish WHERE id = ?").
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(1, 1))
	
	err = repo.DeleteCategory(1)
	assert.NoError(t, err)

	mock.ExpectExec("DELETE FROM category_dish WHERE id = ?").
		WithArgs(1).
		WillReturnError(fmt.Errorf("no rows in result set"))
	
	err = repo.DeleteCategory(1)
	assert.Error(t, err)
}
