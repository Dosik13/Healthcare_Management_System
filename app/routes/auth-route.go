package routes

import (
	"Healthcare_Management_System/app/controllers"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func RegisterAuth(router *mux.Router, db *gorm.DB) {
	AuthHandler := controllers.NewAuthController(db)

	router.HandleFunc("/login", AuthHandler.LoginHandler).Methods("POST")
	router.HandleFunc("/register", AuthHandler.RegisterHandler).Methods("POST")
	router.HandleFunc("/welcome", AuthHandler.WelcomeHandler).Methods("GET")
	router.HandleFunc("/logout", AuthHandler.LogoutHandler).Methods("GET")
}
