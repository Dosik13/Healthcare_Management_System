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
		appointment.Date = r.FormValue("date")
		appointment.Status = "Not scheduled"

		err = tx.Create(&appointment).Error
		if err != nil {
			tx.Rollback()
			http.Error(w, "Failed to create medical record", http.StatusInternalServerError)
			return
		}
		//amount, _ := strconv.Atoi(r.FormValue("amount"))

		billing := models.Billing{
			AppointmentID: appointment.ID,
			Amount:        float64(10),
			Status:        "",
			DueDate:       time.Now().Add(30 * 24 * time.Hour).Format("2006-01-02"),
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

		http.Redirect(w, r, "/appointments", http.StatusSeeOther)
	}
}
func (ac *AppointmentController) GetNotScheduledAppointments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	session, _ := utils.Store.Get(r, "SessionID")
	patientID, ok := session.Values["user"].(uint)
	if !ok || patientID == 0 {
		http.Error(w, "Could not identify patient", http.StatusUnauthorized)
		return
	}

	vars := mux.Vars(r)
	doctorId, err := strconv.Atoi(vars["doctorID"])
	if err != nil {
		http.Error(w, "Invalid appointment ID", http.StatusBadRequest)
		return
	}

	var doctor models.Doctor

	subQuery := ac.DB.Table("appointments").
		Select("doctor_id").
		Where("status = ?", "Not scheduled").
		Group("doctor_id")

	result := ac.DB.
		Where("user_id = ?", doctorId).
		Where("user_id IN (?)", subQuery).
		Preload("Appointments").
		Preload("Ratings").
		Preload("Patients").
		First(&doctor)

	if result.Error != nil {
		http.Error(w, "Doctor not found or no 'Not scheduled' appointments", http.StatusNotFound)
		return
	}

	if len(doctor.Appointments) == 0 {
		w.Write([]byte("No 'Not Scheduled' appointments found"))
		return
	}

	tmpl, err := template.ParseFiles("frontend/templates/patientCalendar.html")
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}

	mp := make(map[string][]string)

	for _, appointment := range doctor.Appointments {
		mp[appointment.Date] = append(mp[appointment.Date], appointment.StartTime+"-"+appointment.EndTime)
	}

	appointmentsJSON, err := json.Marshal(mp)
	if err != nil {
		http.Error(w, "Failed to serialize appointments", http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"appointmentsJS": string(appointmentsJSON),
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

type AppointmentData struct {
	Start   string `json:"start"`
	End     string `json:"end"`
	Date    string `json:"date"`
	AddInfo string `json:"addInfo"`
}

func (ac *AppointmentController) ScheduleAppointment(w http.ResponseWriter, r *http.Request) {

	session, _ := utils.Store.Get(r, "SessionID")
	patientID, ok := session.Values["user"].(uint)
	if !ok || patientID == 0 {
		http.Error(w, "Could not identify patient", http.StatusUnauthorized)
		return
	}
	if r.Method == "POST" {

		var app AppointmentData

		err := json.NewDecoder(r.Body).Decode(&app)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		var appointment models.Appointment
		if err := ac.DB.Where("date = ? AND start_time = ? AND end_time = ?", app.Date, app.Start, app.End).First(&appointment).Error; err != nil {
			http.Error(w, "Appointment not found or already scheduled", http.StatusNotFound)
			return
		}

		tx := ac.DB.Begin()

		updateResult := tx.Model(&appointment).Updates(models.Appointment{PatientID: &patientID, Status: "Scheduled", Reason: app.AddInfo})
		if updateResult.Error != nil {
			tx.Rollback()
			http.Error(w, "Failed to schedule appointment", http.StatusInternalServerError)
			return
		}

		var billing models.Billing

		if result := tx.Where("appointment_id = ?", appointment.ID).First(&billing); result.Error != nil {
			tx.Rollback()
			http.Error(w, "Failed to update/create billing record", http.StatusInternalServerError)
			return
		}

		updateResult = tx.Model(&billing).Updates(models.Billing{PatientID: &patientID, Status: "Pending", DueDate: time.Now().Add(30 * 24 * time.Hour).Format("2006-01-02")})
		if updateResult.Error != nil {
			tx.Rollback()
			http.Error(w, "Failed to schedule appointment", http.StatusInternalServerError)
			return
		}
		var doctor models.Doctor

		if err := tx.Preload("Patients").First(&doctor, "id = ?", appointment.DoctorID).Error; err != nil {
			tx.Rollback()
			http.Error(w, "Failed to fetch doctor", http.StatusInternalServerError)
			return
		}

		var patient models.Patient
		if err := tx.First(&patient, "id = ?", patientID).Error; err != nil {
			tx.Rollback()
			http.Error(w, "Failed to fetch patient", http.StatusInternalServerError)
			return
		}

		doctor.Patients = append(doctor.Patients, patient)

		if err := tx.Save(&doctor).Error; err != nil {
			tx.Rollback()
			http.Error(w, "Failed to update doctor's patients", http.StatusInternalServerError)
			return
		}

		if err := tx.Commit().Error; err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/patient_dashboard", http.StatusSeeOther)
	}
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
	session, _ := utils.Store.Get(r, "SessionID")
	userId, ok := session.Values["user"].(uint)
	if !ok || userId == 0 {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	var doctor models.Doctor

	if result := ac.DB.Preload("Appointments").Preload("Ratings").Preload("Patients").First(&doctor, "user_id = ?", userId); result.Error != nil {
		http.Error(w, "Doctor not found", http.StatusNotFound)
		return
	}

	tmpl, err := template.ParseFiles("frontend/templates/calendar.html")
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}

	mp := make(map[string][]string)

	for _, appointment := range doctor.Appointments {
		mp[appointment.Date] = append(mp[appointment.Date], appointment.StartTime+"-"+appointment.EndTime)
	}

	appointmentsJSON, err := json.Marshal(mp)
	if err != nil {
		http.Error(w, "Failed to serialize appointments", http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"appointmentsJS": string(appointmentsJSON),
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
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
