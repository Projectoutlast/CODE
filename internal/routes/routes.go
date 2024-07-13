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
	r.HandleFunc("/admin/menu/category/{id}", middleware.Logging(adminHandlers.TheCategory)).Methods("GET")
	r.HandleFunc("/admin/menu/category/edit/{category_id}", middleware.Logging(adminHandlers.EditCategory)).Methods("GET", "POST")
	r.HandleFunc("/admin/menu/category/delete/{category_id}", middleware.Logging(adminHandlers.DeleteCategory)).Methods("DELETE")

	r.HandleFunc("/admin/dishes", middleware.Logging(adminHandlers.Dishes)).Methods("GET")
	r.HandleFunc("/admin/menu/dish/{dish_id}", middleware.Logging(adminHandlers.TheDish)).Methods("GET")
	r.HandleFunc("/admin/menu/dish/edit/{dish_id}", middleware.Logging(adminHandlers.EditDish)).Methods("GET", "POST")
	r.HandleFunc("/admin/menu/dish/create", middleware.Logging(adminHandlers.CreateDish)).Methods("GET", "POST")
	r.HandleFunc("/admin/menu/dish/delete/{dish_id}", middleware.Logging(adminHandlers.DeleteDish)).Methods("DELETE")

	r.HandleFunc("/admin/events", middleware.Logging(adminHandlers.Events)).Methods("GET")
	r.HandleFunc("/admin/event/{event_id}", middleware.Logging(adminHandlers.TheEvent)).Methods("GET")
	r.HandleFunc("/admin/event/edit/{event_id}", middleware.Logging(adminHandlers.EditEvent)).Methods("GET", "POST")
	r.HandleFunc("/admin/event/create", middleware.Logging(adminHandlers.CreateEvent)).Methods("GET", "POST")
	r.HandleFunc("/admin/event/delete/{event_id}", middleware.Logging(adminHandlers.DeleteEvent)).Methods("DELETE")

	r.HandleFunc("/admin/employees", middleware.Logging(adminHandlers.Employees)).Methods("GET")
	r.HandleFunc("/admin/register-new-employee", middleware.Logging(adminHandlers.RegisterNewEmployee)).Methods("GET", "POST")
	r.HandleFunc("/admin/employee/{id}", middleware.Logging(adminHandlers.TheEmployee)).Methods("GET")
	r.HandleFunc("/admin/employee/edit/{id}", middleware.Logging(adminHandlers.EditEmployee)).Methods("GET", "POST")
	r.HandleFunc("/admin/employee/delete/{id}", middleware.Logging(adminHandlers.DeleteEmployee)).Methods("DELETE")
}

func SetUpFileServer(r *mux.Router, pathToStatic string) {
	fs := http.FileServer(utils.NeuteredFileSystem{Fs: http.Dir(pathToStatic)})
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
}
