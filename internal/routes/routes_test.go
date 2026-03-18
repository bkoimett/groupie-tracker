package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func test1(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
func test2(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
func TestRoutes(t *testing.T) {
	var testRoutes []Routes

	testRoutes = append(testRoutes, Routes{
		Handler: test1,
		Method:  http.MethodGet,
		Path:    "/",
	},
		Routes{
			Handler: test2,
			Method:  http.MethodGet,
			Path:    "/source",
		})

	mux := NewRoutes(testRoutes)

	tests := []struct {
		Path string
	}{
		{"/"},
		{"/source"},
	}

	for _, tt := range tests {
		req := httptest.NewRequest(http.MethodGet, tt.Path, nil)
		rec := httptest.NewRecorder()

		mux.ServeHTTP(rec, req)

		if rec.Code != http.StatusOK {
			t.Errorf("path %v failed, got %v", tt.Path, rec.Code)
		}
	}
}
