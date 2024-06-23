package admin

import (
	"fmt"
	"net/http"
)

func (h *AdminHandlers) Categories(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Categories page")
}

func (h *AdminHandlers) TheCategory(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "The category page")
}

func (h *AdminHandlers) EditCategory(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Edit category page")
}

func (h *AdminHandlers) CreateCategory(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Create category page")
}

func (h *AdminHandlers) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Delete category page")
}
