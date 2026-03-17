package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"groupie-tracker/internal/models"
)

var baseURL = "https://groupietrackers.herokuapp.com/api"

// fetch is a reusable helper that requests an endpoint
// and decodes the response into the target structure.
func fetch(endpoint string, target interface{}) error {
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

// FormatLocation cleans up location strings
// FormatLocation cleans up location strings
func FormatLocation(location string) string {
    if location == "" {
        return ""
    }
    
    // Replace underscores with spaces
    location = strings.ReplaceAll(location, "_", " ")
    
    // Handle hyphens - replace first hyphen with ", " and rest with spaces
    parts := strings.SplitN(location, "-", 2)
    if len(parts) == 2 {
        location = parts[0] + ", " + parts[1]
        // Replace any remaining hyphens with spaces
        location = strings.ReplaceAll(location, "-", " ")
    }
    
    // Capitalize words
    words := strings.Fields(location)
    for i, word := range words {
        if len(word) > 0 {
            words[i] = strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])
        }
    }
    return strings.Join(words, " ")
}

// FormatDate makes dates more readable
func FormatDate(date string) string {
	// Remove asterisks if present
	date = strings.TrimPrefix(date, "*")

	// Assuming date comes as "YYYY-MM-DD"
	parts := strings.Split(date, "-")
	if len(parts) == 3 {
		months := map[string]string{
			"01": "January", "02": "February", "03": "March",
			"04": "April", "05": "May", "06": "June",
			"07": "July", "08": "August", "09": "September",
			"10": "October", "11": "November", "12": "December",
		}
		month := months[parts[1]]
		return month + " " + parts[2] + ", " + parts[0]
	}
	return date
}
