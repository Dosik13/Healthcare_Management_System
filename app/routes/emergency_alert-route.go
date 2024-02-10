package routes

import (
	"Healthcare_Management_System/app/controllers"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func RegisterEmergencyAlertRoutes(router *mux.Router, db *gorm.DB) {
	emergencyAlertController := controllers.NewEmergencyAlertController(db)

	router.HandleFunc("/emergencyAlerts", emergencyAlertController.CreateEmergencyAlert).Methods("POST")
	router.HandleFunc("/emergencyAlerts", emergencyAlertController.GetAllEmergencyAlerts).Methods("GET")
	router.HandleFunc("/emergencyAlerts/{id}", emergencyAlertController.GetEmergencyAlert).Methods("GET")
	router.HandleFunc("/emergencyAlerts/{id}", emergencyAlertController.UpdateEmergencyAlert).Methods("PUT")
	router.HandleFunc("/emergencyAlerts/{id}", emergencyAlertController.DeleteEmergencyAlert).Methods("DELETE")
}
