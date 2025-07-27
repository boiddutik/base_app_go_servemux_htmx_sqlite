package main

import (
	"log"
	"net/http"
	"text/template"

	"github.com/boiddutik/go_htmx/crud_app/internal/db"
	"github.com/boiddutik/go_htmx/crud_app/internal/routes"
)

var tmpl *template.Template

func init() {
	db.OpenDB()
	tmpl = template.Must(template.ParseGlob("ui/templates/*.html"))
	log.Println("✅ DB initialized and templates parsed")
}

func main() {
	defer db.DB.Close()

	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("ui/static"))))

	routes.SetupRoutes(mux, tmpl)

	log.Println("🚀 Server running at http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", mux))
}
