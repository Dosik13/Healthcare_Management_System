package routes

import (
	"Healthcare_Management_System/app/controllers"
	"Healthcare_Management_System/utils"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"net/http"
)

// RegisterMedicalRecordRoutes registers the routes for the medical record
func RegisterMedicalRecordRoutes(router *mux.Router, db *gorm.DB) {
	medicalRecordController := controllers.NewMedicalRecordController(db)

	router.Handle("/medical-records", utils.AuthDoctorHandler(http.HandlerFunc(medicalRecordController.CreateMedicalRecord))).Methods("GET")
	router.Handle("/medical-records", utils.AuthDoctorHandler(http.HandlerFunc(medicalRecordController.CreateMedicalRecord))).Methods("POST")
	router.HandleFunc("/medical-records", medicalRecordController.GetAllMedicalRecords).Methods("GET")
	router.HandleFunc("/medical-records/{id}", medicalRecordController.GetMedicalRecord).Methods("GET")
	router.HandleFunc("/medical-records/{id}", medicalRecordController.UpdateMedicalRecord).Methods("PUT")
	router.HandleFunc("/medical-records/{id}", medicalRecordController.DeleteMedicalRecord).Methods("DELETE")
}
