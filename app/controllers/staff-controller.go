package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"Healthcare_Management_System/app/models"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type HospitalStaffController struct {
	DB *gorm.DB
}

func NewHospitalStaffController(db *gorm.DB) *HospitalStaffController {
	return &HospitalStaffController{DB: db}
}

func (hsc *HospitalStaffController) CreateHospitalStaff(w http.ResponseWriter, r *http.Request) {
	var hospitalStaff models.HospitalStaff
	if err := json.NewDecoder(r.Body).Decode(&hospitalStaff); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	hsc.DB.Create(&hospitalStaff)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(hospitalStaff)
}

func (hsc *HospitalStaffController) GetHospitalStaff(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["staffID"])
	var hospitalStaff models.HospitalStaff
	if result := hsc.DB.First(&hospitalStaff, id); result.Error != nil {
		http.Error(w, "Hospital staff not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(hospitalStaff)
}

func (hsc *HospitalStaffController) GetAllHospitalStaff(w http.ResponseWriter, r *http.Request) {
	var hospitalStaffs []models.HospitalStaff
	hsc.DB.Find(&hospitalStaffs)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(hospitalStaffs)
}

func (hsc *HospitalStaffController) UpdateHospitalStaff(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["staffID"])
	var hospitalStaff models.HospitalStaff
	if result := hsc.DB.First(&hospitalStaff, id); result.Error != nil {
		http.Error(w, "Hospital staff not found", http.StatusNotFound)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&hospitalStaff); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	hsc.DB.Save(&hospitalStaff)
	json.NewEncoder(w).Encode(hospitalStaff)
}

func (hsc *HospitalStaffController) DeleteHospitalStaff(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["staffID"])
	if result := hsc.DB.Delete(&models.HospitalStaff{}, id); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
