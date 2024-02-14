package models

type Nurse struct {
	User
	HospitalID       uint      `gorm:"not null"`
	YearOfExperience uint      `gorm:"type:int;not null"`
	MoreInfo         string    `gorm:"type:varchar(255);not null"`
	Patients         []Patient `gorm:"foreignKey:NurseID"`
}
