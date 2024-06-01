package main

import (
	"code/site/routes"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	routes.SetUpRoutes(r)

	http.ListenAndServe(":8080", r)
}
