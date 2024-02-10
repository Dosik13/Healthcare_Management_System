package routes

import (
	"Healthcare_Management_System/app/controllers"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// RegisterPrescriptionRoutes registers the routes for the prescription
func RegisterPrescriptionRoutes(router *mux.Router, db *gorm.DB) {
	prescriptionController := controllers.NewPrescriptionController(db)

	router.HandleFunc("/prescriptions", prescriptionController.CreatePrescription).Methods("POST")
	router.HandleFunc("/prescriptions", prescriptionController.GetAllPrescriptions).Methods("GET")
	router.HandleFunc("/prescriptions/{prescriptionID}", prescriptionController.GetPrescription).Methods("GET")
	router.HandleFunc("/prescriptions/{prescriptionID}", prescriptionController.UpdatePrescription).Methods("PUT")
	router.HandleFunc("/prescriptions/{prescriptionID}", prescriptionController.DeletePrescription).Methods("DELETE")
}
