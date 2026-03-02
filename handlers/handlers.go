package handlers

import (
    // "encoding/json"
    "html/template"
    "net/http"
    "strconv"
    // "strings"
    "groupie-tracker/models"
)

// Template cache
var templates *template.Template

// SetTemplateCache receives the parsed templates from main
func SetTemplateCache(tmpl *template.Template) {
    templates = tmpl
}

const baseURL = "https://groupietrackers.herokuapp.com/api"

// HomeHandler fetches artists and displays them
func HomeHandler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        ErrorHandler(w, "Page not found", http.StatusNotFound)
        return
    }

    artists, err := FetchArtists()
    if err != nil {
        ErrorHandler(w, "Failed to fetch artists: "+err.Error(), http.StatusInternalServerError)
        return
    }

    // Use cached template
    err = templates.ExecuteTemplate(w, "index.html", artists)
    if err != nil {
        ErrorHandler(w, "Failed to render template: "+err.Error(), http.StatusInternalServerError)
        return
    }
}

// ArtistHandler displays detailed information for a single artist
func ArtistHandler(w http.ResponseWriter, r *http.Request) {
    // Get artist ID from URL query
    idStr := r.URL.Query().Get("id")
    if idStr == "" {
        ErrorHandler(w, "Artist ID not provided", http.StatusBadRequest)
        return
    }

    id, err := strconv.Atoi(idStr)
    if err != nil || id < 1 {
        ErrorHandler(w, "Invalid artist ID", http.StatusBadRequest)
        return
    }

    // Fetch all artists to get basic info
    artists, err := FetchArtists()
    if err != nil {
        ErrorHandler(w, "Failed to fetch artist data: "+err.Error(), http.StatusInternalServerError)
        return
    }

    // Find the specific artist
    var artist models.Artist
    found := false
    for _, a := range artists {
        if a.ID == id {
            artist = a
            found = true
            break
        }
    }

    if !found {
        ErrorHandler(w, "Artist not found", http.StatusNotFound)
        return
    }

    // Fetch additional data
    locations, err := FetchArtistLocations(id)
    if err != nil {
        // Log but continue with empty locations
        artist.Locations = []string{}
    } else {
        artist.Locations = locations
    }

    dates, err := FetchArtistDates(id)
    if err != nil {
        artist.ConcertDates = []string{}
    } else {
        artist.ConcertDates = dates
    }

    relations, err := FetchArtistRelations(id)
    if err != nil {
        artist.Relations = map[string][]string{}
    } else {
        artist.Relations = relations
    }

    // Use cached template
    err = templates.ExecuteTemplate(w, "artist.html", artist)
    if err != nil {
        ErrorHandler(w, "Failed to render template: "+err.Error(), http.StatusInternalServerError)
        return
    }
}

// ErrorHandler displays error pages
func ErrorHandler(w http.ResponseWriter, message string, status int) {
    w.WriteHeader(status)
    
    // Try to use cached template
    if templates != nil {
        data := struct {
            Status  int
            Message string
        }{
            Status:  status,
            Message: message,
        }
        
        err := templates.ExecuteTemplate(w, "error.html", data)
        if err == nil {
            return
        }
    }
    
    // Fallback to simple error message
    http.Error(w, message, status)
}