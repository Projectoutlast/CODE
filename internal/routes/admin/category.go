package admin

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"

	"github.com/gorilla/mux"
)

func (h *AdminHandlers) Categories(w http.ResponseWriter, r *http.Request) {
	categories, err := h.category.GetAllCategories()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	files := []string{
		"./ui/html/admin/category.page.html",
		baseHTMLLayout,
	}

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		h.log.Error(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if err = tmpl.Execute(w, categories); err != nil {
		h.log.Error(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func (h *AdminHandlers) CreateCategory(w http.ResponseWriter, r *http.Request) {
	categories, err := h.menu.GetAllMenuTypes()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	files := []string{
		"./ui/html/admin/category_create.html",
		baseHTMLLayout,
	}

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		h.log.Error(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if err = tmpl.Execute(w, categories); err != nil {
		h.log.Error(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func (h *AdminHandlers) CreateCategoryProcess(w http.ResponseWriter, r *http.Request) {
	menuTypeID := r.FormValue("menuType")
	category := r.FormValue("category")

	if menuTypeID == "" || category == "" {
		h.log.Warn("empty category or menu type")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if err := h.category.CreateCategory(menuTypeID, category); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/admin/menu/category", http.StatusSeeOther)
}

func (h *AdminHandlers) EditCategory(w http.ResponseWriter, r *http.Request) {
	category_id, err := strconv.Atoi(mux.Vars(r)["category_id"])

	if err != nil {
		h.log.Error(fmt.Sprintf("can't parse category id: %s", err.Error()))
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	category, err := h.category.GetCategory(category_id)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	files := []string{
		"./ui/html/admin/category_edit.page.html",
		baseHTMLLayout,
	}

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		h.log.Error(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if err = tmpl.Execute(w, category); err != nil {
		h.log.Error(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func (h *AdminHandlers) EditCategoryProcess(w http.ResponseWriter, r *http.Request) {
	category_id, err := strconv.Atoi(mux.Vars(r)["category_id"])

	if err != nil {
		h.log.Error(fmt.Sprintf("can't parse category id: %s", err.Error()))
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	category := r.FormValue("category")

	if category == "" {
		h.log.Error("menu type is empty")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if err := h.category.UpdateCategory(category_id, category); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/admin/menu/category", http.StatusSeeOther)
}

func (h *AdminHandlers) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	category_id, err := strconv.Atoi(mux.Vars(r)["category_id"])

	if err != nil {
		h.log.Error(fmt.Sprintf("can't parse category id: %s", err.Error()))
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if err := h.category.DeleteCategory(category_id); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/admin/menu/category", http.StatusSeeOther)
}
