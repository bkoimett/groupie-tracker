package utils

import (
	"embed"
	"html/template"
	"net/http"
)

//go:embed templates/error.html
var errorTemplateFS embed.FS
var errorTemplate *template.Template

func init() {
	errorTemplate = template.Must(template.ParseFS(errorTemplateFS, "templates/error.html"))
}

type ErrorData struct {
	Code    int
	Message string
}

func ErrorHandler(w http.ResponseWriter, r *http.Request, code int, msg string) {
	if code < 100 || code > 999 {
		code = http.StatusInternalServerError
	}
	w.WriteHeader(code)

	data := ErrorData{
		Code:    code,
		Message: msg,
	}

	err := errorTemplate.Execute(w, data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
