package models

import (
	"time"
)

type Appointment struct {
	ID         uint `gorm:"primaryKey"`
	PatientID  *uint
	DoctorID   uint   `gorm:"not null"`
	StartTime  string `gorm:"not null"`
	EndTime    string `gorm:"not null"`
	Date       string `gorm:"not null"`
	Status     string `gorm:"not null"` // e.g., Scheduled, Cancelled, Completed
	Reason     string `gorm:""`
	Billing    Billing
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
	HospitalID *uint
}
