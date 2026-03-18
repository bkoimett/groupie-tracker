package main

import (
	"fmt"
	"groupie-tracker/internal/handlers"
	"groupie-tracker/internal/routes"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

// TemplateCache stores parsed templates for reuse
var TemplateCache *template.Template

func init() {
    // Try multiple possible paths
    possiblePaths := []string{
        // When running from project root (Render)
        "internal/templates/*.html",
        // When running from cmd/ (local development)
        "../internal/templates/*.html",
        // Absolute path from executable
        filepath.Join(getExecutableDir(), "internal", "templates", "*.html"),
    }

    var templates *template.Template
    var err error
    

    // Try each path until one works
    for _, path := range possiblePaths {
        templates, err = template.New("").ParseGlob(path)
        if err == nil {
            TemplateCache = templates
            handlers.SetTemplateCache(TemplateCache)
            fmt.Printf("✅ Templates loaded from: %s\n", path)
            return
        }
    }
    
    // If we get here, none of the paths worked
    log.Fatal("Failed to load templates from any path")
}

// Helper function to get the directory of the executable
func getExecutableDir() string {
    exe, err := os.Executable()
    if err != nil {
        return ""
    }
    return filepath.Dir(exe)
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
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // default for local environment
	}

	fmt.Printf("🚀 Server starting on http://localhost%s\n", port)
	fmt.Println("📦 Press Ctrl+C to stop")
	for _, t := range router {
		fmt.Printf("%v	%v\n", t.Method, t.Path)
	}

	err := http.ListenAndServe(":"+port, r)
	if err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
