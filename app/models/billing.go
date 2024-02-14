package models

import "time"

type Billing struct {
	ID            uint `gorm:"primaryKey"`
	PatientID     *uint
	AppointmentID uint    `gorm:"not null"`
	Amount        float64 `gorm:"not null"`
	Status        string  `gorm:"type:varchar(255);not null"` // e.g., Pending, Paid, Overdue
	DueDate       string
	CreatedAt     time.Time `gorm:"autoCreateTime"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime"`
}
