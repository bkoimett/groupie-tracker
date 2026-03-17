package main

import (
	"fmt"
	"groupie-tracker/internal/handlers"
	"groupie-tracker/internal/routes"
	"html/template"
	"log"
	"net/http"
)

// TemplateCache stores parsed templates for reuse
var TemplateCache *template.Template

func init() {
	// Register template functions
	funcMap := template.FuncMap{
		"formatLocation": handlers.FormatLocation,
		"formatDate":     handlers.FormatDate,
	}

	// Parse all templates with functions
	TemplateCache = template.Must(template.New("").Funcs(funcMap).ParseGlob("../internal/templates/*.html"))

	// Make templates available to handlers
	handlers.SetTemplateCache(TemplateCache)
}

func main() {
	var trees []routes.Tree
	trees = append(trees, routes.Tree{
		Handler: handlers.HomeHandler,
		Method: http.MethodGet,
		Path: "/",
	},
	routes.Tree{
		Handler: handlers.ArtistHandler,
		Method: http.MethodGet,
		Path: "/artist",
	})

	r, methods := routes.Routes(trees)

	// Start server
	port := ":8080"
	fmt.Printf("🚀 Server starting on http://localhost%s\n", port)
	fmt.Println("📦 Press Ctrl+C to stop")
	fmt.Println("✨ New today: Artist details pages with locations and dates!")
	for _, m := range methods {
		fmt.Printf("%v	%v\n", m.Method, m.Path)
	}

	err := http.ListenAndServe(port, r)
	if err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
