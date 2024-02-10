package models

type Patient struct {
	User
	MedicalHistory string `gorm:"type:text"`
	Disease        string `gorm:"type:text"`
	Role           string `gorm:"type:varchar(255);not null"`
}
