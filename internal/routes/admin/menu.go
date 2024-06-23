package admin

import (
	"fmt"
	"html/template"
	"net/http"
)

func (h *AdminHandlers) Menu(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/html/admin/menu.page.html",
	}

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		h.log.Error(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if err = tmpl.Execute(w, nil); err != nil {
		h.log.Error(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func (h *AdminHandlers) TheMenuType(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "The menu type page")
}

func (h *AdminHandlers) EditMenu(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Edit menu page")
}

func (h *AdminHandlers) CreateMenu(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Edit menu page")
}

func (h *AdminHandlers) DeleteMenu(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Delete menu page")
}
