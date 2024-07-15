package tests_test

import (
	"code/internal/config"
	"code/internal/logger"
	"code/internal/middleware"
	"code/internal/repository/sqlite"
	"code/internal/routes"
	"code/internal/routes/admin"
	"code/internal/routes/site"
	"database/sql"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"
)

func TestMenu(t *testing.T) {

	_, adminHandlers := startTestHTTPServer(t)

	// MENU
	rw, req := createRWReq(http.MethodGet, "/admin/menu")
	adminHandlers.Menu(rw, req)
	require.Equal(t, http.StatusOK, rw.Code)
	strings.Contains(rw.Body.String(), "Наименование типа меню")

	rw, req = createRWReq(http.MethodGet, "/admin/menu/create")
	adminHandlers.CreateMenuGet(rw, req)
	require.Equal(t, http.StatusOK, rw.Code)
	strings.Contains(rw.Body.String(), "Введите название")

	rw, req = createRWReq(http.MethodPost, "/admin/menu/create")
	req.Form = url.Values{}
	req.Form.Set("type", "Основное")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	adminHandlers.CreateMenuPost(rw, req)
	require.Equal(t, http.StatusSeeOther, rw.Code)

	rw, req = createRWReq(http.MethodPost, "/admin/menu/create")
	req.Form = url.Values{}
	req.Form.Set("type", "")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	adminHandlers.CreateMenuPost(rw, req)
	require.Equal(t, http.StatusBadRequest, rw.Code)

	rw, req = createRWReq(http.MethodGet, "/admin/menu/edit/hello")
	req = mux.SetURLVars(req, map[string]string{"type_id": "hello"})
	adminHandlers.EditMenuGet(rw, req)
	require.Equal(t, http.StatusBadRequest, rw.Code)

	rw, req = createRWReq(http.MethodGet, "/admin/menu/edit/1")
	req = mux.SetURLVars(req, map[string]string{"type_id": "1"})
	adminHandlers.EditMenuGet(rw, req)
	require.Equal(t, http.StatusOK, rw.Code)
	strings.Contains(rw.Body.String(), "Редактирование типа меню")

	rw, req = createRWReq(http.MethodPost, "/admin/menu/edit/1")
	req.Form = url.Values{}
	req.Form.Set("type", "")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req = mux.SetURLVars(req, map[string]string{"type_id": "1"})
	adminHandlers.EditMenuPost(rw, req)
	require.Equal(t, http.StatusBadRequest, rw.Code)

	rw, req = createRWReq(http.MethodPost, "/admin/menu/edit/1")
	req.Form = url.Values{}
	req.Form.Set("type", "Измененное")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req = mux.SetURLVars(req, map[string]string{"type_id": "1"})
	adminHandlers.EditMenuPost(rw, req)
	require.Equal(t, http.StatusSeeOther, rw.Code)

	rw, req = createRWReq(http.MethodPost, "/admin/menu/edit/hello")
	req.Form = url.Values{}
	req.Form.Set("type", "Измененное")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req = mux.SetURLVars(req, map[string]string{"type_id": "hello"})
	adminHandlers.EditMenuPost(rw, req)
	require.Equal(t, http.StatusBadRequest, rw.Code)

	rw, req = createRWReq(http.MethodGet, "/admin/menu/delete/test")
	req = mux.SetURLVars(req, map[string]string{"type_id": "test"})
	adminHandlers.DeleteMenu(rw, req)
	require.Equal(t, http.StatusBadRequest, rw.Code)

	rw, req = createRWReq(http.MethodGet, "/admin/menu/delete/100")
	req = mux.SetURLVars(req, map[string]string{"type_id": "100"})
	adminHandlers.DeleteMenu(rw, req)
	require.Equal(t, http.StatusSeeOther, rw.Code)

	rw, req = createRWReq(http.MethodGet, "/admin/menu/delete/1")
	req = mux.SetURLVars(req, map[string]string{"type_id": "1"})
	adminHandlers.DeleteMenu(rw, req)
	require.Equal(t, http.StatusSeeOther, rw.Code)

	// CATEGORIES
	rw, req = createRWReq(http.MethodPost, "/admin/menu/category/create")
	req.Form = url.Values{}
	req.Form.Set("menuType", "1")
	req.Form.Set("category", "Салаты")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	adminHandlers.CreateCategoryProcess(rw, req)
	require.Equal(t, http.StatusSeeOther, rw.Code)

	rw, req = createRWReq(http.MethodGet, "/admin/menu/category")
	adminHandlers.Categories(rw, req)
	require.Equal(t, http.StatusOK, rw.Code)
	strings.Contains(rw.Body.String(), "Салаты")

	rw, req = createRWReq(http.MethodGet, "/admin/menu/category/create")
	adminHandlers.CreateCategory(rw, req)
	require.Equal(t, http.StatusOK, rw.Code)
	strings.Contains(rw.Body.String(), "Выберите тип меню, к которому относится новая категория:")

	rw, req = createRWReq(http.MethodGet, "/admin/menu/category/edit/1")
	req = mux.SetURLVars(req, map[string]string{"category_id": "1"})
	adminHandlers.EditCategory(rw, req)
	require.Equal(t, http.StatusOK, rw.Code)
	strings.Contains(rw.Body.String(), "Редактирование категории")

	rw, req = createRWReq(http.MethodPost, "/admin/menu/category/edit/1")
	req.Form = url.Values{}
	req.Form.Set("category", "Измененная категория")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req = mux.SetURLVars(req, map[string]string{"category_id": "1"})
	adminHandlers.EditCategoryProcess(rw, req)
	require.Equal(t, http.StatusSeeOther, rw.Code)

}

func createRWReq(method, route string) (*httptest.ResponseRecorder, *http.Request) {
	rw := httptest.NewRecorder()
	req := httptest.NewRequest(method, route, nil)
	return rw, req
}

func startTestHTTPServer(t *testing.T) (*site.MainHandlers, *admin.AdminHandlers) {
	os.Chdir("..")
	fmt.Println(os.Getwd())
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	newLogger := logger.NewLogger(cfg.LogLevel)

	dsn := fmt.Sprintf("file:%s?mode=memory&cache=shared", t.Name())
	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		log.Fatal(err)
	}

	require.NoError(t, createTables(db))

	sqliteRepository := sqlite.NewSQLiteRepository(newLogger, db)

	siteHandlers := site.NewMainHandlers(newLogger)

	adminHandlers := admin.NewAdminHandlers(sqliteRepository, newLogger, sqliteRepository)

	middlewares := middleware.NewMiddleware(newLogger)

	r := mux.NewRouter()

	routes.SetUpRoutes(r, siteHandlers, adminHandlers, middlewares)
	routes.SetUpFileServer(r, cfg.StaticDir)

	newLogger.Info("Starting server on", slog.String("port", cfg.Port))
	go http.ListenAndServe(cfg.Port, r)

	return siteHandlers, adminHandlers
}
