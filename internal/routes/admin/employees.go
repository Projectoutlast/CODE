package admin

import (
	"code/internal/models"
	"html/template"
	"net/http"
)

func (h *AdminHandlers) Employees(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/html/admin/employee.page.html",
		baseHTMLLayout,
	}

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		h.log.Error(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	employees, err := h.employee.GetAllUsers()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if err = tmpl.Execute(w, employees); err != nil {
		h.log.Error(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func (h *AdminHandlers) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/html/admin/employee_create.html",
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

func (h *AdminHandlers) CreateEmployeeProcess(w http.ResponseWriter, r *http.Request) {
	login := r.FormValue("login")
	password := r.FormValue("password")
	email := r.FormValue("email")
	firstname := r.FormValue("firstname")
	lastname := r.FormValue("lastname")

	if login == "" || password == "" || email == "" || firstname == "" || lastname == "" {
		h.log.Warn("некорректное заполнение формы")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	newEmployee := &models.User{
		Login:     login,
		PasswordHash:  password,
		Email:     email,
		Firstname: firstname,
		Lastname:  lastname,
	}

	if err := h.employee.RegisterUser(newEmployee); err != nil {
		h.log.Error(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/admin/employees", http.StatusSeeOther)

}