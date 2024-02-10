package models

type Administrator struct {
	User
	Role string `gorm:"type:varchar(255);not null"`
}
