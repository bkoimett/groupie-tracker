package handler

import (
    "net/http"
    "groupie-tracker/internal/handlers" // Adjust module name
)

// Handler is the entry point for Vercel
func Handler(w http.ResponseWriter, r *http.Request) {
    // Route to your existing handlers
    switch r.URL.Path {
    case "/":
        handlers.HomeHandler(w, r)
    case "/artist":
        handlers.ArtistHandler(w, r)
    default:
        // Handle static files or 404
        http.NotFound(w, r)
    }
}