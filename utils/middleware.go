package utils

import (
	"Healthcare_Management_System/app/models"
	"context"
	"github.com/gorilla/sessions"
	"log"
	"net/http"
)

var Store = sessions.NewCookieStore([]byte("your-secret-key"))

func GetUserMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := Store.Get(r, "sessionID")
		if err != nil {
			http.Error(w, "Session error", http.StatusInternalServerError)
			return
		}

		user, ok := session.Values["user"].(models.User)
		if !ok || user.UserID == 0 {
			http.Error(w, "Unauthorized access", http.StatusUnauthorized)
			return
		}

		// Embed user in request context
		ctx := context.WithValue(r.Context(), "user", user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func RoleAccessMiddleware(requiredRole string, next http.Handler) http.Handler {
	return GetUserMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, ok := r.Context().Value("user").(models.User)
		if !ok {
			http.Error(w, "Unauthorized access", http.StatusUnauthorized)
			return
		}

		if user.Role != requiredRole {
			log.Printf("Access denied: %v is not a %v or an administrator", user.UserID, requiredRole)
			http.Error(w, "Access denied", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	}))
}
