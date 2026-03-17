package routes

import (
	"net/http"
)

type Tree struct {
	Handler http.HandlerFunc
	Method string
	Path   string
}

func Routes(trees []Tree) (http.Handler, []Tree) {
	mux := http.NewServeMux()
	var newTree []Tree
	// Serve static files
	fs := http.FileServer(http.Dir("../internal/static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
	
	// Register handlers
	for _,i := range trees {
		mux.HandleFunc(i.Path, i.Handler)
		path := i.Path
		if i.Path != "/" {
			path = i.Path + "?id={id}"
		}
		newTree = append(newTree, Tree{
			Method: i.Method,
			Path:   path,
		})
	}
	
	return mux, newTree
}
