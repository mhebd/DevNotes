package controller

import (
	"net/http"

	"devnotes.com/db"
	"devnotes.com/models"
	"devnotes.com/utility"
	"devnotes.com/views"
)

func Home(w http.ResponseWriter, r *http.Request) {
	var notes []models.Note
	err := utility.FetchMany(w, db.Note, &notes)
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

	var data = struct {
		Notes      []models.Note
		Categories []models.Category
	}{
		Notes:      notes,
		Categories: categories,
	}

	views.RenderF(w, "home.gohtml", data)
}
