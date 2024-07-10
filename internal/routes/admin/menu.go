package admin

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (h *AdminHandlers) Menu(w http.ResponseWriter, r *http.Request) {
	menuTypes, err := h.menu.GetAllMenuTypes()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	response := MenuResponse{
		MenuTypes: *menuTypes,
	}

	files := []string{
		"./ui/html/admin/menu.page.html",
		baseHTMLLayout,
	}

	tmpl := template.Must(template.ParseFiles(files...))

	if err = tmpl.Execute(w, response); err != nil {
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

func (h *AdminHandlers) EditMenuGet(w http.ResponseWriter, r *http.Request) {
	type_id, err := strconv.Atoi(mux.Vars(r)["type_id"])

	if err != nil {
		h.log.Error(fmt.Sprintf("can't parse menu type id: %s", err.Error()))
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	menuType, err := h.menu.GetMenuType(type_id)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

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

	if err = tmpl.Execute(w, menuType); err != nil {
		h.log.Error(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func (h *AdminHandlers) EditMenuPost(w http.ResponseWriter, r *http.Request) {
	type_id, err := strconv.Atoi(mux.Vars(r)["type_id"])
	if err != nil {
		h.log.Error(fmt.Sprintf("can't parse menu type id: %s", err.Error()))
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	type_menu := r.FormValue("type")

	if type_menu == "" {
		h.log.Error("menu type is empty")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if err := h.menu.UpdateMenu(type_id, type_menu); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/admin/menu", http.StatusSeeOther)
}

func (h *AdminHandlers) DeleteMenu(w http.ResponseWriter, r *http.Request) {
	type_id, err := strconv.Atoi(mux.Vars(r)["type_id"])

	if err != nil {
		h.log.Error(fmt.Sprintf("can't parse menu type id: %s", err.Error()))
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if err := h.menu.DeleteMenu(type_id); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/admin/menu", http.StatusSeeOther)
}
