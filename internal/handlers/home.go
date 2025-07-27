package handlers

import (
	"html/template"
	"net/http"
)

func HomeHandler(tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, "Template execution error", http.StatusInternalServerError)
		}
	}
}
