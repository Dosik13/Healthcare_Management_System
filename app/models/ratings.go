package models

import (
	"gorm.io/gorm"
	"time"
)

type Rating struct {
	gorm.Model
	Score     int       `gorm:"type:int;not null"`
	Comment   string    `gorm:"type:text"`
	RatedAt   time.Time `gorm:"autoCreateTime"`
	PatientID uint      `gorm:"not null"`
	DoctorID  uint      `gorm:"not null"`
}
