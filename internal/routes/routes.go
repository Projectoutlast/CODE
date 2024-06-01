package routes

import "github.com/gorilla/mux"

func SetUpRoutes(r *mux.Router) {
	main := r.PathPrefix("/").Subrouter()
	main.HandleFunc("/introduction", introduction).Methods("GET")
	main.HandleFunc("/main-proposals", mainProposals).Methods("GET")
	main.HandleFunc("/links-to-key-sections", linksToKeySections).Methods("GET")

	menu := r.PathPrefix("/menu").Subrouter()
	menu.HandleFunc("/main-menu", mainMenu).Methods("GET")
	menu.HandleFunc("/menu-for-lunch", menuForLunch).Methods("GET")
	menu.HandleFunc("/menu-for-catering", menuForCatering).Methods("GET")

	catering := r.PathPrefix("/catering").Subrouter()
	catering.HandleFunc("/catering-order", cateringOrder).Methods("GET")
	catering.HandleFunc("catering-faq", cateringFAQ).Methods("GET")

	about := r.PathPrefix("/about").Subrouter()
	about.HandleFunc("/our-history", ourHistory).Methods("GET")
	about.HandleFunc("/team", team).Methods("GET")
	about.HandleFunc("/mission-and-values", missionAndValues).Methods("GET")

	contacts := r.PathPrefix("/contacts").Subrouter()
	contacts.HandleFunc("/address", address).Methods("GET")
	contacts.HandleFunc("/phone", phone).Methods("GET")
	contacts.HandleFunc("/email", email).Methods("GET")
	contacts.HandleFunc("/social-network", socialNetwork).Methods("GET")

	r.HandleFunc("/articles-and-news", articleAndNews).Methods("GET")

	r.HandleFunc("/privacy-and-policy", privacyPolicy).Methods("GET")
	r.HandleFunc("/user-agreement", userAgreement).Methods("GET")
}
