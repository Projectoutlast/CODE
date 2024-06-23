package admin

import (
	"fmt"
	"net/http"
)

func (h *AdminHandlers) Menu(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Menu page")
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
