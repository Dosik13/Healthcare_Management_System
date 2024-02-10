package models

import (
	"gorm.io/gorm"
	"time"
)

type Rating struct {
	gorm.Model
	Score     int       `gorm:"type:int;not null"` // Assuming a score from 1 to 5
	Comment   string    `gorm:"type:text"`         // Optional review text
	RatedAt   time.Time `gorm:"not null"`          // The date when the rating was given
	PatientID uint      `gorm:"not null"`          // ID of the patient who gave the rating
	DoctorID  uint      `gorm:"not null"`          // ID of the doctor being rated
}
