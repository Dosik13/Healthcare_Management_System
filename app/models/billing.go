package models

import "time"

type Billing struct {
	ID            uint      `gorm:"primaryKey"`
	PatientID     uint      `gorm:"not null"` // Foreign key from Patient model
	AppointmentID uint      `gorm:"not null"` // Foreign key from Appointment model
	Amount        float64   `gorm:"not null"`
	Status        string    `gorm:"not null"` // e.g., Pending, Paid, Overdue
	DueDate       time.Time `gorm:"not null"`
	CreatedAt     time.Time `gorm:"autoCreateTime"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime"`
}
