package models

type Nurse struct {
	User
	NurseID          uint      `gorm:"primaryKey"`
	HospitalID       uint      `gorm:"not null"`
	YearOfExperience uint      `gorm:"type:int;not null"`
	WorkHours        string    `gorm:"type:varchar(255);not null"`
	Patients         []Patient `gorm:"foreignKey:NurseID"`
}
