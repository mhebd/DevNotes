package controller

import (
	"net/http"

	"devnotes.com/views"
)

func NotfoundPage(w http.ResponseWriter, r *http.Request) {
	views.Render(w, "views/pages/frontend/404.gohtml", nil)
}
