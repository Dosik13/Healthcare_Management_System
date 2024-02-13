package models

import (
	"time"
)

//type Role string
//
//const (
//	A Role = "Administator"
//	P Role = "Patient"
//	D Role = "Doctor"
//	N Role = "Nurse"
//)

type User struct {
	//gorm.Model
	UserID      uint      `gorm:"primaryKey"`
	FirstName   string    `gorm:"type:varchar(255);not null"`
	MiddleName  string    `gorm:"type:varchar(255);not null"`
	LastName    string    `gorm:"type:varchar(255);not null"`
	Email       string    `gorm:"uniqueIndex;type:varchar(255);not null"`
	Password    string    `gorm:"type:varchar(255);not null"`
	UCN         string    `gorm:"type:varchar(255);not null"`
	Address     string    `gorm:"type:varchar(255);not null"`
	PhoneNumber string    `gorm:"type:varchar(255);not null"`
	Gender      string    `gorm:"type:varchar(255);not null"`
	Role        string    `gorm:"type:varchar(255);not null"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}
