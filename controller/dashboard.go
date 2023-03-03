package controller

import (
	"net/http"

	"devnotes.com/views"
)

func Dashboard(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user")
	views.RenderB(w, "dashboard.gohtml", user)
}
