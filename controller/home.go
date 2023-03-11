package controller

import (
	"net/http"

	"devnotes.com/db"
	"devnotes.com/middlewares"
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

	var slides []models.Note
	for _, v := range notes {
		if v.Featured {
			slides = append(slides, v)
		}
	}

	rootData, err := middlewares.RootData(w)
	if err != nil {
		utility.ServerErr(w, err)
		return
	}

	var data = struct {
		Notes      []models.Note
		Categories []models.Category
		Slides     []models.Note
		Menus      []models.Menu
	}{
		Notes:      notes,
		Categories: rootData.Categories,
		Slides:     slides,
		Menus:      rootData.Menus,
	}

	views.RenderF(w, "home.gohtml", data)
}
