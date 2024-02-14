package controllers

import (
	"Healthcare_Management_System/app/models"
	"Healthcare_Management_System/utils"
	"encoding/json"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"html/template"
	"net/http"
	"strconv"
	"time"
)

type AppointmentController struct {
	DB *gorm.DB
}

func NewAppointmentController(db *gorm.DB) *AppointmentController {
	return &AppointmentController{DB: db}
}

func (ac *AppointmentController) CreateAppointment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	if r.Method == "POST" {

		session, err := utils.Store.Get(r, "SessionID")
		if err != nil {
			http.Error(w, "Session error", http.StatusInternalServerError)
			return
		}
		userId, ok := session.Values["user"].(uint)
		if !ok || userId == 0 {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		tx := ac.DB.Begin()

		var appointment models.Appointment
		appointment.DoctorID = userId
		appointment.StartTime = r.FormValue("startTime")
		appointment.EndTime = r.FormValue("endTime")
		appointment.Status = "Not scheduled"

		err = tx.Create(&appointment).Error
		if err != nil {
			tx.Rollback()
			http.Error(w, "Failed to create medical record", http.StatusInternalServerError)
			return
		}
		amount, _ := strconv.Atoi(r.FormValue("amount"))

		billing := models.Billing{
			AppointmentID: appointment.ID,
			Amount:        float64(amount),
			Status:        "",
		}

		if err := tx.Create(&billing).Error; err != nil {
			tx.Rollback()
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := tx.Commit().Error; err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/doctor_dashboard", http.StatusSeeOther)

	} else {
		tmpl, err := template.ParseFiles("frontend/templates/calendar.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func (ac *AppointmentController) GetNotScheduledAppointments(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html")

	session, _ := utils.Store.Get(r, "SessionID")
	patientID, ok := session.Values["patientID"].(uint)
	if !ok || patientID == 0 {
		http.Error(w, "Could not identify patient", http.StatusUnauthorized)
		return
	}

	var appointments []models.Appointment
	if result := ac.DB.Where("status = ?", "Not Scheduled").Find(&appointments); result.Error != nil {
		http.Error(w, "Failed to fetch appointments", http.StatusInternalServerError)
		return
	}

	if len(appointments) == 0 {
		w.Write([]byte("No 'Not Scheduled' appointments found"))
		return
	}

	//Tuka trqbva da davam na klienta HTML-a za calendar
	tmpl, err := template.ParseFiles("frontend/templates/calendar.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, appointments)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func (ac *AppointmentController) ScheduleAppointment(w http.ResponseWriter, r *http.Request) {
	session, _ := utils.Store.Get(r, "SessionID")
	patientID, ok := session.Values["user"].(uint)
	if !ok || patientID == 0 {
		http.Error(w, "Could not identify patient", http.StatusUnauthorized)
		return
	}

	appointmentID, err := strconv.Atoi(r.FormValue("appointmentID"))
	if err != nil {
		http.Error(w, "Invalid appointment ID", http.StatusBadRequest)
		return
	}

	var appointment models.Appointment
	if err := ac.DB.Where("id = ? AND status = ?", appointmentID, "Not Scheduled").First(&appointment).Error; err != nil {
		http.Error(w, "Appointment not found or already scheduled", http.StatusNotFound)
		return
	}

	tx := ac.DB.Begin()

	// Update the appointment with the patient's ID and change its status
	updateResult := tx.Model(&appointment).Updates(models.Appointment{PatientID: patientID, Status: "Scheduled"})
	if updateResult.Error != nil {
		tx.Rollback()
		http.Error(w, "Failed to schedule appointment", http.StatusInternalServerError)
		return
	}

	var billing models.Billing
	billingResult := tx.Where("appointment_id = ?", appointment.ID).FirstOrCreate(&billing, models.Billing{
		PatientID: patientID,
		Status:    "Pending",
		DueDate:   time.Now().Add(30 * 24 * time.Hour), // Example: due date 30 days from now
	})
	if billingResult.Error != nil {
		tx.Rollback()
		http.Error(w, "Failed to update/create billing record", http.StatusInternalServerError)
		return
	}

	if err := tx.Commit().Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

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
