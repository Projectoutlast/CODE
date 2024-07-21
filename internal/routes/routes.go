package routes

import (
	"code/internal/middleware"
	"code/internal/routes/admin"
	"code/internal/routes/site"
	"code/internal/utils"
	"net/http"

	"github.com/gorilla/mux"
)

func SetUpRoutes(
	r *mux.Router,
	siteHandlers *site.MainHandlers,
	adminHandlers *admin.AdminHandlers,
	middleware *middleware.Middleware,
) {
	// Site routes
	r.HandleFunc("/", middleware.Logging(siteHandlers.Index)).Methods("GET")
	r.HandleFunc("/main-menu", middleware.Logging(siteHandlers.MainMenu)).Methods("GET")
	r.HandleFunc("/menu-for-catering", middleware.Logging(siteHandlers.MenuForCatering)).Methods("GET")
	r.HandleFunc("/about", middleware.Logging(siteHandlers.About)).Methods("GET")
	r.HandleFunc("/contacts", middleware.Logging(siteHandlers.Contacts)).Methods("GET")
	r.HandleFunc("/news-and-events", middleware.Logging(siteHandlers.NewsAndEvents)).Methods("GET")
	r.HandleFunc("/privacy-policy", middleware.Logging(siteHandlers.PrivacyPolicy)).Methods("GET")
	r.HandleFunc("/user-agreement", middleware.Logging(siteHandlers.UserAgreement)).Methods("GET")

	// Admin routes
	r.HandleFunc("/admin/main", middleware.Logging(adminHandlers.Index)).Methods("GET")

	r.HandleFunc("/admin/menu", middleware.Logging(adminHandlers.Menu)).Methods("GET")
	r.HandleFunc("/admin/menu/create", middleware.Logging(adminHandlers.CreateMenuGet)).Methods("GET")
	r.HandleFunc("/admin/menu/create", middleware.Logging(adminHandlers.CreateMenuPost)).Methods("POST")
	r.HandleFunc("/admin/menu/edit/{type_id}", middleware.Logging(adminHandlers.EditMenuGet)).Methods("GET")
	r.HandleFunc("/admin/menu/edit/{type_id}", middleware.Logging(adminHandlers.EditMenuPost)).Methods("POST")
	r.HandleFunc("/admin/menu/delete/{type_id}", middleware.Logging(adminHandlers.DeleteMenu)).Methods("DELETE")

	r.HandleFunc("/admin/menu/category", middleware.Logging(adminHandlers.Categories)).Methods("GET")
	r.HandleFunc("/admin/menu/category/create", middleware.Logging(adminHandlers.CreateCategory)).Methods("GET")
	r.HandleFunc("/admin/menu/category/create", middleware.Logging(adminHandlers.CreateCategoryProcess)).Methods("POST")
	r.HandleFunc("/admin/menu/category/edit/{category_id}", middleware.Logging(adminHandlers.EditCategory)).Methods("GET")
	r.HandleFunc("/admin/menu/category/edit/{category_id}", middleware.Logging(adminHandlers.EditCategoryProcess)).Methods("POST")
	r.HandleFunc("/admin/menu/category/delete/{category_id}", middleware.Logging(adminHandlers.DeleteCategory)).Methods("DELETE")
}

func SetUpFileServer(r *mux.Router, pathToStatic string) {
	fs := http.FileServer(utils.NeuteredFileSystem{Fs: http.Dir(pathToStatic)})
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
}
