package routes

import (
	"code/internal/utils"
	"github.com/gorilla/mux"
	"net/http"
)

func SetUpRoutes(r *mux.Router) {
	r.HandleFunc("/", index).Methods("GET")

	r.HandleFunc("/main-menu", mainMenu).Methods("GET")
	r.HandleFunc("/menu-for-catering", menuForCatering).Methods("GET")

	r.HandleFunc("/about", about).Methods("GET")

	r.HandleFunc("/contacts", contacts)

	r.HandleFunc("/articles-and-news", articleAndNews).Methods("GET")

	r.HandleFunc("/privacy-and-policy", privacyPolicy).Methods("GET")
	r.HandleFunc("/user-agreement", userAgreement).Methods("GET")
}

func SetUpFileServer(r *mux.Router) {
	fs := http.FileServer(utils.NeuteredFileSystem{Fs: http.Dir("./ui/static")})
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static", fs))
}
