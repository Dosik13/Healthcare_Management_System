package routes

import (
	"Healthcare_Management_System/app/controllers"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func RegisterHospitalRoutes(router *mux.Router, db *gorm.DB) {
	hospitalController := controllers.NewHospitalController(db)

	router.HandleFunc("/hospitals", hospitalController.CreateHospital).Methods("POST")
	router.HandleFunc("/hospitals", hospitalController.GetAllHospitals).Methods("GET")
	router.HandleFunc("/hospitals/{id}", hospitalController.GetHospital).Methods("GET")
	router.HandleFunc("/hospitals/{id}", hospitalController.UpdateHospital).Methods("PUT")
	router.HandleFunc("/hospitals/{id}", hospitalController.DeleteHospital).Methods("DELETE")
}
