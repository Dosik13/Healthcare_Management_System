package routes

import (
	"Healthcare_Management_System/app/controllers"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// RegisterNurseRoutes registers the routes for the nurse
func RegisterNurseRoutes(router *mux.Router, db *gorm.DB) {
	nurseController := controllers.NewNurseController(db)

	router.HandleFunc("/nurses", nurseController.CreateNurse).Methods("POST")
	router.HandleFunc("/nurses", nurseController.GetAllNurses).Methods("GET")
	router.HandleFunc("/nurses/{nurseID}", nurseController.GetNurse).Methods("GET")
	router.HandleFunc("/nurses/{nurseID}", nurseController.UpdateNurse).Methods("PUT")
	router.HandleFunc("/nurses/{nurseID}", nurseController.DeleteNurse).Methods("DELETE")
}
