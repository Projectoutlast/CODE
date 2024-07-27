package admin

import (
	"net/http"
	"text/template"

	"code/internal/repository/models"
)

func (h *AdminHandlers) GetAllDishes(w http.ResponseWriter, r *http.Request) {
	dishes, err := h.dish.GetAllDishes()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	files := []string{
		"./ui/html/admin/dishes.page.html",
		baseHTMLLayout,
	}

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		h.log.Error(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if err = tmpl.Execute(w, dishes); err != nil {
		h.log.Error(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func (h *AdminHandlers) CreateNewDish(w http.ResponseWriter, r *http.Request) {
	menu_types, err := h.menu.GetAllMenuTypes()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return 
	}

	categories, err := h.category.GetAllCategories()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	dataForHTML := &models.CreateDish{
		MenuTypes: *menu_types,
		Categories: categories,
	}

	files := []string{
		"./ui/html/admin/dish_create.html",
		baseHTMLLayout,
	}

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		h.log.Error(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, dataForHTML); err != nil {
		h.log.Error(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

}
