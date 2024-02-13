package models

import "gorm.io/gorm"

type Doctor struct {
	gorm.Model
	User
	Specialization   string        `gorm:"type:varchar(255);not null"`
	YearOfExperience uint          `gorm:"type:int;not null"`
	WorkHours        string        `gorm:"type:varchar(255);not null"`
	Patients         []Patient     `gorm:"foreignKey:DoctorID"`
	Ratings          []Rating      `gorm:"foreignKey:DoctorID"`
	Appointments     []Appointment `gorm:"foreignKey:DoctorID"`
	HospitalID       uint          `gorm:"not null"`
}
