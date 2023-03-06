package controller

import (
	"context"
	"io"
	"net/http"
	"os"
	"time"

	"devnotes.com/db"
	"devnotes.com/models"
	"devnotes.com/utility"
	"devnotes.com/views"
	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Fetch all notes for dashboard notes page
func NotesPage(w http.ResponseWriter, r *http.Request) {
	var notes []models.Note
	err := utility.FetchMany(w, db.Note, &notes)
	if err != nil {
		utility.ServerErr(w, err)
		return
	}
	views.RenderB(w, "notes.gohtml", notes)
}

// Single note page
func SingleNote(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return
	}

	// fmt.Print(id)

	var article models.Note
	err = db.Note.FindOne(context.TODO(), bson.D{{Key: "_id", Value: oid}}).Decode(&article)
	if err != nil {
		utility.ServerErr(w, err)
		return
	}

	views.RenderF(w, "article.gohtml", article)
}

// Create a note form page
func CreateNotePage(w http.ResponseWriter, r *http.Request) {
	var categories []models.Category
	err := utility.FetchMany(w, db.Category, &categories)
	if err != nil {
		utility.ServerErr(w, err)
		return
	}

	views.RenderB(w, "create-notes.gohtml", categories)
}

// Create a note
func CreateNote(w http.ResponseWriter, r *http.Request) {
	// Parse the multipart form
	err := r.ParseMultipartForm(32 << 20) // 32MB maximum form size
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Get the file from the request
	file, handler, err := r.FormFile("cover-img")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Create the file in the uploads directory
	f, err := os.OpenFile("./static/images/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer f.Close()

	// Copy the contents of the uploaded file to the created file
	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	title := r.Form.Get("title")
	category := r.Form.Get("category")
	content := r.Form.Get("content")
	excerpt := r.Form.Get("excerpt")
	tag := r.Form.Get("tag")

	newNote := models.Note{
		Title:    title,
		Category: category,
		Image:    "/static/images/" + handler.Filename,
		Content:  content,
		Writer:   "Admin",
		Excerpt:  excerpt,
		Tag:      tag,
		Created:  string(time.Now().Format("January 2, 2006 15:04:05")),
	}

	_, err = db.Note.InsertOne(context.Background(), newNote)
	if err != nil {
		utility.ServerErr(w, err)
		return
	}

	http.Redirect(w, r, "/dashboard/notes", http.StatusFound)
}
