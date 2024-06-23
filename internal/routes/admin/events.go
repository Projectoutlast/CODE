package admin

import (
	"fmt"
	"net/http"
)

func (h *AdminHandlers) Events(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Events page")
}

func (h *AdminHandlers) TheEvent(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "The event page")
}

func (h *AdminHandlers) EditEvent(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Edit event page")
}

func (h *AdminHandlers) CreateEvent(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Create event page")
}

func (h *AdminHandlers) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Delete event page")
}
