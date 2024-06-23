package admin

import (
	"fmt"
	"net/http"
)

func (h *AdminHandlers) Employees(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Employees page")
}

func (h *AdminHandlers) RegisterNewEmployee(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Register new employee page")
}

func (h *AdminHandlers) TheEmployee(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "The employee page")
}

func (h *AdminHandlers) EditEmployee(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Edit employee page")
}

func (h *AdminHandlers) DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Delete employee page")
}
