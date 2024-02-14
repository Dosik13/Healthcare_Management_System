package models

import (
	"time"
)

type EmergencyAlert struct {
	ID          uint      `gorm:"primaryKey"`
	PatientID   uint      `gorm:"not null"`
	Description string    `gorm:"type:varchar(255);not null"`
	Status      string    `gorm:"type:varchar(255);not null"` // e.g., Reported, Responding, Resolved
	HospitalID  uint      `gorm:"not null"`
	ReportedAt  time.Time `gorm:"not null"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}
