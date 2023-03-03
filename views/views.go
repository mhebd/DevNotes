package views

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

// Frontend page render
func RenderF(w http.ResponseWriter, file string, data interface{}) {
	// Set content type header
	w.Header().Set("Content-Type", "text/html")

	// Make a file array and create a executable template
	file = "views/pages/frontend/" + file
	files := []string{file}
	files = append(files, "views/layout/frontend-layout.gohtml")
	files = append(files, fComponents()...)
	t, err := template.ParseFiles(files...)
	if err != nil {
		log.Fatal(err)
	}

	// Execute template with name and data
	t.ExecuteTemplate(w, "frontend-layout", data)
}

// Backend page render
func RenderB(w http.ResponseWriter, file string, data interface{}) {
	// Set content type header
	w.Header().Set("Content-Type", "text/html")

	// Make a file array and create a executable template
	file = "views/pages/backend/" + file
	files := []string{file}
	files = append(files, "views/layout/backend-layout.gohtml")
	files = append(files, bComponents()...)
	t, err := template.ParseFiles(files...)
	if err != nil {
		log.Fatal(err)
	}

	// Execute template with name and data
	t.ExecuteTemplate(w, "backend-layout", data)
}

// Render a simple file
func Render(w http.ResponseWriter, file string, data interface{}) {
	w.Header().Set("Content-Type", "text/html")
	t, err := template.ParseFiles(file)
	if err != nil {
		log.Fatal(err)
	}

	t.Execute(w, data)
}

// Make frontend components array
func fComponents() []string {
	files, err := filepath.Glob("views/components/frontend/*.gohtml")
	if err != nil {
		log.Fatal(err)
	}
	return files
}

// Make backend components array
func bComponents() []string {
	files, err := filepath.Glob("views/components/backend/*.gohtml")
	if err != nil {
		log.Fatal(err)
	}
	return files
}
