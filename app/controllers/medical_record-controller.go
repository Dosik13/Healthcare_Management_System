package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"Healthcare_Management_System/app/models" // Adjust the import path as necessary
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type MedicalRecordController struct {
	DB *gorm.DB
}

func NewMedicalRecordController(db *gorm.DB) *MedicalRecordController {
	return &MedicalRecordController{DB: db}
}

func (mrc *MedicalRecordController) CreateMedicalRecord(w http.ResponseWriter, r *http.Request) {
	var medicalRecord models.MedicalRecord
	if err := json.NewDecoder(r.Body).Decode(&medicalRecord); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Handle creation with associated prescriptions if any
	mrc.DB.Create(&medicalRecord)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(medicalRecord)
}

func (mrc *MedicalRecordController) GetMedicalRecord(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid medical record ID", http.StatusBadRequest)
		return
	}
	var medicalRecord models.MedicalRecord
	if result := mrc.DB.Preload("Prescript").First(&medicalRecord, id); result.Error != nil {
		http.Error(w, "Medical record not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(medicalRecord)
}

func (mrc *MedicalRecordController) GetAllMedicalRecords(w http.ResponseWriter, r *http.Request) {
	var medicalRecords []models.MedicalRecord
	// Preload prescriptions for all medical records
	mrc.DB.Preload("Prescript").Find(&medicalRecords)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(medicalRecords)
}

func (mrc *MedicalRecordController) UpdateMedicalRecord(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid medical record ID", http.StatusBadRequest)
		return
	}
	var medicalRecord models.MedicalRecord
	if result := mrc.DB.First(&medicalRecord, id); result.Error != nil {
		http.Error(w, "Medical record not found", http.StatusNotFound)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&medicalRecord); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Update operation might need special handling for prescriptions
	mrc.DB.Save(&medicalRecord)
	json.NewEncoder(w).Encode(medicalRecord)
}

func (mrc *MedicalRecordController) DeleteMedicalRecord(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid medical record ID", http.StatusBadRequest)
		return
	}
	if result := mrc.DB.Delete(&models.MedicalRecord{}, id); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
