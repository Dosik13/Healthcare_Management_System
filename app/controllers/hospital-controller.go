package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"Healthcare_Management_System/app/models"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type HospitalController struct {
	DB *gorm.DB
}

func NewHospitalController(db *gorm.DB) *HospitalController {
	return &HospitalController{DB: db}
}

func (hc *HospitalController) CreateHospital(w http.ResponseWriter, r *http.Request) {
	var hospital models.Hospital
	if err := json.NewDecoder(r.Body).Decode(&hospital); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if result := hc.DB.Create(&hospital); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(hospital)
}

func (hc *HospitalController) GetHospital(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	hospitalID, err := strconv.ParseUint(vars["hospitalID"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid hospital ID", http.StatusBadRequest)
		return
	}
	var hospital models.Hospital
	if result := hc.DB.Preload("Doctors").Preload("Nurses").Preload("Patients").Preload("EmergencyAlerts").Preload("Appointments").First(&hospital, "hospital_id = ?", hospitalID); result.Error != nil {
		http.Error(w, "Hospital not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(hospital)
}

func (hc *HospitalController) GetAllHospitals(w http.ResponseWriter, r *http.Request) {
	var hospitals []models.Hospital
	if result := hc.DB.Preload("Doctors").Preload("Nurses").Preload("Patients").Preload("EmergencyAlerts").Preload("Appointments").Find(&hospitals); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(hospitals)
}

func (hc *HospitalController) UpdateHospital(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	hospitalID, err := strconv.ParseUint(vars["hospitalID"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid hospital ID", http.StatusBadRequest)
		return
	}
	var hospitalUpdates models.Hospital
	if err := json.NewDecoder(r.Body).Decode(&hospitalUpdates); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if result := hc.DB.Model(&models.Hospital{}).Where("hospital_id = ?", hospitalID).Updates(hospitalUpdates); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Hospital updated successfully")
}

func (hc *HospitalController) DeleteHospital(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	hospitalID, err := strconv.ParseUint(vars["hospitalID"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid hospital ID", http.StatusBadRequest)
		return
	}
	if result := hc.DB.Delete(&models.Hospital{}, hospitalID); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
