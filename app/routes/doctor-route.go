package routes

import (
	"Healthcare_Management_System/app/controllers"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

//func (dc *DoctorController) SetupRoutes(router *mux.Router) {
//	router.Handle("/doctors", middleware.AuthMiddleware(http.HandlerFunc(dc.GetAllDoctors))).Methods("GET")
//	router.Handle("/doctors/{id}", middleware.AuthMiddleware(http.HandlerFunc(dc.GetDoctor))).Methods("GET")
//	router.Handle("/doctors", middleware.AuthMiddleware(http.HandlerFunc(dc.CreateDoctor))).Methods("POST")
//	router.Handle("/doctors/{id}", middleware.AuthMiddleware(http.HandlerFunc(dc.UpdateDoctor))).Methods("PUT")
//	router.Handle("/doctors/{id}", middleware.AuthMiddleware(http.HandlerFunc(dc.DeleteDoctor))).Methods("DELETE")
//}

func RegisterDoctorRoutes(router *mux.Router, db *gorm.DB) {
	doctorController := controllers.NewDoctorController(db)

	router.HandleFunc("/doctors", doctorController.CreateDoctor).Methods("POST")
	router.HandleFunc("/doctors", doctorController.GetAllDoctors).Methods("GET")
	router.HandleFunc("/doctors/{doctorID}", doctorController.GetDoctor).Methods("GET")
	router.HandleFunc("/doctors/{doctorID}", doctorController.UpdateDoctor).Methods("PUT")
	router.HandleFunc("/doctors/{doctorID}", doctorController.DeleteDoctor).Methods("DELETE")
}
