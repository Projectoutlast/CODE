package main

import (
	"code/internal/config"
	"code/internal/logger"
	"code/internal/middleware"
	"code/internal/repository/sqlite"
	"code/internal/routes"
	"code/internal/routes/admin"
	"code/internal/routes/site"
	"database/sql"
	"log"
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("sqlite3", cfg.SQLite.DataSource)
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	newLogger := logger.NewLogger(cfg.LogLevel)

	sqliteRepository := sqlite.NewSQLiteRepository(newLogger, db)

	siteHandlers := site.NewMainHandlers(newLogger)
	adminHandlers := admin.NewAdminHandlers(sqliteRepository, sqliteRepository, newLogger, sqliteRepository, sqliteRepository)

	middlewares := middleware.NewMiddleware(newLogger)

	r := mux.NewRouter()

	routes.SetUpRoutes(r, siteHandlers, adminHandlers, middlewares)
	routes.SetUpFileServer(r, cfg.StaticDir)

	newLogger.Info("Starting server on", slog.String("port", cfg.Port))
	err = http.ListenAndServe(cfg.Port, r)
	log.Fatal(err)
}
