package models

import (
	"time"
)

type Appointment struct {
	ID          uint      `gorm:"primaryKey"`
	PatientID   uint      `gorm:"not null"`
	DoctorID    uint      `gorm:"not null"`
	HospitalID  uint      `gorm:"not null"`
	ScheduledAt time.Time `gorm:"not null"`
	Status      string    `gorm:"not null"` // e.g., Scheduled, Cancelled, Completed
	Reason      string    `gorm:"not null"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}
