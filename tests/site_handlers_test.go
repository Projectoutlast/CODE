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

func TestAdminHandlers(t *testing.T) {

	_, adminHandlers := startTestHTTPServer(t)

	rw, req := createRWReq(http.MethodGet, "/admin/menu")
	adminHandlers.Index(rw, req)
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
