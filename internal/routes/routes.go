package routes

import (
	"code/internal/middleware"
	"code/internal/utils"
	"github.com/gorilla/mux"
	"net/http"
)

func SetUpRoutes(r *mux.Router, handlers *Handlers, middleware *middleware.Middleware) {
	r.HandleFunc("/", middleware.Logging(handlers.index)).Methods("GET")

	r.HandleFunc("/main-menu", middleware.Logging(handlers.mainMenu)).Methods("GET")
	r.HandleFunc("/menu-for-catering", middleware.Logging(handlers.menuForCatering)).Methods("GET")

	r.HandleFunc("/about", middleware.Logging(handlers.about)).Methods("GET")

	r.HandleFunc("/contacts", middleware.Logging(handlers.contacts)).Methods("GET")

	r.HandleFunc("/news-and-events", middleware.Logging(handlers.newsAndEvents)).Methods("GET")

	r.HandleFunc("/privacy-policy", middleware.Logging(handlers.privacyPolicy)).Methods("GET")
	r.HandleFunc("/user-agreement", middleware.Logging(handlers.userAgreement)).Methods("GET")
}

func SetUpFileServer(r *mux.Router, pathToStatic string) {
	fs := http.FileServer(utils.NeuteredFileSystem{Fs: http.Dir(pathToStatic)})
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static", fs))
}
