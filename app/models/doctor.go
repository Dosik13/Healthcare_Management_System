package models

type Doctor struct {
	User
	Specialization string `gorm:"not null"`
}
