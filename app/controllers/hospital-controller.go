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
	hc.DB.Create(&hospital)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(hospital)
}

func (hc *HospitalController) GetHospital(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	var hospital models.Hospital
	hc.DB.First(&hospital, id)
	json.NewEncoder(w).Encode(hospital)
}

func (hc *HospitalController) GetAllHospitals(w http.ResponseWriter, r *http.Request) {
	var hospitals []models.Hospital
	hc.DB.Find(&hospitals)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(hospitals)
}

func (hc *HospitalController) UpdateHospital(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	var hospital models.Hospital
	hc.DB.First(&hospital, id)
	json.NewDecoder(r.Body).Decode(&hospital)
	hc.DB.Save(&hospital)
	json.NewEncoder(w).Encode(hospital)
}

func (hc *HospitalController) DeleteHospital(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	hc.DB.Delete(&models.Hospital{}, id)
	w.WriteHeader(http.StatusNoContent)
}
