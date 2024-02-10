package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"Healthcare_Management_System/app/models"
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
	mrc.DB.Create(&medicalRecord)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(medicalRecord)
}

func (mrc *MedicalRecordController) GetMedicalRecord(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	var medicalRecord models.MedicalRecord
	mrc.DB.First(&medicalRecord, id)
	json.NewEncoder(w).Encode(medicalRecord)
}

func (mrc *MedicalRecordController) GetAllMedicalRecords(w http.ResponseWriter, r *http.Request) {
	var medicalRecords []models.MedicalRecord
	mrc.DB.Find(&medicalRecords)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(medicalRecords)
}

func (mrc *MedicalRecordController) UpdateMedicalRecord(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	var medicalRecord models.MedicalRecord
	mrc.DB.First(&medicalRecord, id)
	json.NewDecoder(r.Body).Decode(&medicalRecord)
	mrc.DB.Save(&medicalRecord)
	json.NewEncoder(w).Encode(medicalRecord)
}

func (mrc *MedicalRecordController) DeleteMedicalRecord(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	mrc.DB.Delete(&models.MedicalRecord{}, id)
	w.WriteHeader(http.StatusNoContent)
}
