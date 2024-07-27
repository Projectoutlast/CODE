package admin

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"

	"code/internal/models"

	"github.com/gorilla/mux"
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
		MenuTypes:  *menu_types,
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

func (h *AdminHandlers) ProcessCreateNewDish(w http.ResponseWriter, r *http.Request) {
	menuTypeID, errTypeID := strconv.Atoi(r.FormValue("menuType"))
	categoryID, errCategoryID := strconv.Atoi(r.FormValue("category"))
	price, errPrice := strconv.ParseFloat(r.FormValue("price"), 64)
	weight, errWeight := strconv.Atoi(r.FormValue("weight"))

	if errTypeID != nil || errCategoryID != nil || errPrice != nil || errWeight != nil {
		h.log.Error("некорректное значение одного из полей (меню, категория, цена, вес)")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	dishName := r.FormValue("name")
	composition := r.FormValue("composition")
	description := r.FormValue("description")

	tags := r.FormValue("tags")

	if menuTypeID == 0 ||
		categoryID == 0 ||
		dishName == "" ||
		composition == "" ||
		price == 0.0 ||
		weight == 0 {
		h.log.Warn("не заполнено одно или несколько полей: меню, категория, наименование блюда, состав, цена, вес.")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	newDish := &models.Dish{
		Name:           dishName,
		CategoryDishID: categoryID,
		MenuTypeID:     menuTypeID,
		Composition:    composition,
		Description:    description,
		Price:          price,
		Weight:         weight,
		Tags:           tags,
	}

	if err := h.dish.CreateNewDish(newDish); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/admin/menu/dish", http.StatusSeeOther)

}

func (h *AdminHandlers) DishEdit(w http.ResponseWriter, r *http.Request) {
	dishID, err := strconv.Atoi(mux.Vars(r)["dish_id"])
	if err != nil {
		h.log.Error(fmt.Sprintf("can't parse dish id: %s", err.Error()))
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	dish, err := h.dish.GetDish(dishID)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	files := []string{
		"./ui/html/admin/dish_edit.html",
		baseHTMLLayout,
	}

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		h.log.Error(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if err = tmpl.Execute(w, dish); err != nil {
		h.log.Error(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func (h *AdminHandlers) DishEditProcess(w http.ResponseWriter, r *http.Request) {
	dishID, err := strconv.Atoi(mux.Vars(r)["dish_id"])
	if err != nil {
		h.log.Error(fmt.Sprintf("can't parse dish id: %s", err.Error()))
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	price, errPrice := strconv.ParseFloat(r.FormValue("price"), 64)
	weight, errWeight := strconv.Atoi(r.FormValue("weight"))

	if errPrice != nil || errWeight != nil {
		h.log.Error("некорректное значение одного из полей (цена, вес)")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	dishName := r.FormValue("name")
	composition := r.FormValue("composition")
	description := r.FormValue("description")
	tags := r.FormValue("tags")

	if dishName == "" || composition == "" || price == 0.0 || weight == 0 {
		h.log.Warn("не заполнено одно или несколько полей: наименование блюда, состав, цена, вес.")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	dish := &models.Dish{
		ID:             dishID,
		Name:           dishName,
		Composition:    composition,
		Description:    description,
		Price:          price,
		Weight:         weight,
		Tags:           tags,
	}

	if err := h.dish.UpdateDish(dish); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/admin/menu/dish", http.StatusSeeOther)
}

func (h *AdminHandlers) DeleteDish(w http.ResponseWriter, r *http.Request) {
	dishID, err := strconv.Atoi(mux.Vars(r)["dish_id"])
	if err != nil {
		h.log.Error(fmt.Sprintf("can't parse dish id: %s", err.Error()))
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if err := h.dish.DeleteDish(dishID); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/admin/menu/dish", http.StatusSeeOther)
}