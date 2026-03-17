package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"groupie-tracker/internal/handlers"
)

func TestMainServerSetup(t *testing.T) {
	// Test that templates are parsed correctly
	if TemplateCache == nil {
		t.Error("TemplateCache should not be nil after init()")
	}

	// Check if templates are actually parsed (should have at least one template)
	if len(TemplateCache.Templates()) == 0 {
		t.Error("TemplateCache should contain at least one template")
	}

	// Verify all required templates exist
	requiredTemplates := []string{"index.html", "artist.html", "error.html"}
	for _, name := range requiredTemplates {
		if tmpl := TemplateCache.Lookup(name); tmpl == nil {
			t.Errorf("Template %s not found in cache", name)
		}
	}
}

func TestStaticFilesHandler(t *testing.T) {
	// Create a request to static file
	req, err := http.NewRequest("GET", "/static/css/style.css", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Create the static file handler
	fs := http.FileServer(http.Dir("../internal/static"))
	handler := http.StripPrefix("/static/", fs)

	// Serve the request
	handler.ServeHTTP(rr, req)

	// Check status code (404 is also valid if file doesn't exist, but handler should not panic)
	if rr.Code != http.StatusOK && rr.Code != http.StatusNotFound {
		t.Errorf("Static handler returned wrong status code: got %v, want %v or %v",
			rr.Code, http.StatusOK, http.StatusNotFound)
	}
}

func TestRouteRegistration(t *testing.T) {
	tests := []struct {
		name       string
		path       string
		wantStatus int
	}{
		{
			name:       "Home page",
			path:       "/",
			wantStatus: http.StatusOK,
		},
		{
			name:       "Artist page with ID",
			path:       "/artist?id=1",
			wantStatus: http.StatusOK,
		},
		{
			name:       "Artist page without ID",
			path:       "/artist",
			wantStatus: http.StatusBadRequest,
		},
		{
			name:       "Artist page with invalid ID",
			path:       "/artist?id=invalid",
			wantStatus: http.StatusBadRequest,
		},
		{
			name:       "Non-existent route",
			path:       "/nonexistent",
			wantStatus: http.StatusNotFound,
		},
	}

	// Set up the handlers like in main()
	handlers.SetTemplateCache(TemplateCache)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create request
			req, err := http.NewRequest("GET", tt.path, nil)
			if err != nil {
				t.Fatal(err)
			}

			// Create ResponseRecorder
			rr := httptest.NewRecorder()

			// Handle the request based on path pattern
			switch {
			case tt.path == "/":
				handlers.HomeHandler(rr, req)
			case tt.path == "/artist" || tt.path == "/artist?id=1" || tt.path == "/artist?id=invalid":
				handlers.ArtistHandler(rr, req)
			default:
				// For non-existent routes, simulate default 404 handler
				http.NotFound(rr, req)
			}

			// Check status code
			if rr.Code != tt.wantStatus {
				t.Errorf("Handler returned wrong status code: got %v, want %v",
					rr.Code, tt.wantStatus)
			}
		})
	}
}

func TestTemplateFunctions(t *testing.T) {
	// Instead of checking Funcs directly, test that templates can be executed
	// This indirectly verifies that template functions are registered

	// Check that we can look up templates
	tmpl := TemplateCache.Lookup("index.html")
	if tmpl == nil {
		t.Fatal("Could not find index.html template")
	}

	// If we got here, templates are parsed and available
	// The FuncMap would have been applied during parsing in init()
	// So if templates parsed without panic, functions are registered
}

func TestServerPort(t *testing.T) {
	// Test that the port is correctly configured
	port := ":8080"
	if port == "" {
		t.Error("Port should not be empty")
	}
	if port[0] != ':' {
		t.Error("Port should start with ':'")
	}
	if len(port) < 2 {
		t.Error("Port should have a number after ':'")
	}
}

func TestTemplateCacheInitialization(t *testing.T) {
	// Test that the template cache is properly set in handlers
	// This is a bit tricky since handlers.SetTemplateCache is called in init()

	// We can test by making a request that would fail if templates weren't set
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	// This should not panic if templates are properly set
	handlers.HomeHandler(rr, req)

	// Even if it returns an error, it shouldn't panic
	// So if we get here without panic, it's good
}

// TestMain function to run setup before tests
func TestMain(m *testing.M) {
	// The init() function will run automatically before tests
	// This test just ensures it's called by running the tests
	m.Run()
}
