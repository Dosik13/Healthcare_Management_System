package models

import "time"

type Hospital struct {
	HospitalID     uint             `gorm:"primaryKey"`
	Name           string           `gorm:"type:varchar(255);not null"`
	Address        string           `gorm:"type:varchar(255);not null"`
	PhoneNumber    string           `gorm:"type:varchar(255);not null"`
	Description    string           `gorm:"type:text"` // Optional
	Doctors        []Doctor         `gorm:"foreignKey:HospitalID"`
	Nurses         []Nurse          `gorm:"foreignKey:HospitalID"`
	Patients       []Patient        `gorm:"foreignKey:HospitalID"`
	EmergencyAlert []EmergencyAlert `gorm:"foreignKey:HospitalID"`
	Appointments   []Appointment    `gorm:"foreignKey:HospitalID"`
	CreatedAt      time.Time        `gorm:"autoCreateTime"`
	UpdatedAt      time.Time        `gorm:"autoUpdateTime"`
}
