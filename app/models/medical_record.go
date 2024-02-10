package models

import (
	"time"
)

type MedicalRecord struct {
	ID         uint      `gorm:"primaryKey"`
	PatientID  uint      `gorm:"not null"` // Foreign key from Patient model
	DoctorID   uint      // Optional: Foreign key from Doctor model, if the record was created by a doctor
	RecordType string    `gorm:"type:varchar(255);not null"` // e.g., Diagnosis, Treatment, Prescription
	Details    string    `gorm:"type:text;not null"`
	Date       time.Time `gorm:"not null"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
}
