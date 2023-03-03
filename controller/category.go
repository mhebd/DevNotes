package controller

import (
	"context"
	"net/http"
	"time"

	"devnotes.com/db"
	"devnotes.com/models"
	"devnotes.com/utility"
	"devnotes.com/views"
)

// Fetch all categories for dashboard categories page
func CategoryPage(w http.ResponseWriter, r *http.Request) {
	var categories []models.Category
	utility.FetchMany(w, db.Category, &categories)
	views.RenderB(w, "categories.gohtml", categories)
}

// Render create category form page
func CreateCategoryPage(w http.ResponseWriter, r *http.Request) {
	views.RenderB(w, "create-category.gohtml", nil)
}

// Create a category
func CreateCategory(w http.ResponseWriter, r *http.Request) {
	// Parse form
	err := r.ParseForm()
	if err != nil {
		utility.ServerErr(w, err)
		return
	}

	// Get value
	category := r.Form.Get("category")
	if category == "" {
		utility.ClientErr(w, utility.ErrMsg("Please write a category name"))
		return
	}

	// Create category instant
	newCategory := models.Category{
		Category: category,
		Created:  string(time.Now().Format("January 2, 2006")),
	}

	// Create category doc in database
	_, err = db.Category.InsertOne(context.Background(), newCategory)
	if err != nil {
		utility.ServerErr(w, err)
		return
	}

	http.Redirect(w, r, "/dashboard/categories", http.StatusFound)
}

// Delete a category
func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	err := utility.DeleteById(w, r, db.Category)
	if err != nil {
		utility.ServerErr(w, err)
		return
	}
	http.Redirect(w, r, "/dashboard/categories", http.StatusFound)
}
