package models

import (
	"time"
)

type Appointment struct {
	ID          uint      `gorm:"primaryKey"`
	PatientID   uint      `gorm:"not null"` // Foreign key from Patient model
	DoctorID    uint      `gorm:"not null"` // Foreign key from Doctor model
	ScheduledAt time.Time `gorm:"not null"`
	Status      string    `gorm:"not null"` // e.g., Scheduled, Cancelled, Completed
	Reason      string    `gorm:"not null"` // Reason for the appointment
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}
