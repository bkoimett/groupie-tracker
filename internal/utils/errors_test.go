package utils

import (
	"html/template"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
)

// Normal render using the existing template
func TestErrorHandler_NormalRender(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()

	ErrorHandler(rr, req, 404, "Not Found")

	if rr.Code != 404 {
		t.Fatalf("expected 404 got %d", rr.Code)
	}

	body := rr.Body.String()

	if !strings.Contains(body, "404") {
		t.Errorf("response missing code")
	}

	if !strings.Contains(body, "Not Found") {
		t.Errorf("response missing message")
	}
}

// Missing template: simulate by providing a wrong path in a temporary handler
func TestErrorHandler_TemplateMissing(t *testing.T) {
	// Force a broken template
	original := errorTemplate
	errorTemplate = template.Must(template.New("bad").Parse("{{.NonExistent.Field}}"))
	defer func() { errorTemplate = original }()

	req := httptest.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()

	ErrorHandler(rr, req, 500, "Internal Server Error")

	if !strings.Contains(rr.Body.String(), "Internal Server Error") {
		t.Errorf("expected fallback error message")
	}
}

// Empty message
func TestErrorHandler_EmptyMessage(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()

	ErrorHandler(rr, req, 400, "")

	if rr.Code != 400 {
		t.Fatalf("expected 400 got %d", rr.Code)
	}

	body := rr.Body.String()
	if !strings.Contains(body, "400") {
		t.Errorf("expected body to contain code 400")
	}
}

// Long message
func TestErrorHandler_LongMessage(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()

	longMsg := strings.Repeat("error ", 1000)
	ErrorHandler(rr, req, 500, longMsg)

	if !strings.Contains(rr.Body.String(), "error") {
		t.Errorf("expected long message in response")
	}
}

// HTML escaping
func TestErrorHandler_HTMLEscaping(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()

	msg := "<script>alert(1)</script>"
	ErrorHandler(rr, req, 400, msg)

	body := rr.Body.String()
	if strings.Contains(body, "<script>") {
		t.Errorf("expected HTML to be escaped")
	}
}

// Invalid status code
func TestErrorHandler_InvalidStatusCode(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()

	ErrorHandler(rr, req, 999, "invalid")

	if rr.Code != 999 {
		t.Errorf("expected 999 got %d", rr.Code)
	}
}

// Status code zero
func TestErrorHandler_StatusCodeZero(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()

	ErrorHandler(rr, req, 0, "zero")

	if rr.Code == 0 {
		t.Errorf("response recorder should not keep status 0")
	}
}

// Concurrent requests
func TestErrorHandler_ConcurrentRequests(t *testing.T) {
	var wg sync.WaitGroup
	for range 50 {
		wg.Add(1)

		go func() {
			defer wg.Done()
			req := httptest.NewRequest("GET", "/", nil)
			rr := httptest.NewRecorder()

			ErrorHandler(rr, req, 500, "Concurrent")

			if rr.Code != 500 {
				t.Errorf("expected 500 got %d", rr.Code)
			}
		}()
	}

	wg.Wait()
}
