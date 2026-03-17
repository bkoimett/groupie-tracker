package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"groupie-tracker/internal/models"
)

// helper to start a mock server and override baseURL
func mockFetch(t *testing.T, response interface{}) (*httptest.Server, func()) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}))

	oldBaseURL := baseURL
	baseURL = server.URL
	return server, func() {
		server.Close()
		baseURL = oldBaseURL
	}
}

func TestFetchArtists(t *testing.T) {
	mockData := []models.Artist{
		{ID: 1, Name: "Queen"},
		{ID: 2, Name: "Adele"},
	}

	_, cleanup := mockFetch(t, mockData)
	defer cleanup()

	artists, err := GetArtists()
	if err != nil {
		t.Fatalf("FetchArtists failed: %v", err)
	}
	if len(artists) != 2 {
		t.Errorf("expected 2 artists, got %d", len(artists))
	}
}

func TestFetchArtistLocations(t *testing.T) {
	// FIX: Pass the Entry directly, NOT the wrapper with Index
	mockData := models.LocationEntry{
		ID:        1,
		Locations: []string{"london-uk", "paris-france"},
	}

	_, cleanup := mockFetch(t, mockData)
	defer cleanup()

	locs, err := FetchArtistLocations(1)
	if err != nil {
		t.Fatalf("FetchArtistLocations failed: %v", err)
	}
	if len(locs) != 2 || locs[0] != "london-uk" {
		t.Errorf("unexpected locations: %v", locs)
	}
}

func TestFetchArtistDates(t *testing.T) {
	// FIX: Pass the Entry directly
	mockData := models.DateEntry{
		ID:    1,
		Dates: []string{"*20-03-2026", "*15-04-2026"},
	}

	_, cleanup := mockFetch(t, mockData)
	defer cleanup()

	dates, err := FetchArtistDates(1)
	if err != nil {
		t.Fatalf("FetchArtistDates failed: %v", err)
	}
	if len(dates) != 2 || dates[0] != "*20-03-2026" {
		t.Errorf("unexpected dates: %v", dates)
	}
}

func TestFetchArtistRelations(t *testing.T) {
	// FIX: Pass the Entry directly
	mockData := models.RelationEntry{
		ID: 1,
		DatesLocations: map[string][]string{
			"20-03-2026": {"london-uk"},
		},
	}

	_, cleanup := mockFetch(t, mockData)
	defer cleanup()

	relations, err := FetchArtistRelations(1)
	if err != nil {
		t.Fatalf("FetchArtistRelations failed: %v", err)
	}
	if locs, ok := relations["20-03-2026"]; !ok || locs[0] != "london-uk" {
		t.Errorf("unexpected relations: %v", relations)
	}
}

func TestNotFound(t *testing.T) {
	// To simulate a "Not Found", we make the mock server return a 404 status
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))

	oldBaseURL := baseURL
	baseURL = server.URL
	defer func() {
		server.Close()
		baseURL = oldBaseURL
	}()

	_, err := FetchArtistLocations(999)
	if err == nil {
		t.Fatal("expected error for 404 status, got nil")
	}
}
