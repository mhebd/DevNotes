package controller

import (
	"context"
	"net/http"
	"strings"

	"devnotes.com/db"
	"devnotes.com/models"
	"devnotes.com/utility"
	"devnotes.com/views"
)

// Render Menu page
func MenusPage(w http.ResponseWriter, r *http.Request) {
	var menus []models.Menu
	err := utility.FetchMany(w, db.Menu, &menus)
	if err != nil {
		utility.ServerErr(w, err)
		return
	}

	views.RenderB(w, "menus.gohtml", menus)
}

// Create a menu
func CreateMenu(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		utility.ClientErr(w, err)
		return
	}

	title := r.Form.Get("title")
	link := r.Form.Get("link")

	menu := models.Menu{
		Title: strings.ToLower(title),
		Link:  link,
	}

	_, err = db.Menu.InsertOne(context.Background(), menu)
	if err != nil {
		utility.ServerErr(w, err)
		return
	}

	http.Redirect(w, r, "/dashboard/menus", http.StatusFound)
}
