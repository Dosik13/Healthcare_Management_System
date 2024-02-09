package models

import "time"

type User struct {
	ID          uint      `gorm:"primaryKey"`
	UserID      string    `gorm:"uniqueIndex;not null"`
	Password    string    `gorm:"not null"`
	FirstName   string    `gorm:"not null"`
	MiddleName  string    `gorm:"not null"`
	LastName    string    `gorm:"not null"`
	Email       string    `gorm:"uniqueIndex;not null"`
	DateOfBirth time.Time `gorm:"not null"`
	Age         int       `gorm:"not null"`
	Address     string    `gorm:"not null"`
	PhoneNumber string    `gorm:"not null"`
	Gender      string    `gorm:"not null"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}
