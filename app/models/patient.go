package models

type Patient struct {
	User
	PatientID      uint            `gorm:"primaryKey"`
	MedicalHistory string          `gorm:"type:text"`
	Allergies      string          `gorm:"type:text"`
	MedicalRecords []MedicalRecord `gorm:"foreignKey:PatientID"`
}
