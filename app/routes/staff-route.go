package routes

import (
	"Healthcare_Management_System/app/controllers"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// RegisterStaffRoutes registers the routes for the staff
func RegisterStaffRoutes(router *mux.Router, db *gorm.DB) {
	staffController := controllers.NewHospitalStaffController(db)

	router.HandleFunc("/staff", staffController.CreateHospitalStaff).Methods("POST")
	router.HandleFunc("/staff", staffController.GetAllHospitalStaff).Methods("GET")
	router.HandleFunc("/staff/{staffID}", staffController.GetHospitalStaff).Methods("GET")
	router.HandleFunc("/staff/{staffID}", staffController.UpdateHospitalStaff).Methods("PUT")
	router.HandleFunc("/staff/{staffID}", staffController.DeleteHospitalStaff).Methods("DELETE")
}
