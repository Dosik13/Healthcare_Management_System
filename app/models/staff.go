package models

type HospitalStaff struct {
	User
	StaffID uint   `gorm:"primaryKey"`
	Role    string `gorm:"type:varchar(255);not null"`
}
