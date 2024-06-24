package site

import (
	"html/template"
	"net/http"
)

func (h *MainHandlers) NewsAndEvents(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/html/site/newsAndEvents.page.html",
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