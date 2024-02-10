package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"Healthcare_Management_System/app/models"
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
	id, _ := strconv.Atoi(vars["nurseID"])
	var nurse models.Nurse
	if result := nc.DB.First(&nurse, id); result.Error != nil {
		http.Error(w, "Nurse not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(nurse)
}

func (nc *NurseController) GetAllNurses(w http.ResponseWriter, r *http.Request) {
	var nurses []models.Nurse
	nc.DB.Find(&nurses)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(nurses)
}

func (nc *NurseController) UpdateNurse(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["nurseID"])
	var nurse models.Nurse
	if result := nc.DB.First(&nurse, id); result.Error != nil {
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
	id, _ := strconv.Atoi(vars["nurseID"])
	if result := nc.DB.Delete(&models.Nurse{}, id); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
