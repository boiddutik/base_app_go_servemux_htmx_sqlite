package routes

import (
	"html/template"
	"net/http"

	"github.com/boiddutik/go_htmx/crud_app/internal/handlers"
)

func SetupRoutes(mux *http.ServeMux, tmpl *template.Template) {
	mux.HandleFunc("/", handlers.HomeHandler(tmpl))
}
