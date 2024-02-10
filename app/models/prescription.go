package models

import "time"

type Prescription struct {
	ID              uint      `gorm:"primaryKey"`
	MedicalRecordID uint      `gorm:"not null"` // Foreign key from Medical Record model
	Medication      string    `gorm:"type:varchar(255);not null"`
	Dosage          string    `gorm:"not null"`
	Frequency       string    `gorm:"not null"`
	Duration        string    `gorm:"not null"`
	CreatedAt       time.Time `gorm:"autoCreateTime"`
	UpdatedAt       time.Time `gorm:"autoUpdateTime"`
}
