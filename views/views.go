package views

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

func Render(w http.ResponseWriter, file string, data interface{}) {
	// Set content type header
	w.Header().Set("Content-Type", "text/html")

	// Make a file array and create a executable template
	files := []string{file}
	files = append(files, "views/layout/base-layout.gohtml")
	files = append(files, components()...)
	t, err := template.ParseFiles(files...)
	if err != nil {
		log.Fatal(err)
	}

	// Execute template with name and data
	t.ExecuteTemplate(w, "base-layout", data)
}

func components() []string {
	files, err := filepath.Glob("views/components/*.gohtml")
	if err != nil {
		log.Fatal(err)
	}
	return files
}
