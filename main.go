package main

import (
	"Healthcare_Management_System/app/models"
	"Healthcare_Management_System/app/routes"
	"Healthcare_Management_System/config"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	db := config.ConnectDB()

	defer config.DisconnectDB(db)

	err := db.AutoMigrate(&models.User{})
	if err != nil {
		return
	} // Auto migrate your models

	r := mux.NewRouter()

	routes.RegisterUserRoutes(r, db)
	routes.RegisterDoctorRoutes(r, db)
	routes.RegisterEmergencyAlertRoutes(r, db)
	routes.RegisterAppointmentRoutes(r, db)
	routes.RegisterBillingRoutes(r, db)
	routes.RegisterHospitalRoutes(r, db)
	routes.RegisterPatientRoutes(r, db)
	routes.RegisterMedicalRecordRoutes(r, db)
	routes.RegisterPatientRoutes(r, db)
	routes.RegisterPrescriptionRoutes(r, db)

	http.Handle("/", r)
	log.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
