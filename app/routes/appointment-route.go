package routes

import (
	"Healthcare_Management_System/app/controllers" // Ensure this import path matches your project structure
	"Healthcare_Management_System/utils"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"net/http"
)

// RegisterAppointmentRoutes sets up the routes for appointment management.
func RegisterAppointmentRoutes(router *mux.Router, db *gorm.DB) {
	appointmentController := controllers.NewAppointmentController(db)

	router.Handle("/appointments/{doctorID}", utils.AuthenticatedPatient(http.HandlerFunc(appointmentController.GetNotScheduledAppointments))).Methods("GET")
	router.Handle("/appointments/schedule", utils.AuthenticatedPatient(http.HandlerFunc(appointmentController.ScheduleAppointment))).Methods("POST")

	router.Handle("/appointments", utils.AuthDoctorHandler(http.HandlerFunc(appointmentController.CreateAppointment))).Methods("POST")
	router.Handle("/appointments", utils.AuthDoctorHandler(http.HandlerFunc(appointmentController.GetAllAppointments))).Methods("GET")

	router.HandleFunc("/appointments/{appointmentID}", appointmentController.GetAppointment).Methods("GET")
	router.HandleFunc("/appointments/{appointmentID}", appointmentController.UpdateAppointment).Methods("PUT")
	router.HandleFunc("/appointments/{appointmentID}", appointmentController.DeleteAppointment).Methods("DELETE")
}
