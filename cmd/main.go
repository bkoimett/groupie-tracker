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
	// Parse all templates with functions
	TemplateCache = template.Must(template.New("").ParseGlob("../internal/templates/*.html"))

	// Make templates available to handlers
	handlers.SetTemplateCache(TemplateCache)
}

func main() {
	var router []routes.Routes
	router = append(router,
		routes.Routes{
			Method:  http.MethodGet,
			Path:    "/",
			Handler: handlers.HomeHandler,
		},
		routes.Routes{
			Method:  http.MethodGet,
			Path:    "/artist",
			Handler: handlers.ArtistHandler,
		},
	)

	r := routes.NewRoutes(router)

	// Start server
	port := ":8080"
	fmt.Printf("🚀 Server starting on http://localhost%s\n", port)
	fmt.Println("📦 Press Ctrl+C to stop")
	for _, t := range router {
		fmt.Printf("%v	%v\n", t.Method, t.Path)
	}

	err := http.ListenAndServe(port, r)
	if err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
