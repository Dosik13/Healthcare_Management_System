package models

type Doctor struct {
	User
	DoctorID       uint   `gorm:"primaryKey"`
	Specialization string `gorm:"type:varchar(255);not null"`
	Role           string `gorm:"type:varchar(255);not null"`
}
