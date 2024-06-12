package main

import (
	"code/internal/config"
	"code/internal/logger"
	"code/internal/middleware"
	"code/internal/routes"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"log/slog"
	"net/http"
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

	handlers := routes.NewHandlers(newLogger)

	middlewares := middleware.NewMiddleware(newLogger)

	r := mux.NewRouter()
	routes.SetUpRoutes(r, handlers, middlewares)
	routes.SetUpFileServer(r)

	newLogger.Info("Starting server on", slog.String("port", cfg.PORT))
	err = http.ListenAndServe(cfg.PORT, r)
	log.Fatal(err)
}
