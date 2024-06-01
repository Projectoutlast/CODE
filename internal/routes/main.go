package routes

import (
	"html/template"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./ui/html/base.layout.html")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	if err := tmpl.Execute(w, nil); err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func introduction(w http.ResponseWriter, r *http.Request) {
	panic("not implemented")
}

func mainProposals(w http.ResponseWriter, r *http.Request) {
	panic("not implemented")
}

func linksToKeySections(w http.ResponseWriter, r *http.Request) {
	panic("not implemented")
}
