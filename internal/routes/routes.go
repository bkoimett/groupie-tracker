package routes

import (
    // "groupie-tracker/internal/handlers"
    "net/http"
    "path/filepath"
    "strings"
)

type Routes struct {
    Method  string
    Path    string
    Handler http.HandlerFunc
}

func NewRoutes(routeList []Routes) http.Handler {
    mux := http.NewServeMux()
    
    // Custom static file handler
    fs := http.FileServer(http.Dir("internal/static"))
    mux.Handle("/static/", http.StripPrefix("/static/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Set correct MIME types
        ext := strings.ToLower(filepath.Ext(r.URL.Path))
        switch ext {
        case ".css":
            w.Header().Set("Content-Type", "text/css; charset=utf-8")
        case ".js":
            w.Header().Set("Content-Type", "application/javascript; charset=utf-8")
        }
        
        // Remove nosniff header or keep it but with correct types
        // w.Header().Del("X-Content-Type-Options")
        
        fs.ServeHTTP(w, r)
    })))
    
    // Register all routes
    for _, route := range routeList {
        mux.HandleFunc(route.Path, route.Handler)
    }
    
    return mux
}