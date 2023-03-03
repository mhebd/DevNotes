package controller

import (
	"context"
	"net/http"
	"time"

	"devnotes.com/db"
	"devnotes.com/models"
	"devnotes.com/utility"
	"devnotes.com/views"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

// Render register page
func RegisterPage(w http.ResponseWriter, r *http.Request) {
	views.Render(w, "views/pages/backend/register.gohtml", nil)
}

// Create a user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	// Parse form from req body
	err := r.ParseForm()
	if err != nil {
		utility.ClientErr(w, err)
		return
	}

	// Get all data
	username := r.Form.Get("username")
	email := r.Form.Get("email")
	password := r.Form.Get("password")
	password2 := r.Form.Get("password2")

	// Validate password
	if len(password) < 6 {
		utility.ClientErr(w, utility.ErrMsg("Password is too short."))
		return
	}
	if password != password2 {
		utility.ClientErr(w, utility.ErrMsg("Your password and retype password don't match."))
		return
	}

	// Check user already exist
	var hasUser models.User
	err = db.User.FindOne(context.TODO(), bson.D{{Key: "email", Value: email}}).Decode(&hasUser)
	if err == nil {
		utility.ClientErr(w, utility.ErrMsg("User already exist."))
		return
	}

	// Make a hash string password
	hashPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		utility.ClientErr(w, err)
		return
	}

	user := models.User{
		Username: username,
		Email:    email,
		Roll:     "user",
		Password: string(hashPass),
		Created:  string(time.Now().Format("January 2, 2006")),
	}

	// Create a new user
	_, err = db.User.InsertOne(context.Background(), user)
	if err != nil {
		utility.ServerErr(w, err)
		return
	}

	http.Redirect(w, r, "/dashboard", http.StatusFound)
}
