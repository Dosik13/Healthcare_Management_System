package models

type Patient struct {
	User
	MedicalHistory string          `gorm:"type:text"`
	MedicalRecords []MedicalRecord `gorm:"foreignKey:PatientID"`
	Billing        []Billing       `gorm:"foreignKey:PatientID"`
	HospitalID     *uint
	DoctorID       *uint
	NurseID        *uint
}
