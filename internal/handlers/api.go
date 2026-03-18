package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"groupie-tracker/internal/models"
)

var baseURL = "https://groupietrackers.herokuapp.com/api"

// fetch is a reusable helper that requests an endpoint
// and decodes the response into the target structure.
func fetch(endpoint string, target any) error {
	resp, err := http.Get(baseURL + endpoint)
	if err != nil {
		return fmt.Errorf("failed to connect to API: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("API returned status: %s", resp.Status)
	}

	return json.NewDecoder(resp.Body).Decode(target)
}

// FetchArtists retrieves all artists
func GetArtists() ([]models.Artist, error) {
	var artists []models.Artist
	if err := fetch("/artists", &artists); err != nil {
		return nil, err
	}
	return artists, nil
}

// FetchArtistLocations - Updated to handle the direct object
func FetchArtistLocations(id int) ([]string, error) {
	var entry models.LocationEntry // NOT the wrapper
	if err := fetch(fmt.Sprintf("/locations/%d", id), &entry); err != nil {
		return nil, err
	}
	return entry.Locations, nil
}

// FetchArtistDates - Updated
func FetchArtistDates(id int) ([]string, error) {
	var entry models.DateEntry
	if err := fetch(fmt.Sprintf("/dates/%d", id), &entry); err != nil {
		return nil, err
	}
	return entry.Dates, nil
}

// FetchArtistRelations - Updated
func FetchArtistRelations(id int) (map[string][]string, error) {
	var entry models.RelationEntry
	if err := fetch(fmt.Sprintf("/relation/%d", id), &entry); err != nil {
		return nil, err
	}
	return entry.DatesLocations, nil
}
