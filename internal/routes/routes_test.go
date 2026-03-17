package routes

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func test1(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
func test2(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
func TestRoutes(t *testing.T) {
	var testTree []Tree

	testTree = append(testTree, Tree{
		Handler: test1,
		Method: http.MethodGet,
		Path: "/",
	},
	Tree{
		Handler: test2,
		Method: http.MethodGet,
		Path: "/source",
	})

	mux, routes := Routes(testTree)

	expectedTree := []Tree{
		{Method: http.MethodGet, Path: "/"},
		{Method: http.MethodGet, Path: "/source?id={id}"},
	}

	if !reflect.DeepEqual(expectedTree, routes) {
		t.Errorf("the routes arent displayed properly\n")
	}

	tests := []struct{
		Path string
	} {
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