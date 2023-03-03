package middlewares

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"devnotes.com/models"
	"devnotes.com/utility"
	"github.com/golang-jwt/jwt"
)

func Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get token form cookie
		cookie, err := r.Cookie("token")
		if err != nil {
			http.Redirect(w, r, "/dashboard/login", http.StatusFound)
		} else {
			// Parse token from jwt with secret key
			token, err := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
				_, ok := token.Method.(*jwt.SigningMethodHMAC)
				if !ok {
					return nil, fmt.Errorf("unexpected signing method")
				}
				return []byte(os.Getenv("JWT_SECRET")), nil
			})
			if err != nil {
				utility.ServerErr(w, err)
				return
			}

			// Get data form token
			claims, ok := token.Claims.(jwt.MapClaims)
			if ok && token.Valid {
				username := claims["username"].(string)
				email := claims["email"].(string)
				id := claims["id"].(string)

				user := &models.User{
					Username: username,
					Email:    email,
					Id:       id,
				}

				// Set data to req context
				ctx := context.WithValue(r.Context(), "user", user)
				r = r.WithContext(ctx)
				next.ServeHTTP(w, r)

			} else {
				utility.ServerErr(w, err)
				return
			}
		}
	})
}
