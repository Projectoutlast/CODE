package site

import (
	"html/template"
	"net/http"
)

func (h *MainHandlers) MenuForCatering(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/html/cateringMenu.page.html",
		"./ui/html/base.layout.html",
	}

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		h.log.Warn("failed to parse template", "err", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if err = tmpl.Execute(w, nil); err != nil {
		h.log.Warn("failed to execute template", "err", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
