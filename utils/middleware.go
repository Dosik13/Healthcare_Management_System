package utils

import (
	"Healthcare_Management_System/app/models"
	"context"
	"github.com/gorilla/sessions"
	"log"
	"net/http"
	"os"
)

var Store = sessions.NewCookieStore([]byte(os.Getenv("SK")))

//func GetUserMiddleware(next http.Handler) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		session, err := Store.Get(r, "sessionID")
//		if err != nil {
//			http.Error(w, "Session error", http.StatusInternalServerError)
//			return
//		}
//
//		role, ok := session.Values["role"].(string)
//		if !ok || role == "" {
//			http.Error(w, "Unauthorized access", http.StatusUnauthorized)
//			return
//		}
//
//		// Embed user in request context
//		ctx := context.WithValue(r.Context(), "user", user)
//		next.ServeHTTP(w, r.WithContext(ctx))
//	})
//}
//
//func RoleAccessMiddleware(requiredRole string, next http.Handler) http.Handler {
//	return GetUserMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		user, ok := r.Context().Value("user").(models.User)
//		if !ok {
//			http.Error(w, "Unauthorized access", http.StatusUnauthorized)
//			return
//		}
//
//		if user.Role != requiredRole {
//			log.Printf("Access denied: %v is not a %v or an administrator", user.UserID, requiredRole)
//			http.Error(w, "Access denied", http.StatusForbidden)
//			return
//		}
//
//		next.ServeHTTP(w, r)
//	}))
//}

func GetUserMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := Store.Get(r, "sessionID")
		if err != nil {
			next.ServeHTTP(w, r)
			return
		}

		role, ok := session.Values["role"].(string)

		switch role {
		case "doctor":
			user, ok := session.Values["user"].(*models.Doctor)
		case "nurse":
			user, ok := session.Values["user"].(*models.Nurse)
		case "patient": user, ok := session.Values["user"].(*models.Patient)
		case ""
		}

		if !ok || user == nil {
			next.ServeHTTP(w, r)
			return
		}

		req := r.WithContext(context.WithValue(r.Context(), "user", *user))
		next.ServeHTTP(w, req)
	})
}

func UserAccessRightsMiddleware(next http.Handler) http.Handler {
	return GetUserMiddleware(GetIdParamMiddleware("userId", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, ok := r.Context().Value("user").(entities.User) // logged user
		userId, _ := r.Context().Value("userId").(uint)       // requested user

		forbidden := false
		if !ok {
			forbidden = true
		} else if !user.Administrator && user.ID != userId {
			log.Printf("%t %v %v", user.Administrator, user.ID, userId)
			forbidden = true
		}

		if forbidden {
			http.Error(w, "You are not allowed to access information about this user", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})))
}

func AccessRightsMiddleware(d app.DAO, admin bool, next http.Handler) http.Handler {
	return GetUserMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, ok := r.Context().Value("user").(entities.User)
		if !ok {
			//missing value means user is not logged in
			if admin {
				http.Error(w, "You don't have access rights for this page.", http.StatusForbidden)
			} else {
				http.Redirect(w, r, "/sign-in", http.StatusSeeOther)
			}
			return
		}

		if admin && !user.Administrator {
			http.Error(w, "You don't have access rights for this page.", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	}))
}
