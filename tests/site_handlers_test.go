package tests_test

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
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"
)

func startTestHTTPServer() *site.MainHandlers {
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

	return siteHandlers
}

func TestMain(t *testing.T) {
	err := os.Chdir("..")
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()

	newLogger := logger.NewLogger(cfg.LogLevel)
	siteHandlers := site.NewMainHandlers(newLogger)
	adminHandlers := admin.NewAdminHandlers(newLogger)
	routes.SetUpRoutes(r, siteHandlers, adminHandlers, middleware.NewMiddleware(newLogger))
	routes.SetUpFileServer(r, cfg.StaticDir)

	rw := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)

	siteHandlers.Index(rw, req)

	require.Equal(t, http.StatusOK, rw.Code)

}
