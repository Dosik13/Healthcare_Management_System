package models

type Patient struct {
	User
	MedicalHistory string          `gorm:"type:text"`
	MedicalRecords []MedicalRecord `gorm:"foreignKey:PatientID"`
	Billings       []Billing       `gorm:"foreignKey:PatientID"`
	Appointments   []Appointment   `gorm:"foreignKey:PatientID"`
	HospitalID     *uint
	DoctorID       *uint
	NurseID        *uint
}
