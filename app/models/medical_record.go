package models

import (
	"time"
)

type MedicalRecord struct {
	ID         uint         `gorm:"primaryKey"`
	PatientID  uint         `gorm:"not null"`
	DoctorID   uint         `gorm:"not null"`
	RecordType string       `gorm:"type:varchar(255);not null"` // e.g., Diagnosis, Treatment, Prescription
	Details    string       `gorm:"type:text;not null"`
	Prescript  Prescription `gorm:"foreignKey:MedicalRecordID"`
	Date       time.Time    `gorm:"not null"`
	CreatedAt  time.Time    `gorm:"autoCreateTime"`
	UpdatedAt  time.Time    `gorm:"autoUpdateTime"`
}
