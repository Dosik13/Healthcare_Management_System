package controllers

import (
	"Healthcare_Management_System/app/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"net/http"
	"strconv"
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

func (dc *DoctorController) GetDoctor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["doctorID"], 10, 32)
	if err != nil {
		http.Error(w, "Invalid doctor ID", http.StatusBadRequest)
		return
	}
	var doctor models.Doctor
	dc.DB.First(&doctor, id)
	json.NewEncoder(w).Encode(doctor)
}

func (dc *DoctorController) GetAllDoctors(w http.ResponseWriter, r *http.Request) {
	var doctors []models.Doctor
	dc.DB.Find(&doctors)
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
	dc.DB.First(&doctor, id)
	json.NewDecoder(r.Body).Decode(&doctor)
	dc.DB.Save(&doctor)
	json.NewEncoder(w).Encode(doctor)
}

func (dc *DoctorController) DeleteDoctor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["doctorID"], 10, 32)
	if err != nil {
		http.Error(w, "Invalid doctor ID", http.StatusBadRequest)
		return
	}
	dc.DB.Delete(&models.Doctor{}, id)
	w.WriteHeader(http.StatusNoContent)
}
