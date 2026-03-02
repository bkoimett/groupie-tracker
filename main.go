package main

import (
    "fmt"
    "log"
    "net/http"
    "groupie-tracker/handlers"
)

func main() {
    // Serve static files
    fs := http.FileServer(http.Dir("static"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

    // Register handlers
    http.HandleFunc("/", handlers.HomeHandler)

    // Start server
    port := ":8080"
    fmt.Printf("🚀 Server starting on http://localhost%s\n", port)
    fmt.Println("📦 Press Ctrl+C to stop")
    
    err := http.ListenAndServe(port, nil)
    if err != nil {
        log.Fatal("Server failed to start:", err)
    }
}