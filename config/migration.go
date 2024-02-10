package config

import (
	"Healthcare_Management_System/app/models"
	"gorm.io/gorm"
	"log"
)

func Migration(db *gorm.DB) {
	err := db.AutoMigrate(
		&models.User{},
		&models.Doctor{},
		&models.Nurse{},
		&models.EmergencyAlert{},
		&models.Appointment{},
		&models.Billing{},
		&models.Hospital{},
		&models.Patient{},
		&models.MedicalRecord{},
		&models.Prescription{},
		&models.Rating{},
	)
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
}
