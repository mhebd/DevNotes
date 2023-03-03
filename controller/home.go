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
	utility.FetchMany(w, db.Note, &notes)
	views.RenderF(w, "home.gohtml", notes)
}
