package models

type Patient struct {
	User
	PatientID      uint            `gorm:"primaryKey"`
	MedicalHistory string          `gorm:"type:text"`
	Allergies      string          `gorm:"type:text"`
	MedicalRecords []MedicalRecord `gorm:"foreignKey:PatientID"`
	Billing        []Billing       `gorm:"foreignKey:PatientID"`
	HospitalID     uint            `gorm:"not null"`
	DoctorID       uint
	NurseID        uint
}
