package routes

import (
	"database/sql"
	"log"
	"log/slog"
	"testing"

	"github.com/gorilla/mux"

	"code/internal/middleware"
	"code/internal/repository/sqlite"

	"code/internal/routes/admin"
	"code/internal/routes/site"
)

func TestSetUpRoutes(t *testing.T) {

	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	logger := slog.Default()

	sqliteRepository := sqlite.NewSQLiteRepository(logger, db)

	siteHandlers := site.NewMainHandlers(logger)
	adminHandlers := admin.NewAdminHandlers(sqliteRepository, sqliteRepository, logger, sqliteRepository)

	middlewares := middleware.NewMiddleware(logger)

	r := mux.NewRouter()

	SetUpRoutes(r, siteHandlers, adminHandlers, middlewares)
	SetUpFileServer(r, "./ui/static")

}
