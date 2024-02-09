package routes

import (
	"Healthcare_Management_System/app/controllers"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var RegisterUserRoutes = func(router *mux.Router, db *gorm.DB) {
	userController := controllers.NewUserController(db)

	router.HandleFunc("/users", userController.CreateUser).Methods("POST")
	router.HandleFunc("/users", userController.GetAllUsers).Methods("GET")
	router.HandleFunc("/users/{userID}", userController.GetUser).Methods("GET")
	router.HandleFunc("/users/{userID}", userController.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{userID}", userController.DeleteUser).Methods("DELETE")
}
