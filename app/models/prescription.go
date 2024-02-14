package models

import "time"

type Prescription struct {
	ID              uint      `gorm:"primaryKey"`
	MedicalRecordID uint      `gorm:"not null"` // Foreign key from Medical Record model
	MedicationInfo  string    `gorm:"type:varchar(255);not null"`
	CreatedAt       time.Time `gorm:"autoCreateTime"`
	UpdatedAt       time.Time `gorm:"autoUpdateTime"`
}
