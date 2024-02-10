package models

import (
	"time"
)

type EmergencyAlert struct {
	ID          uint      `gorm:"primaryKey"`
	PatientID   uint      `gorm:"not null"` // Foreign key from Patient model
	Description string    `gorm:"type:varchar(255);not null"`
	Status      string    `gorm:"type:varchar(255);not null"` // e.g., Reported, Responding, Resolved
	ReportedAt  time.Time `gorm:"not null"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}
