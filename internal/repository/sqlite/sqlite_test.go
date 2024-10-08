package sqlite

import (
	"code/internal/models"
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

func TestGetAllDishes(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewSQLiteRepository(slog.Default(), db)

	rows := mock.NewRows([]string{
		"id",
		"dish_name",
		"menu_type_id",
		"category_dish_id",
		"composition_of_the_dish",
		"dish_description",
		"price",
		"dish_weight",
		"dish_image",
		"tags",
	}).
		AddRow(1, "Beef Burger", 1, 1, "Beef, cheese, lettuce", "Beef with cheese and lettuce", 1000, 500, []byte{}, "Beef, Cheese, Lettuce").
		AddRow(2, "Chicken Burger", 1, 1, "Chicken, cheese, lettuce", "Chicken with cheese and lettuce", 1000, 500, []byte{}, "Chicken, Cheese, Lettuce")

	mock.ExpectQuery(`SELECT \* FROM dishes`).WillReturnRows(rows)

	dishes, err := repo.GetAllDishes()
	assert.NoError(t, err)
	assert.NotNil(t, dishes)
	assert.Equal(t, 2, len(dishes))

	mock.ExpectQuery(`SELECT \* FROM dishes`).WillReturnError(fmt.Errorf("no rows in result set"))

	_, err = repo.GetAllDishes()
	assert.Error(t, err)
}

func TestCreateDish(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewSQLiteRepository(slog.Default(), db)

	mock.ExpectExec(`INSERT INTO dishes \(
	dish_name,
	menu_type_id,
	category_dish_id,
	composition_of_the_dish,
	dish_description,
	price, 
	dish_weight,
	tags\) VALUES \(\?, \?, \?, \?, \?, \?, \?, \?\)`).
		WithArgs("Beef Burger", 1, 1, "Beef, cheese, lettuce", "Beef with cheese and lettuce", 10.00, 500, "Beef, Cheese, Lettuce").
		WillReturnResult(sqlmock.NewResult(1, 1))

	newDish := &models.Dish{
		Name:           "Beef Burger",
		CategoryDishID: 1,
		MenuTypeID:     1,
		Composition:    "Beef, cheese, lettuce",
		Description:    "Beef with cheese and lettuce",
		Price:          10.00,
		Weight:         500,
		Tags:           "Beef, Cheese, Lettuce",
	}
	err = repo.CreateNewDish(newDish)

	assert.NoError(t, err)

	mock.ExpectExec(`INSERT INTO dishes \(
		dish_name,
		menu_type_id,
		category_dish_id,
		composition_of_the_dish,
		dish_description,
		price, 
		dish_weight,
		tags\) VALUES \(\?, \?, \?, \?, \?, \?, \?, \?\)`).
		WithArgs().
		WillReturnError(fmt.Errorf("some error occured"))

	err = repo.CreateNewDish(&models.Dish{})
	assert.Error(t, err)
}

func TestGetDish(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewSQLiteRepository(slog.Default(), db)

	rows := sqlmock.NewRows([]string{
		"id",
		"dish_name",
		"menu_type_id",
		"category_dish_id",
		"composition_of_the_dish",
		"dish_description",
		"price",
		"dish_weight",
		"dish_image",
		"tags",
	}).
		AddRow(1, "Beef Burger", 1, 1, "Beef, cheese, lettuce", "Beef with cheese and lettuce", 1000, 500, []byte{}, "Beef, Cheese, Lettuce")

	mock.ExpectQuery(`SELECT \* FROM dishes WHERE id = \?`).
		WithArgs(1).
		WillReturnRows(rows)

	dishes, err := repo.GetDish(1)
	assert.NoError(t, err)
	assert.NotNil(t, dishes)

	mock.ExpectQuery(`SELECT \* FROM dishes WHERE id = \?`).WillReturnError(fmt.Errorf("no rows in result set"))

	_, err = repo.GetAllDishes()
	assert.Error(t, err)
}

func TestUpdateDish(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewSQLiteRepository(slog.Default(), db)

	mock.ExpectExec(`UPDATE dishes SET dish_name = \?, composition_of_the_dish = \?, dish_description = \?, price = \?, dish_weight = \?, tags = \? WHERE id = \?`).
		WithArgs("Beef Burger", "Beef, cheese, lettuce", "Beef with cheese and lettuce", 10.00, 500, "Beef, Cheese, Lettuce", 1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	dish := &models.Dish{
		ID:          1,
		Name:        "Beef Burger",
		Composition: "Beef, cheese, lettuce",
		Description: "Beef with cheese and lettuce",
		Price:       10.00,
		Weight:      500,
		Tags:        "Beef, Cheese, Lettuce",
	}

	err = repo.UpdateDish(dish)
	assert.NoError(t, err)

	mock.ExpectExec(`UPDATE dishes SET dish_name = \?, composition_of_the_dish = \?, dish_description = \?, price = \?, dish_weight = \?, tags = \? WHERE id = \?`).
		WithArgs().
		WillReturnError(fmt.Errorf("some error occured"))

	err = repo.UpdateDish(&models.Dish{})
	assert.Error(t, err)
}

func TestGetAllUsers(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewSQLiteRepository(slog.Default(), db)

	mock.ExpectQuery(`SELECT
		id, user_login, email, firstname, lastname, user_role, create_date, update_date
	FROM
		admin_panel_users`).
		WillReturnRows(sqlmock.NewRows([]string{"id", "user_login", "email", "firstname", "lastname", "user_role", "create_date", "update_date"}).
			AddRow(1, "Any", "test@example.com", "John", "Doe", "менеджер", "2024-07-31 00:00:00", "2024-07-31 12:00:00").
			AddRow(1, "Many", "best@example.com", "Jean", "Doe", "управляющий", "2024-07-31 00:00:00", "2024-07-31 12:00:00"))

	users, err := repo.GetAllUsers()
	assert.NoError(t, err)
	assert.NotNil(t, users)
	assert.Len(t, users, 2)
	assert.Equal(t, "test@example.com", (users)[0].Email)
	assert.Equal(t, "Many", (users)[1].Login)
	assert.NoError(t, mock.ExpectationsWereMet())

	mock.ExpectQuery(`SELECT
		id, user_login, email, firstname, lastname, user_role, create_date, update_date
	FROM
		admin_panel_users`).
		WillReturnError(fmt.Errorf("no rows in result set"))

	users, err = repo.GetAllUsers()
	assert.Error(t, err)
	assert.Nil(t, users)
	assert.Len(t, users, 0)
}

func TestRegisterUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewSQLiteRepository(slog.Default(), db)

	mock.ExpectExec(`INSERT INTO admin_panel_users \(
		user_login, user_password, email, firstname, lastname\) VALUES \(
		\?, \?, \?, \?, \?\)`).
		WithArgs("Any", "somepass", "test@example.com", "John", "Doe").
		WillReturnResult(sqlmock.NewResult(1, 1))

	user := &models.User{
		Login:        "Any",
		PasswordHash: "somepass",
		Email:        "test@example.com",
		Firstname:    "John",
		Lastname:     "Doe",
	}
	err = repo.RegisterUser(user)

	assert.NoError(t, err)

	mock.ExpectExec(`INSERT INTO admin_panel_users \(
		user_login, user_password, email, firstname, lastname\) VALUES \(
		\?, \?, \?, \?, \?\)`).
		WithArgs("Any", "somepass", "test@example.com", "John", "Doe").
		WillReturnError(fmt.Errorf("no rows in result set"))

	err = repo.RegisterUser(user)

	assert.Error(t, err)
}

func TestViewUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewSQLiteRepository(slog.Default(), db)

	mock.ExpectQuery(`SELECT
		id, user_login, email, firstname, lastname, user_role, create_date, update_date
	FROM
		admin_panel_users WHERE id = \?`).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "user_login", "email", "firstname", "lastname", "user_role", "create_date", "update_date"}).
			AddRow(1, "Any", "test@example.com", "John", "Doe", "менеджер", "2024-07-31 00:00:00", "2024-07-31 12:00:00"))

	user, err := repo.ViewUser(1)
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, "test@example.com", user.Email)
	assert.NoError(t, mock.ExpectationsWereMet())

	mock.ExpectQuery(`SELECT
		id, user_login, email, firstname, lastname, user_role, create_date, update_date
	FROM
		admin_panel_users WHERE id = \?`).
		WithArgs(1).
		WillReturnError(fmt.Errorf("no rows in result set"))

	user, err = repo.ViewUser(1)
	assert.Error(t, err)
	assert.Nil(t, user)
}

func TestUpdateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewSQLiteRepository(slog.Default(), db)

	mock.ExpectExec(`UPDATE admin_panel_users
	SET user_login = \?, email = \?, firstname = \?, lastname = \?
	WHERE id = \?`).
		WithArgs("Any", "test@example.com", "John", "Doe", 1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	user := &models.User{
		ID:        1,
		Login:     "Any",
		Email:     "test@example.com",
		Firstname: "John",
		Lastname:  "Doe",
	}
	err = repo.UpdateUser(user)

	assert.NoError(t, err)

	mock.ExpectExec(`UPDATE admin_panel_users
	SET user_login = \?, email = \?, firstname = \?, lastname = \?
	WHERE id = \?`).
		WithArgs("Any", "test@example.com", "John", "Doe", 1).
		WillReturnError(fmt.Errorf("no rows in result set"))

	err = repo.UpdateUser(user)

	assert.Error(t, err)
}

func TestDeleteUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewSQLiteRepository(slog.Default(), db)

	mock.ExpectExec(`DELETE admin_panel_users WHERE id = ?`).
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.DeleteUser(1)
	assert.NoError(t, err)

	mock.ExpectExec(`DELETE FROM admin_panel_users WHERE id = \?`).
		WithArgs(1).
		WillReturnError(fmt.Errorf("no rows in result set"))

	err = repo.DeleteUser(1)
	assert.Error(t, err)
}
