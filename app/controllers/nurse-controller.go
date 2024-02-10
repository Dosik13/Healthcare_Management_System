package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"Healthcare_Management_System/app/models" // Update the import path as necessary
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type NurseController struct {
	DB *gorm.DB
}

func NewNurseController(db *gorm.DB) *NurseController {
	return &NurseController{DB: db}
}

func (nc *NurseController) CreateNurse(w http.ResponseWriter, r *http.Request) {
	var nurse models.Nurse
	if err := json.NewDecoder(r.Body).Decode(&nurse); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	nc.DB.Create(&nurse)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(nurse)
}

func (nc *NurseController) GetNurse(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["nurseID"], 10, 32)
	if err != nil {
		http.Error(w, "Invalid nurse ID", http.StatusBadRequest)
		return
	}
	var nurse models.Nurse
	if result := nc.DB.Preload("Patients").First(&nurse, "nurse_id = ?", id); result.Error != nil {
		http.Error(w, "Nurse not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(nurse)
}

func (nc *NurseController) GetAllNurses(w http.ResponseWriter, r *http.Request) {
	var nurses []models.Nurse
	nc.DB.Preload("Patients").Find(&nurses)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(nurses)
}

func (nc *NurseController) UpdateNurse(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["nurseID"], 10, 32)
	if err != nil {
		http.Error(w, "Invalid nurse ID", http.StatusBadRequest)
		return
	}
	var nurse models.Nurse
	if result := nc.DB.First(&nurse, "nurse_id = ?", id); result.Error != nil {
		http.Error(w, "Nurse not found", http.StatusNotFound)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&nurse); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	nc.DB.Save(&nurse)
	json.NewEncoder(w).Encode(nurse)
}

func (nc *NurseController) DeleteNurse(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["nurseID"], 10, 32)
	if err != nil {
		http.Error(w, "Invalid nurse ID", http.StatusBadRequest)
		return
	}
	nc.DB.Delete(&models.Nurse{}, "nurse_id = ?", id)
	w.WriteHeader(http.StatusNoContent)
}
