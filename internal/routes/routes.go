package routes

import (
	"net/http"
)

type Routes struct {
	Method  string
	Path    string
	Handler http.HandlerFunc
}

func NewRoutes(routes []Routes) http.Handler {
	mux := http.NewServeMux()
	// Serve static files
	fs := http.FileServer(http.Dir("../internal/static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	// Register handlers
	for _, i := range routes {
		mux.HandleFunc(i.Path, i.Handler)
	}

	return mux
}
