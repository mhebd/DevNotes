package controller

import (
	"context"
	"fmt"
	"html/template"
	"net/http"
	"strings"
	"time"

	"devnotes.com/db"
	"devnotes.com/middlewares"
	"devnotes.com/models"
	"devnotes.com/utility"
	"devnotes.com/views"
	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/bson"
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

	fmt.Print(content)

	page := models.Page{
		Title:   title,
		Content: template.HTML(content),
		Slug:    strings.ReplaceAll(strings.ToLower(title), " ", "_"),
		Created: string(time.Now().Format("January 2, 2006")),
	}

	_, err = db.Page.InsertOne(context.Background(), page)
	if err != nil {
		utility.ServerErr(w, err)
		return
	}

	http.Redirect(w, r, "/dashboard/pages", http.StatusFound)
}

// Delete page by id
func DeletePage(w http.ResponseWriter, r *http.Request) {
	err := utility.DeleteById(w, r, db.Page)
	if err != nil {
		utility.ServerErr(w, err)
		return
	}

	http.Redirect(w, r, "/dashboard/pages", http.StatusFound)
}

// Render single page
func RenderPage(w http.ResponseWriter, r *http.Request) {
	slug := chi.URLParam(r, "slug")

	var page models.Page
	err := db.Page.FindOne(context.Background(), bson.M{"slug": slug}).Decode(&page)
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
		Page  models.Page
		Menus []models.Menu
	}{
		Page:  page,
		Menus: rootData.Menus,
	}

	views.RenderF(w, "page.gohtml", data)
}
