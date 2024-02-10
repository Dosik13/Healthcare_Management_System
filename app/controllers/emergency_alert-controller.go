package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"Healthcare_Management_System/app/models"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type EmergencyAlertController struct {
	DB *gorm.DB
}

func NewEmergencyAlertController(db *gorm.DB) *EmergencyAlertController {
	return &EmergencyAlertController{DB: db}
}

func (eac *EmergencyAlertController) CreateEmergencyAlert(w http.ResponseWriter, r *http.Request) {
	var emergencyAlert models.EmergencyAlert
	if err := json.NewDecoder(r.Body).Decode(&emergencyAlert); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	eac.DB.Create(&emergencyAlert)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(emergencyAlert)
}

func (eac *EmergencyAlertController) GetEmergencyAlert(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	var emergencyAlert models.EmergencyAlert
	eac.DB.First(&emergencyAlert, id)
	json.NewEncoder(w).Encode(emergencyAlert)
}

func (eac *EmergencyAlertController) GetAllEmergencyAlerts(w http.ResponseWriter, r *http.Request) {
	var emergencyAlerts []models.EmergencyAlert
	eac.DB.Find(&emergencyAlerts)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(emergencyAlerts)
}

func (eac *EmergencyAlertController) UpdateEmergencyAlert(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	var emergencyAlert models.EmergencyAlert
	eac.DB.First(&emergencyAlert, id)
	json.NewDecoder(r.Body).Decode(&emergencyAlert)
	eac.DB.Save(&emergencyAlert)
	json.NewEncoder(w).Encode(emergencyAlert)
}

func (eac *EmergencyAlertController) DeleteEmergencyAlert(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	eac.DB.Delete(&models.EmergencyAlert{}, id)
	w.WriteHeader(http.StatusNoContent)
}
