package models

type Doctor struct {
	User
	DoctorID       uint   `gorm:"primaryKey"`
	Specialization string `gorm:"not null"`
}
