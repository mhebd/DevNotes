package controller

import (
	"net/http"

	"devnotes.com/views"
)

func Home(w http.ResponseWriter, r *http.Request) {
	views.Render(w, "views/pages/home.gohtml", nil)
}
