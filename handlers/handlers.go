package handlers

import (
    "encoding/json"
    "html/template"
    "net/http"
	
    "groupie-tracker/models"
)

// API endpoint
const apiURL = "https://groupietrackers.herokuapp.com/api/artists"

// HomeHandler fetches artists and displays them
func HomeHandler(w http.ResponseWriter, r *http.Request) {
    // Only handle root path
    if r.URL.Path != "/" {
        ErrorHandler(w, "Page not found", http.StatusNotFound)
        return
    }

    // Fetch artists from API
    response, err := http.Get(apiURL)
    if err != nil {
        ErrorHandler(w, "Failed to fetch artists", http.StatusInternalServerError)
        return
    }
    defer response.Body.Close()

    // Check if API returned success
    if response.StatusCode != http.StatusOK {
        ErrorHandler(w, "API returned error", http.StatusInternalServerError)
        return
    }

    // Parse JSON into our Artist structs
    var artists []models.Artist
    err = json.NewDecoder(response.Body).Decode(&artists)
    if err != nil {
        ErrorHandler(w, "Failed to parse artists data", http.StatusInternalServerError)
        return
    }

    // Parse and execute template
    tmpl, err := template.ParseFiles("templates/index.html")
    if err != nil {
        ErrorHandler(w, "Failed to load template", http.StatusInternalServerError)
        return
    }

    tmpl.Execute(w, artists)
}

// ErrorHandler displays error pages
func ErrorHandler(w http.ResponseWriter, message string, status int) {
    w.WriteHeader(status)
    tmpl, err := template.ParseFiles("templates/error.html")
    if err != nil {
        http.Error(w, message, status)
        return
    }
    
    data := struct {
        Status  int
        Message string
    }{
        Status:  status,
        Message: message,
    }
    
    tmpl.Execute(w, data)
}