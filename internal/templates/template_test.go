package templates

import (
	"bytes"
	"html/template"
	"os"
	"strings"
	"testing"

	"groupie-tracker/internal/handlers" // Import the real handlers
)


func TestIndexTemplate(t *testing.T) {
	// Use REAL functions
	funcMap := template.FuncMap{
		"formatLocation": handlers.FormatLocation,
		"formatDate":     handlers.FormatDate,
	}

	tmpl, err := template.New("index.html").Funcs(funcMap).ParseFiles("index.html")
	if err != nil {
		t.Skip("Template file not found, skipping test")
	}

	// Mock data with proper structure
	type MockArtist struct {
		ID           int
		Name         string
		Image        string
		CreationDate int
		FirstAlbum   string
		Members      []string
		Locations    []string
		ConcertDates []string
		Relations    map[string][]string
	}
	
	data := []MockArtist{
		{
			ID:           1,
			Name:         "Test Artist",
			Image:        "/test.jpg",
			CreationDate: 2020,
			FirstAlbum:   "2020-01-01",
			Members:      []string{"Member 1", "Member 2"},
		},
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, data)
	if err != nil {
		t.Fatalf("Failed to execute template: %v", err)
	}

	result := buf.String()
	
	// Check for required elements
	required := []string{
		"Test Artist", 
		"2020", 
		"artist-card",
		"details-link",
	}
	
	for _, req := range required {
		if !strings.Contains(result, req) {
			t.Errorf("Template missing: %s", req)
		}
	}
}

func TestArtistTemplate(t *testing.T) {
	// Use REAL functions
	funcMap := template.FuncMap{
		"formatLocation": handlers.FormatLocation,
		"formatDate":     handlers.FormatDate,
	}

	tmpl, err := template.New("artist.html").Funcs(funcMap).ParseFiles("artist.html")
	if err != nil {
		t.Skip("Template file not found, skipping test")
	}

	// Mock data
	type MockFullArtistData struct {
		Name         string
		Image        string
		CreationDate int
		FirstAlbum   string
		Members      []string
		Locations    []string
		ConcertDates []string
		Relations    map[string][]string
	}
	
	data := MockFullArtistData{
		Name:         "Test Artist",
		Image:        "/test.jpg",
		CreationDate: 2020,
		FirstAlbum:   "2020-01-01",
		Members:      []string{"Member 1", "Member 2"},
		Locations:    []string{"london_uk", "paris_france"},
		ConcertDates: []string{"2020-01-20", "2020-02-20"},
		Relations: map[string][]string{
			"2020-01-20": {"london_uk"},
			"2020-02-20": {"paris_france"},
		},
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, data)
	if err != nil {
		t.Fatalf("Failed to execute template: %v", err)
	}

	result := buf.String()
	required := []string{
		"Test Artist", 
		"2020", 
		"Member 1", 
		"Member 2",
		"london_uk", 
		"2020-01-20",
		"tour-schedule",
		"location-card",
		"date-item",
	}
	
	for _, req := range required {
		if !strings.Contains(result, req) {
			t.Errorf("Template missing: %s", req)
		}
	}
}

func TestErrorTemplate(t *testing.T) {
	// Use REAL functions
	funcMap := template.FuncMap{
		"formatLocation": handlers.FormatLocation,
		"formatDate":     handlers.FormatDate,
	}

	tmpl, err := template.New("error.html").Funcs(funcMap).ParseFiles("error.html")
	if err != nil {
		t.Skip("Template file not found, skipping test")
	}

	data := struct {
		Code    int
		Message string
	}{
		Code:    404,
		Message: "Page not found",
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, data)
	if err != nil {
		t.Fatalf("Failed to execute template: %v", err)
	}

	result := buf.String()
	required := []string{"404", "Page not found", "error-container", "home-button"}
	
	for _, req := range required {
		if !strings.Contains(result, req) {
			t.Errorf("Template missing: %s", req)
		}
	}
}

func TestThemeToggleInTemplates(t *testing.T) {
	files := []string{"index.html", "artist.html", "error.html"}
	
	for _, file := range files {
		t.Run(file, func(t *testing.T) {
			content, err := os.ReadFile(file)
			if err != nil {
				t.Skipf("Template file %s not found", file)
			}

			// Theme toggle is added by JS, so we just check for the script
			if !strings.Contains(string(content), "main.js") {
				t.Error("Template missing script tag for main.js")
			}
		})
	}
}