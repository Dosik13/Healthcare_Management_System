package models

type Doctor struct {
	User
	Specialization   string          `gorm:"type:varchar(255);not null"`
	YearOfExperience uint            `gorm:"type:int;not null"`
	MoreInfo         string          `gorm:"type:varchar(255);not null"`
	Patients         []Patient       `gorm:"foreignKey:DoctorID"`
	Ratings          []Rating        `gorm:"foreignKey:DoctorID"`
	Appointments     []Appointment   `gorm:"foreignKey:DoctorID"`
	MedicalRecord    []MedicalRecord `gorm:"foreignKey:DoctorID"`
	HospitalID       *uint
}
