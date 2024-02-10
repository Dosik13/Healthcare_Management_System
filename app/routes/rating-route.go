package routes

import (
	"Healthcare_Management_System/app/controllers"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// RegisterRatingRoutes registers the routes for the rating
func RegisterRatingRoutes(router *mux.Router, db *gorm.DB) {
	ratingController := controllers.NewRatingController(db)

	router.HandleFunc("/ratings", ratingController.CreateRating).Methods("POST")
	router.HandleFunc("/ratings", ratingController.GetAllRatings).Methods("GET")
	router.HandleFunc("/ratings/{ratingID}", ratingController.GetRating).Methods("GET")
	router.HandleFunc("/ratings/{ratingID}", ratingController.UpdateRating).Methods("PUT")
	router.HandleFunc("/ratings/{ratingID}", ratingController.DeleteRating).Methods("DELETE")
}
