package main

import (
    "fmt"
    "html/template"
    "log"
    "net/http"
    "groupie-tracker/handlers"
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
    TemplateCache = template.Must(template.New("").Funcs(funcMap).ParseGlob("templates/*.html"))
    
    // Make templates available to handlers
    handlers.SetTemplateCache(TemplateCache)
}

func main() {
    // Serve static files
    fs := http.FileServer(http.Dir("static"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

    // Register handlers
    http.HandleFunc("/", handlers.HomeHandler)
    http.HandleFunc("/artist", handlers.ArtistHandler)

    // Start server
    port := ":8080"
    fmt.Printf("🚀 Server starting on http://localhost%s\n", port)
    fmt.Println("📦 Press Ctrl+C to stop")
    fmt.Println("✨ New today: Artist details pages with locations and dates!")
    
    err := http.ListenAndServe(port, nil)
    if err != nil {
        log.Fatal("Server failed to start:", err)
    }
}