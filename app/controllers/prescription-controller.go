package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"Healthcare_Management_System/app/models"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type PrescriptionController struct {
	DB *gorm.DB
}

func NewPrescriptionController(db *gorm.DB) *PrescriptionController {
	return &PrescriptionController{DB: db}
}

func (pc *PrescriptionController) CreatePrescription(w http.ResponseWriter, r *http.Request) {
	var prescription models.Prescription
	if err := json.NewDecoder(r.Body).Decode(&prescription); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	pc.DB.Create(&prescription)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(prescription)
}

func (pc *PrescriptionController) GetPrescription(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	var prescription models.Prescription
	if result := pc.DB.First(&prescription, id); result.Error != nil {
		http.Error(w, "Prescription not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(prescription)
}

func (pc *PrescriptionController) GetAllPrescriptions(w http.ResponseWriter, r *http.Request) {
	var prescriptions []models.Prescription
	pc.DB.Find(&prescriptions)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(prescriptions)
}

func (pc *PrescriptionController) UpdatePrescription(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	var prescription models.Prescription
	if result := pc.DB.First(&prescription, id); result.Error != nil {
		http.Error(w, "Prescription not found", http.StatusNotFound)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&prescription); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	pc.DB.Save(&prescription)
	json.NewEncoder(w).Encode(prescription)
}

func (pc *PrescriptionController) DeletePrescription(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	if result := pc.DB.Delete(&models.Prescription{}, id); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
