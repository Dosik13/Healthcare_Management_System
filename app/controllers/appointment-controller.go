package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"Healthcare_Management_System/app/models"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type AppointmentController struct {
	DB *gorm.DB
}

// NewAppointmentController creates a new AppointmentController with database connection
func NewAppointmentController(db *gorm.DB) *AppointmentController {
	return &AppointmentController{DB: db}
}

// CreateAppointment handles POST requests to add a new appointment
func (ac *AppointmentController) CreateAppointment(w http.ResponseWriter, r *http.Request) {
	var appointment models.Appointment
	if err := json.NewDecoder(r.Body).Decode(&appointment); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if result := ac.DB.Create(&appointment); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(appointment)
}

// GetAppointment handles GET requests to retrieve an appointment by ID
func (ac *AppointmentController) GetAppointment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	appointmentID, err := strconv.Atoi(vars["appointmentID"])
	if err != nil {
		http.Error(w, "Invalid appointment ID", http.StatusBadRequest)
		return
	}
	var appointment models.Appointment
	if result := ac.DB.First(&appointment, appointmentID); result.Error != nil {
		http.Error(w, "Appointment not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(appointment)
}

// GetAllAppointments handles GET requests to retrieve all appointments
func (ac *AppointmentController) GetAllAppointments(w http.ResponseWriter, r *http.Request) {
	var appointments []models.Appointment
	if result := ac.DB.Find(&appointments); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(appointments)
}

// UpdateAppointment handles PUT requests to update an existing appointment
func (ac *AppointmentController) UpdateAppointment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	appointmentID, err := strconv.Atoi(vars["appointmentID"])
	if err != nil {
		http.Error(w, "Invalid appointment ID", http.StatusBadRequest)
		return
	}
	var updatedAppointment models.Appointment
	if err := json.NewDecoder(r.Body).Decode(&updatedAppointment); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var appointment models.Appointment
	if result := ac.DB.First(&appointment, appointmentID); result.Error != nil {
		http.Error(w, "Appointment not found", http.StatusNotFound)
		return
	}
	ac.DB.Model(&appointment).Updates(updatedAppointment)
	json.NewEncoder(w).Encode(appointment)
}

// DeleteAppointment handles DELETE requests to remove an appointment
func (ac *AppointmentController) DeleteAppointment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	appointmentID, err := strconv.Atoi(vars["appointmentID"])
	if err != nil {
		http.Error(w, "Invalid appointment ID", http.StatusBadRequest)
		return
	}
	if result := ac.DB.Delete(&models.Appointment{}, appointmentID); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
