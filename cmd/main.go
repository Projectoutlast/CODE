package main

import (
	"code/internal/config"
	"code/internal/routes"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
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

	r := mux.NewRouter()

	routes.SetUpRoutes(r)
	http.ListenAndServe(cfg.PORT, r)
}
