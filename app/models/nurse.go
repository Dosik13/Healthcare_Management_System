package models

type Nurse struct {
	User
	NurseID        uint   `gorm:"primaryKey"`
	Specialization string `gorm:"type:varchar(255);not null"`
	Role           string `gorm:"type:varchar(255);not null"`
}
