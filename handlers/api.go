package handlers

import (
    "encoding/json"
    "fmt"
    "net/http"
    // "strconv"
    "strings"
    "groupie-tracker/models"
)

// const baseURL = "https://groupietrackers.herokuapp.com/api"

// FetchArtists gets all artists from the API
func FetchArtists() ([]models.Artist, error) {
    resp, err := http.Get(baseURL + "/artists")
    if err != nil {
        return nil, fmt.Errorf("failed to connect to API: %v", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("API returned status: %s", resp.Status)
    }

    var artists []models.Artist
    err = json.NewDecoder(resp.Body).Decode(&artists)
    if err != nil {
        return nil, fmt.Errorf("failed to parse API response: %v", err)
    }
    
    return artists, nil
}

// FetchArtistLocations gets locations for a specific artist
func FetchArtistLocations(artistID int) ([]string, error) {
    url := fmt.Sprintf("%s/locations/%d", baseURL, artistID)
    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("API returned status: %s", resp.Status)
    }

    var locationData struct {
        Locations []string `json:"locations"`
    }
    err = json.NewDecoder(resp.Body).Decode(&locationData)
    return locationData.Locations, err
}

// FetchArtistDates gets concert dates for a specific artist
func FetchArtistDates(artistID int) ([]string, error) {
    url := fmt.Sprintf("%s/dates/%d", baseURL, artistID)
    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("API returned status: %s", resp.Status)
    }

    var dateData struct {
        Dates []string `json:"dates"`
    }
    err = json.NewDecoder(resp.Body).Decode(&dateData)
    return dateData.Dates, err
}

// FetchArtistRelations gets the relations (dates mapped to locations)
func FetchArtistRelations(artistID int) (map[string][]string, error) {
    url := fmt.Sprintf("%s/relation/%d", baseURL, artistID)
    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("API returned status: %s", resp.Status)
    }

    var relationData struct {
        DatesLocations map[string][]string `json:"datesLocations"`
    }
    err = json.NewDecoder(resp.Body).Decode(&relationData)
    return relationData.DatesLocations, err
}

// FormatLocation cleans up location strings
func FormatLocation(location string) string {
    // Replace underscores with spaces and format nicely
    location = strings.ReplaceAll(location, "_", " ")
    location = strings.ReplaceAll(location, "-", ", ")
    
    // Capitalize first letter of each word
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