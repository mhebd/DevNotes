package controller

import (
	"context"
	"net/http"
	"time"

	"devnotes.com/db"
	"devnotes.com/models"
	"devnotes.com/utility"
	"devnotes.com/views"
)

// Render pages page
func PagesPage(w http.ResponseWriter, r *http.Request) {
	var pages []models.Page
	err := utility.FetchMany(w, db.Page, &pages)
	if err != nil {
		utility.ServerErr(w, err)
		return
	}

	views.RenderB(w, "pages.gohtml", pages)
}

// Render create page form
func CreatePageForm(w http.ResponseWriter, r *http.Request) {
	views.RenderB(w, "create-page.gohtml", nil)
}

// Create a new page
func CreatePage(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		utility.ClientErr(w, err)
		return
	}

	title := r.Form.Get("title")
	content := r.Form.Get("content")

	page := models.Page{
		Title:   title,
		Content: content,
		Created: string(time.Now().Format("January 2, 2006 15:04:05")),
	}

	_, err = db.Page.InsertOne(context.Background(), page)
	if err != nil {
		utility.ServerErr(w, err)
		return
	}

	http.Redirect(w, r, "/dashboard/pages", http.StatusFound)
}
