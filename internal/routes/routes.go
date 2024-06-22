package routes

import (
	"code/internal/middleware"
	"code/internal/routes/site"
	"code/internal/utils"
	"net/http"

	"github.com/gorilla/mux"
)

func SetUpRoutes(r *mux.Router, siteHandlers *site.Handlers, middleware *middleware.Middleware) {
	r.HandleFunc("/", middleware.Logging(siteHandlers.Index)).Methods("GET")

	r.HandleFunc("/main-menu", middleware.Logging(siteHandlers.MainMenu)).Methods("GET")
	r.HandleFunc("/menu-for-catering", middleware.Logging(siteHandlers.MenuForCatering)).Methods("GET")

	r.HandleFunc("/about", middleware.Logging(siteHandlers.About)).Methods("GET")

	r.HandleFunc("/contacts", middleware.Logging(siteHandlers.Contacts)).Methods("GET")

	r.HandleFunc("/news-and-events", middleware.Logging(siteHandlers.NewsAndEvents)).Methods("GET")

	r.HandleFunc("/privacy-policy", middleware.Logging(siteHandlers.PrivacyPolicy)).Methods("GET")
	r.HandleFunc("/user-agreement", middleware.Logging(siteHandlers.UserAgreement)).Methods("GET")
}

func SetUpFileServer(r *mux.Router, pathToStatic string) {
	fs := http.FileServer(utils.NeuteredFileSystem{Fs: http.Dir(pathToStatic)})
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static", fs))
}
