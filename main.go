package main

import (
	"Healthcare_Management_System/app/routes"
	"Healthcare_Management_System/config"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	db := config.ConnectDB()

	defer config.DisconnectDB(db)
	config.Migration(db)

	r := mux.NewRouter()

	routes.RegisterDoctorRoutes(r, db)
	routes.RegisterNurseRoutes(r, db)
	routes.RegisterEmergencyAlertRoutes(r, db)
	routes.RegisterAppointmentRoutes(r, db)
	routes.RegisterBillingRoutes(r, db)
	routes.RegisterPatientRoutes(r, db)
	routes.RegisterMedicalRecordRoutes(r, db)
	routes.RegisterPrescriptionRoutes(r, db)
	routes.RegisterRatingRoutes(r, db)
	routes.RegisterHospitalRoutes(r, db)

	http.Handle("/", r)
	log.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
