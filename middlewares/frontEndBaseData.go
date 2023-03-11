package middlewares

import (
	"net/http"

	"devnotes.com/db"
	"devnotes.com/models"
	"devnotes.com/utility"
)

func RootData(w http.ResponseWriter) (models.RootData, error) {
	var data models.RootData

	var categories []models.Category
	err := utility.FetchMany(w, db.Category, &categories)
	if err != nil {
		return data, err
	}
	data.Categories = categories

	var menus []models.Menu
	err = utility.FetchMany(w, db.Menu, &menus)
	if err != nil {
		return data, err
	}
	data.Menus = menus

	return data, nil
}
