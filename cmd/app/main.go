package main

import (
	"code/internal/config"
	"code/internal/logger"
	"code/internal/middleware"
	"code/internal/routes"
	"code/internal/routes/admin"
	"code/internal/routes/site"
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

	newLogger := logger.NewLogger(cfg.LogLevel)

	siteHandlers := site.NewMainHandlers(newLogger)
	adminHandlers := admin.NewAdminHandlers(newLogger)

	middlewares := middleware.NewMiddleware(newLogger)

	r := mux.NewRouter()

	routes.SetUpRoutes(r, siteHandlers, adminHandlers, middlewares)
	routes.SetUpFileServer(r, cfg.StaticDir)

	newLogger.Info("Starting server on", slog.String("port", cfg.Port))
	err = http.ListenAndServe(cfg.Port, r)
	log.Fatal(err)
}
