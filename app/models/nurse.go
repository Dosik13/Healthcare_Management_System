package models

type Nurse struct {
	User
	Department string `gorm:"not null"`
	//	Patients []*Patient `gorm:"many2many:patient_nurses;"`
}
