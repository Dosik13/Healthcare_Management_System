package utils

import (
	"github.com/gorilla/sessions"
	"net/http"
	"os"
)

var Store = sessions.NewCookieStore([]byte(os.Getenv("SK")))

func AuthenticatedPatient(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := Store.Get(r, "SessionID")
		if err != nil {
			http.Error(w, "Session error", http.StatusInternalServerError)
			return
		}
		if role, ok := session.Values["role"].(string); !ok || role != "patient" {
			http.Error(w, "Access denied", http.StatusForbidden)
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		next.ServeHTTP(w, r)
	})
}
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := Store.Get(r, "SessionID")
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		if userId, ok := session.Values["user"].(uint); !ok || userId == 0 {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func AuthDoctorHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := Store.Get(r, "SessionID")
		if err != nil {
			http.Error(w, "Session error", http.StatusInternalServerError)
			return
		}
		if role, ok := session.Values["role"].(string); !ok || role != "doctor" {
			http.Error(w, "Access denied", http.StatusForbidden)
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func AuthNurseHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := Store.Get(r, "SessionID")
		if err != nil {
			http.Error(w, "Session error", http.StatusInternalServerError)
			return
		}
		if role, ok := session.Values["role"].(string); !ok || role != "nurse" {
			http.Error(w, "Access denied", http.StatusForbidden)
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func CheckIfLogged(w http.ResponseWriter, r *http.Request) bool {

	session, err := Store.Get(r, "SessionID")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return true
	}
	role, ok := session.Values["role"].(string)

	if ok && role == "doctor" {
		http.Redirect(w, r, "/doctor_dashboard", http.StatusSeeOther)
		return true
	}
	if ok && role == "nurse" {
		http.Redirect(w, r, "/nurse_dashboard", http.StatusSeeOther)
		return true
	}
	if ok && role == "patient" {
		http.Redirect(w, r, "/patient_dashboard", http.StatusSeeOther)
		return true
	}
	return false
}
