package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"Healthcare_Management_System/app/models"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type DoctorController struct {
	DB *gorm.DB
}

func NewDoctorController(db *gorm.DB) *DoctorController {
	return &DoctorController{DB: db}
}

func (dc *DoctorController) CreateDoctor(w http.ResponseWriter, r *http.Request) {
	var doctor models.Doctor
	if err := json.NewDecoder(r.Body).Decode(&doctor); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	dc.DB.Create(&doctor)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(doctor)
}

//func (dc *DoctorController) CreateDoctor(w http.ResponseWriter, r *http.Request) {
//	var doctor models.Doctor
//	if err := json.NewDecoder(r.Body).Decode(&doctor); err != nil {
//		http.Error(w, err.Error(), http.StatusBadRequest)
//		return
//	}
//
//	// Populate the Doctor fields
//	doctor.User = populateUser(r)
//	doctor.Specialization = r.Form.Get("specialization")
//	yearOfExperience, _ := strconv.Atoi(r.Form.Get("year_of_experience"))
//	doctor.YearOfExperience = uint(yearOfExperience)
//
//	// Save the Doctor to the database
//	dc.DB.Create(&doctor)
//
//	w.WriteHeader(http.StatusCreated)
//	json.NewEncoder(w).Encode(doctor)
//}

func (dc *DoctorController) GetDoctor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["doctorID"], 10, 32)
	if err != nil {
		http.Error(w, "Invalid doctor ID", http.StatusBadRequest)
		return
	}
	var doctor models.Doctor
	if result := dc.DB.Preload("Patients").Preload("Ratings").Preload("Appointments").First(&doctor, "doctor_id = ?", id); result.Error != nil {
		http.Error(w, "Doctor not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(doctor)
}

func (dc *DoctorController) GetAllDoctors(w http.ResponseWriter, r *http.Request) {
	var doctors []models.Doctor
	dc.DB.Preload("Patients").Preload("Ratings").Preload("Appointments").Find(&doctors) // Optionally preload patients if needed for the use case
	json.NewEncoder(w).Encode(doctors)
}

func (dc *DoctorController) UpdateDoctor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["doctorID"], 10, 32)
	if err != nil {
		http.Error(w, "Invalid doctor ID", http.StatusBadRequest)
		return
	}
	var doctor models.Doctor
	if result := dc.DB.First(&doctor, "doctor_id = ?", id); result.Error != nil {
		http.Error(w, "Doctor not found", http.StatusNotFound)
		return
	}
	var updates models.Doctor
	if err := json.NewDecoder(r.Body).Decode(&updates); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Assumes updates to existing fields; adjust as necessary for your application logic.
	dc.DB.Model(&doctor).Updates(updates)
	json.NewEncoder(w).Encode(doctor)
}

func (dc *DoctorController) DeleteDoctor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["doctorID"], 10, 32)
	if err != nil {
		http.Error(w, "Invalid doctor ID", http.StatusBadRequest)
		return
	}
	dc.DB.Delete(&models.Doctor{}, "doctor_id = ?", id)
	w.WriteHeader(http.StatusNoContent)
}
