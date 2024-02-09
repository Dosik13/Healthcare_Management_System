package models

type Administrator struct {
	ID        uint   `gorm:"primaryKey"`
	UserID    string `gorm:"uniqueIndex;type:varchar(255);not null"`
	Password  string `gorm:"type:varchar(255);not null"`
	FirstName string `gorm:"type:varchar(255);not null"`
	LastName  string `gorm:"type:varchar(255);not null"`
	Email     string `gorm:"uniqueIndex;type:varchar(255);not null"`
}
