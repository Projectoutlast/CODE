package admin

import (
	"fmt"
	"html/template"
	"net/http"
)

func (h *AdminHandlers) Menu(w http.ResponseWriter, r *http.Request) {
	menuTypes, err := h.menu.GetAllMenuTypes()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	response := MenuResponse{
		PageTitle: "Меню",
		MenuTypes: *menuTypes,
	}

	files := []string{
		"./ui/html/admin/menu.page.html",
		baseHTMLLayout,
	}

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		h.log.Error(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if err = tmpl.Execute(w, response); err != nil {
		h.log.Error(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func (h *AdminHandlers) TheMenuType(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/html/admin/menu_type_view.page.html",
		baseHTMLLayout,
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

func (h *AdminHandlers) EditMenu(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/html/admin/menu_type_edit.page.html",
		baseHTMLLayout,
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

func (h *AdminHandlers) CreateMenuGet(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/html/admin/menu_type_create.page.html",
		baseHTMLLayout,
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

func (h *AdminHandlers) CreateMenuPost(w http.ResponseWriter, r *http.Request) {
	menuType := r.FormValue("type")

	if menuType == "" {
		h.log.Error("menu type is empty")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	err := h.menu.CreateMenu(menuType)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/admin/menu", http.StatusSeeOther)

}

func (h *AdminHandlers) DeleteMenu(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Delete menu page")
}
