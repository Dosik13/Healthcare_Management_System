package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"Healthcare_Management_System/app/models"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type PatientController struct {
	DB *gorm.DB
}

func NewPatientController(db *gorm.DB) *PatientController {
	return &PatientController{DB: db}
}

func (pc *PatientController) CreatePatient(w http.ResponseWriter, r *http.Request) {
	var patient models.Patient
	if err := json.NewDecoder(r.Body).Decode(&patient); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	pc.DB.Create(&patient)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(patient)
}

func (pc *PatientController) GetPatient(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["patientID"])
	var patient models.Patient
	if result := pc.DB.Preload("MedicalRecords").First(&patient, id); result.Error != nil {
		http.Error(w, "Patient not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(patient)
}

func (pc *PatientController) GetAllPatients(w http.ResponseWriter, r *http.Request) {
	var patients []models.Patient
	pc.DB.Preload("MedicalRecords").Find(&patients)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(patients)
}

func (pc *PatientController) UpdatePatient(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["patientID"])
	var patient models.Patient
	if result := pc.DB.First(&patient, id); result.Error != nil {
		http.Error(w, "Patient not found", http.StatusNotFound)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&patient); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	pc.DB.Save(&patient)
	json.NewEncoder(w).Encode(patient)
}

func (pc *PatientController) DeletePatient(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["patientID"])
	if result := pc.DB.Delete(&models.Patient{}, id); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
