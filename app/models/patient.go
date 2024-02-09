package models

type Patient struct {
	User
	MedicalHistory string `gorm:"type:text"`
	Disease        string
	//StartDate   string
	//DoctorID    uint
	//RoomID      uint
	//TreatmentID uint
	//Doctor      *Doctor
	//Room        *Room
	//Treatment   *Treatment
	//Numbers     *Number
	//Nurses      []*Nurse `gorm:"many2many:patient_nurses;"`
}
