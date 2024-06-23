package site

import (
	"html/template"
	"net/http"
)

func (h *MainHandlers) PrivacyPolicy(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/html/site/privacyPolicy.page.html",
		baseHTMLLayout,
	}

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		h.log.Warn("failed to parse template", "err", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if err = tmpl.Execute(w, nil); err != nil {
		h.log.Warn("failed to execute template", "err", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func (h *MainHandlers) UserAgreement(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/html/site/userAgreement.page.html",
		baseHTMLLayout,
	}

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		h.log.Warn("failed to parse template", "err", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if err = tmpl.Execute(w, nil); err != nil {
		h.log.Warn("failed to execute template", "err", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
