package handlers

import (
	"html/template"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	templ, err := template.ParseFiles(
		"internal/templates/base.html",
		"internal/templates/home.html",
	)

	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}

	templ.Execute(w, nil)
}
