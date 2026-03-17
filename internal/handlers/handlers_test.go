package handlers

import (
	"html/template"
	"net/http"
	"net/http/httptest"
	"testing"
)

func init() {
	// Minimal templates so ExecuteTemplate doesn't panic
	templates = template.Must(template.New("index.html").Parse(`<html>{{range .}}{{.ID}}{{end}}</html>`))
	template.Must(templates.New("artist.html").Parse(`<html>{{.ID}}</html>`))
}

func TestHomeHandler_ValidGet(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	HomeHandler(rec, req)
	if rec.Code != http.StatusOK {
		t.Errorf("expected 200 got %d", rec.Code)
	}
}

func TestHomeHandler_WrongMethod(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	rec := httptest.NewRecorder()
	HomeHandler(rec, req)
	if rec.Code != http.StatusMethodNotAllowed {
		t.Errorf("expected 405 got %d", rec.Code)
	}
}

func TestHomeHandler_WrongPath(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/unknown", nil)
	rec := httptest.NewRecorder()
	HomeHandler(rec, req)
	if rec.Code != http.StatusNotFound {
		t.Errorf("expected 404 got %d", rec.Code)
	}
}

func TestArtistHandler_ValidID(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/artist?id=1", nil)
	rec := httptest.NewRecorder()
	ArtistHandler(rec, req)
	if rec.Code != http.StatusOK {
		t.Errorf("expected 200 got %d", rec.Code)
	}
}

func TestArtistHandler_InvalidID_String(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/artist?id=invalid", nil)
	rec := httptest.NewRecorder()
	ArtistHandler(rec, req)
	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected 400 got %d", rec.Code)
	}
}

func TestArtistHandler_InvalidID_Negative(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/artist?id=-2", nil)
	rec := httptest.NewRecorder()
	ArtistHandler(rec, req)
	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected 400 got %d", rec.Code)
	}
}

func TestArtistHandler_InvalidID_Zero(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/artist?id=0", nil)
	rec := httptest.NewRecorder()
	ArtistHandler(rec, req)
	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected 400 got %d", rec.Code)
	}
}

func TestArtistHandler_MissingID(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/artist", nil)
	rec := httptest.NewRecorder()
	ArtistHandler(rec, req)
	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected 400 got %d", rec.Code)
	}
}

func TestArtistHandler_NotFound(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/artist?id=999999", nil)
	rec := httptest.NewRecorder()
	ArtistHandler(rec, req)
	if rec.Code != http.StatusNotFound {
		t.Errorf("expected 404 got %d", rec.Code)
	}
}

func TestArtistHandler_WrongPath(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/artist/1", nil)
	rec := httptest.NewRecorder()
	ArtistHandler(rec, req)
	if rec.Code != http.StatusNotFound {
		t.Errorf("expected 404 got %d", rec.Code)
	}
}
