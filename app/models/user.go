package models

import "time"

type User struct {
	ID          uint      `gorm:"primaryKey"`
	UserID      string    `gorm:"uniqueIndex;type:varchar(255);not null"`
	Password    uint      `gorm:"type:varchar(255);not null"`
	FirstName   string    `gorm:"type:varchar(255);not null"`
	MiddleName  string    `gorm:"type:varchar(255);not null"`
	LastName    string    `gorm:"type:varchar(255);not null"`
	Email       string    `gorm:"uniqueIndex;type:varchar(255);not null"`
	DateOfBirth time.Time `gorm:"not null"`
	Age         int       `gorm:"not null"`
	Address     string    `gorm:"type:varchar(255);not null"`
	PhoneNumber string    `gorm:"type:varchar(255);not null"`
	Gender      string    `gorm:"type:varchar(255);not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
