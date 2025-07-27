package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/boiddutik/go_htmx/crud_app/internal/db"
	"github.com/boiddutik/go_htmx/crud_app/internal/routes"
)

func main() {
	db.OpenDB()
	defer db.DB.Close()
	log.Println("✅ Connected to SQLite database.")

	tmpl := template.Must(template.ParseGlob("ui/templates/*.html"))

	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("ui/static"))))
	routes.SetupRoutes(mux, tmpl)

	log.Println("🚀 Server running at http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", mux))
}
