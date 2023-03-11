package controller

import (
	"context"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"time"

	"devnotes.com/db"
	"devnotes.com/middlewares"
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
	var article models.Note
	err := utility.FetchById(w, r, &article)
	if err != nil {
		utility.ServerErr(w, err)
		return
	}

	rootData, err := middlewares.RootData(w)
	if err != nil {
		utility.ServerErr(w, err)
		return
	}

	data := struct {
		Article    models.Note
		Categories []models.Category
		Menus      []models.Menu
	}{
		Article:    article,
		Categories: rootData.Categories,
		Menus:      rootData.Menus,
	}

	views.RenderF(w, "article.gohtml", data)
}

// Create note form page
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
		utility.ServerErr(w, err)
		return
	}

	// Get the file from the request
	file, handler, err := r.FormFile("cover-img")
	if err != nil {
		utility.ClientErr(w, err)
		return
	}
	defer file.Close()

	// Create the file in the uploads directory
	f, err := os.OpenFile("./static/images/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		utility.ServerErr(w, err)
		return
	}
	defer f.Close()

	// Copy the contents of the uploaded file to the created file
	_, err = io.Copy(f, file)
	if err != nil {
		utility.ServerErr(w, err)
		return
	}

	title := r.Form.Get("title")
	category := r.Form.Get("category")
	content := r.Form.Get("content")
	excerpt := r.Form.Get("excerpt")
	tag := r.Form.Get("tag")
	featured := false

	if r.Form.Get("featured") == "on" {
		featured = true
	}

	newNote := models.Note{
		Title:    title,
		Category: category,
		Image:    "/static/images/" + handler.Filename,
		Content:  template.HTML(content),
		Writer:   "Admin",
		Excerpt:  excerpt,
		Tag:      tag,
		Featured: featured,
		Created:  string(time.Now().Format("January 2, 2006 15:04:05")),
	}

	_, err = db.Note.InsertOne(context.Background(), newNote)
	if err != nil {
		utility.ServerErr(w, err)
		return
	}

	http.Redirect(w, r, "/dashboard/notes", http.StatusFound)
}

// Delete a note by id
func DeleteNote(w http.ResponseWriter, r *http.Request) {
	err := utility.DeleteById(w, r, db.Note)
	if err != nil {
		utility.ServerErr(w, err)
		return
	}

	http.Redirect(w, r, "/dashboard/notes", http.StatusFound)
}

// Render Update a note page
func UpdateNotePage(w http.ResponseWriter, r *http.Request) {
	var article models.Note
	err := utility.FetchById(w, r, &article)
	if err != nil {
		utility.ServerErr(w, err)
		return
	}

	var categories []models.Category
	err = utility.FetchMany(w, db.Category, &categories)
	if err != nil {
		utility.ServerErr(w, err)
		return
	}

	data := struct {
		Categories []models.Category
		Article    models.Note
	}{
		Categories: categories,
		Article:    article,
	}
	views.RenderB(w, "update-note.gohtml", data)
}

// Update a note by id
func UpdateNote(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		utility.ClientErr(w, err)
		return
	}

	filter := bson.M{"_id": oid}
	var article models.Note
	if err := db.Note.FindOne(context.Background(), filter).Decode(&article); err != nil {
		utility.ServerErr(w, err)
		return
	}

	err = r.ParseMultipartForm(32 << 20) // 32MB maximum form size
	if err != nil {
		utility.ServerErr(w, err)
		return
	}

	title := r.Form.Get("title")
	category := r.Form.Get("category")
	content := r.Form.Get("content")
	excerpt := r.Form.Get("excerpt")
	tag := r.Form.Get("tag")
	featured := false

	if r.Form.Get("featured") == "on" {
		featured = true
	}

	updateNote := models.Note{
		Title:    title,
		Category: category,
		Content:  template.HTML(content),
		Writer:   "Admin",
		Image:    article.Image,
		Excerpt:  excerpt,
		Tag:      tag,
		Featured: featured,
		Created:  article.Created,
	}

	// Get the file from the request
	file, handler, err := r.FormFile("cover-img")
	if err != nil {
		// err handler
		fmt.Print("no file")
	} else {
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
			utility.ServerErr(w, err)
			return
		}

		if handler.Filename != "" {
			updateNote.Image = "/static/images/" + handler.Filename
		}
	}

	_, err = db.Note.UpdateOne(context.Background(), filter, bson.M{"$set": updateNote})
	if err != nil {
		utility.ServerErr(w, err)
		return
	}

	http.Redirect(w, r, "/dashboard/notes", http.StatusFound)

}
