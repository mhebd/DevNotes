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
	r.Get("/page/{slug}", controller.RenderPage)
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
		r.Get("/dashboard/notes/update-note/{id}", controller.UpdateNotePage)
		r.Post("/dashboard/notes/update-note/{id}", controller.UpdateNote)
		r.Get("/dashboard/notes/delete/{id}", controller.DeleteNote)

		r.Get("/dashboard/categories", controller.CategoryPage)
		r.Get("/dashboard/categories/{id}", controller.DeleteCategory)
		r.Get("/dashboard/categories/create-category", controller.CreateCategoryPage)
		r.Post("/dashboard/categories/create-category", controller.CreateCategory)

		r.Get("/dashboard/pages", controller.PagesPage)
		r.Get("/dashboard/pages/create-page", controller.CreatePageForm)
		r.Post("/dashboard/pages/create-page", controller.CreatePage)
		r.Get("/dashboard/pages/delete/{id}", controller.DeletePage)

		r.Get("/dashboard/menus", controller.MenusPage)
		r.Post("/dashboard/menu/create", controller.CreateMenu)
	})

	fs := http.FileServer(http.Dir("static"))
	r.Handle("/static/*", http.StripPrefix("/static", fs))

	return r
}
