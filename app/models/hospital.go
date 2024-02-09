package models

import "time"

type Hospital struct {
	ID          uint      `gorm:"primaryKey"`
	Name        string    `gorm:"not null"`
	Address     string    `gorm:"not null"`
	PhoneNumber string    `gorm:"not null"`
	Description string    `gorm:"type:text"` // Optional
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}
