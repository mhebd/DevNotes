package router

import (
	"net/http"

	"devnotes.com/controller"
	"devnotes.com/middlewares"
	"github.com/go-chi/chi/v5"
)

func Router() chi.Router {
	r := chi.NewRouter()
	// r.Use(middleware.Logger)
	r.Get("/", controller.Home)
	r.Get("/article/{id}", controller.SingleNote)
	r.Get("/dashboard/login", controller.LoginPage)
	r.Post("/dashboard/login", controller.Login)
	r.Get("/dashboard/register", controller.RegisterPage)
	r.Post("/dashboard/register", controller.CreateUser)
	r.NotFound(controller.NotfoundPage)

	// Private route
	r.Group(func(r chi.Router) {
		r.Use(middlewares.Authentication)
		r.Get("/dashboard", controller.Dashboard)
		r.Get("/dashboard/notes", controller.NotesPage)
		r.Get("/dashboard/notes/create-note", controller.CreateNotePage)
		r.Post("/dashboard/notes/create-note", controller.CreateNote)
		r.Get("/dashboard/categories", controller.CategoryPage)
		r.Get("/dashboard/categories/{id}", controller.DeleteCategory)
		r.Get("/dashboard/categories/create-category", controller.CreateCategoryPage)
		r.Post("/dashboard/categories/create-category", controller.CreateCategory)
	})

	fs := http.FileServer(http.Dir("static"))
	r.Handle("/static/*", http.StripPrefix("/static", fs))

	return r
}
