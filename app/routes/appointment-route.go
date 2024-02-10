package routes

import (
	"Healthcare_Management_System/app/controllers" // Ensure this import path matches your project structure
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// RegisterAppointmentRoutes sets up the routes for appointment management.
func RegisterAppointmentRoutes(router *mux.Router, db *gorm.DB) {
	appointmentController := controllers.NewAppointmentController(db)

	router.HandleFunc("/appointments", appointmentController.CreateAppointment).Methods("POST")
	router.HandleFunc("/appointments", appointmentController.GetAllAppointments).Methods("GET")
	router.HandleFunc("/appointments/{appointmentID}", appointmentController.GetAppointment).Methods("GET")
	router.HandleFunc("/appointments/{appointmentID}", appointmentController.UpdateAppointment).Methods("PUT")
	router.HandleFunc("/appointments/{appointmentID}", appointmentController.DeleteAppointment).Methods("DELETE")
}
