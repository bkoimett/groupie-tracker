package handlers

import (
	"groupie-tracker/internal/utils"
	"html/template"
	"net/http"
	"strconv"
	"sync"

	"groupie-tracker/internal/models"
)

// Template cache
var templates *template.Template

// SetTemplateCache receives the parsed templates from main
func SetTemplateCache(tmpl *template.Template) {
	templates = tmpl
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		utils.ErrorHandler(w, r, http.StatusNotFound, "page not found")
		return
	}
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	artists, err := GetArtists()
	if err != nil {
		utils.ErrorHandler(w, r, http.StatusInternalServerError, "failed to get artists")
		return
	}

	err = templates.ExecuteTemplate(w, "index.html", artists)
	if err != nil {
		utils.ErrorHandler(w, r, http.StatusInternalServerError, "Failed to render template: "+err.Error())
		return
	}
}

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/artist" {
		utils.ErrorHandler(w, r, http.StatusNotFound, "page not found")
		return
	}
	id := r.URL.Query().Get("id")
	artistID, err := strconv.Atoi(id)
	if err != nil || artistID < 1 {
		utils.ErrorHandler(w, r, http.StatusBadRequest, "invalid artist id")
		return
	}

	artists, err := GetArtists()
	if err != nil {
		utils.ErrorHandler(w, r, http.StatusInternalServerError, "failed to get artists")
		return
	}

	var locations, dates []string
	var relations map[string][]string
	var locErr, dateErr, relErr error

	// wait group - sync
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		defer wg.Done()
		locations, locErr = FetchArtistLocations(artistID)
	}()
	go func() {
		defer wg.Done()
		dates, dateErr = FetchArtistDates(artistID)
	}()
	go func() {
		defer wg.Done()
		relations, relErr = FetchArtistRelations(artistID)
	}()
	wg.Wait()

	if locErr != nil {
		utils.ErrorHandler(w, r, http.StatusInternalServerError, "Failed to get locations: "+err.Error())
		return
	}
	if dateErr != nil {
		utils.ErrorHandler(w, r, http.StatusInternalServerError, "Failed to get dates: "+err.Error())
		return
	}
	if relErr != nil {
		utils.ErrorHandler(w, r, http.StatusInternalServerError, "Failed to get relations: "+err.Error())
		return
	}

	var artist models.FullArtistData

	found := false
	for _, a := range artists {
		if a.ID == artistID {
			artist = models.FullArtistData{
				Artist:       a,
				Locations:    locations,
				ConcertDates: dates,
				Relations:    relations,
			}
			found = true
			break
		}
	}
	if !found {
		utils.ErrorHandler(w, r, http.StatusNotFound, "Artist not found")
		return
	}

	// Use cached template
	err = templates.ExecuteTemplate(w, "artist.html", artist)
	if err != nil {
		utils.ErrorHandler(w, r, http.StatusInternalServerError, "Failed to render template: "+err.Error())
		return
	}
}
