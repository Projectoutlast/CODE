package admin

import (
	"fmt"
	"net/http"
)

func (h *AdminHandlers) Dishes(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Dishes page")
}

func (h *AdminHandlers) TheDish(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "The dish page")
}

func (h *AdminHandlers) EditDish(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Edit dish page")
}

func (h *AdminHandlers) CreateDish(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Create dish page")
}

func (h *AdminHandlers) DeleteDish(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Delete dish page")
}
