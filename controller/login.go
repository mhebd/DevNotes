package controller

import (
	"context"
	"net/http"
	"os"
	"time"

	"devnotes.com/db"
	"devnotes.com/models"
	"devnotes.com/utility"
	"devnotes.com/views"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

// Render login page
func LoginPage(w http.ResponseWriter, r *http.Request) {
	views.Render(w, "views/pages/backend/login.gohtml", nil)
}

// Login a existing user
func Login(w http.ResponseWriter, r *http.Request) {
	// Parse form from req body
	err := r.ParseForm()
	if err != nil {
		utility.ClientErr(w, err)
		return
	}

	// Get all data
	email := r.Form.Get("email")
	password := r.Form.Get("password")

	// Find user
	var user models.User
	err = db.User.FindOne(context.TODO(), bson.D{{Key: "email", Value: email}}).Decode(&user)
	if err != nil {
		utility.ClientErr(w, err)
		return
	}

	// Compare bcrypt password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		utility.ServerErr(w, err)
		return
	}

	// Create JWT token
	jwtSecret := []byte(os.Getenv("JWT_SECRET"))
	jwtData := jwt.MapClaims{
		"username": user.Username,
		"email":    user.Email,
		"id":       user.Id,
		"iat":      time.Now().Unix(),
		"exp":      time.Now().Add(time.Hour * 48).Unix(),
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtData)
	jwtTokenString, err := jwtToken.SignedString(jwtSecret)
	if err != nil {
		utility.ServerErr(w, err)
		return
	}

	// Create a req cookie and save it
	token := http.Cookie{
		Name:   "token",
		Value:  jwtTokenString,
		MaxAge: 3600 * 48,
	}
	http.SetCookie(w, &token)

	// Redirect to dashboard
	http.Redirect(w, r, "/dashboard", http.StatusFound)
}
