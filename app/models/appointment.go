package models

import (
	"time"
)

type Appointment struct {
	ID          uint      `gorm:"primaryKey"`
	PatientID   uint      `gorm:"type:varchar(255);not null"` // Foreign key from Patient model
	DoctorID    uint      `gorm:"type:varchar(255);not null"` // Foreign key from Doctor model
	HospitalID  uint      `gorm:"type:varchar(255);not null"` // Foreign key from Hospital model
	ScheduledAt time.Time `gorm:"not null"`
	Status      string    `gorm:"not null"` // e.g., Scheduled, Cancelled, Completed
	Reason      string    `gorm:"not null"` // Reason for the appointment
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}
