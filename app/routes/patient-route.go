package routes

import (
	"Healthcare_Management_System/app/controllers"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// RegisterPatientRoutes registers the routes for the patient
func RegisterPatientRoutes(router *mux.Router, db *gorm.DB) {
	patientController := controllers.NewPatientController(db)

	router.HandleFunc("/patients", patientController.CreatePatient).Methods("POST")
	router.HandleFunc("/patients", patientController.GetAllPatients).Methods("GET")
	router.HandleFunc("/patients/{patientID}", patientController.GetPatient).Methods("GET")
	router.HandleFunc("/patients/{patientID}", patientController.UpdatePatient).Methods("PUT")
	router.HandleFunc("/patients/{patientID}", patientController.DeletePatient).Methods("DELETE")
}
