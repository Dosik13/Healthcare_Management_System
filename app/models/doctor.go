package models

type Doctor struct {
	User
	DoctorID         uint      `gorm:"primaryKey"`
	Specialization   string    `gorm:"type:varchar(255);not null"`
	YearOfExperience uint      `gorm:"type:int;not null"`
	WorkHours        string    `gorm:"type:varchar(255);not null"`
	Patients         []Patient `gorm:"foreignKey:DoctorID"`
}
