package handler

import (
    "net/http"
    "groupie-tracker/internal/handlers"
)

func Handler(w http.ResponseWriter, r *http.Request) {
    switch r.URL.Path {
    case "/":
        handlers.HomeHandler(w, r)
    case "/artist":
        handlers.ArtistHandler(w, r)
    default:
        http.NotFound(w, r)
    }
}