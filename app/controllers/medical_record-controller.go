package controllers

import (
	"Healthcare_Management_System/utils"
	"encoding/json"
	"html/template"
	"net/http"
	"strconv"
	"time"

	"Healthcare_Management_System/app/models" // Adjust the import path as necessary
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type MedicalRecordController struct {
	DB *gorm.DB
}

func NewMedicalRecordController(db *gorm.DB) *MedicalRecordController {
	return &MedicalRecordController{DB: db}
}

func (mrc *MedicalRecordController) CreateMedicalRecord(w http.ResponseWriter, r *http.Request) {
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

		tx := mrc.DB.Begin()

		patientID, _ := strconv.Atoi(r.FormValue("patientId"))
		recordType := r.FormValue("recordType")
		details := r.FormValue("details")
		date, _ := time.Parse("2006-01-02", r.FormValue("date"))

		medicalRecord := models.MedicalRecord{
			PatientID:  uint(patientID),
			DoctorID:   userId,
			RecordType: recordType,
			Details:    details,
			Date:       date,
		}

		err = tx.Create(&medicalRecord).Error
		if err != nil {
			tx.Rollback()
			http.Error(w, "Failed to create medical record", http.StatusInternalServerError)
			return
		}
		var prescription models.Prescription
		prescription.MedicationInfo = r.FormValue("medicationInfo")
		prescription.MedicalRecordID = medicalRecord.ID

		if err := tx.Create(&prescription).Error; err != nil {
			tx.Rollback()
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := tx.Commit().Error; err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("Medical record created successfully"))

	} else {
		tmpl, err := template.ParseFiles("frontend/templates/medical-record.html")
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

func (mrc *MedicalRecordController) GetMedicalRecord(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	recordID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid medical record ID", http.StatusBadRequest)
		return
	}

	var medicalRecord models.MedicalRecord
	if err := mrc.DB.Preload("Prescript").First(&medicalRecord, recordID).Error; err != nil {
		http.Error(w, "Medical record not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(medicalRecord)
}

func (mrc *MedicalRecordController) GetAllMedicalRecords(w http.ResponseWriter, r *http.Request) {
	var medicalRecords []models.MedicalRecord
	if err := mrc.DB.Preload("Prescript").Find(&medicalRecords).Error; err != nil {
		http.Error(w, "Failed to retrieve medical records", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(medicalRecords)
}

func (mrc *MedicalRecordController) UpdateMedicalRecord(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	recordID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid medical record ID", http.StatusBadRequest)
		return
	}

	var updateInfo models.MedicalRecord
	if err := json.NewDecoder(r.Body).Decode(&updateInfo); err != nil {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}

	if err := mrc.DB.Model(&models.MedicalRecord{}).Where("id = ?", recordID).Updates(updateInfo).Error; err != nil {
		http.Error(w, "Failed to update medical record", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Medical record updated successfully"))
}

func (mrc *MedicalRecordController) DeleteMedicalRecord(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	recordID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid medical record ID", http.StatusBadRequest)
		return
	}

	if err := mrc.DB.Delete(&models.MedicalRecord{}, recordID).Error; err != nil {
		http.Error(w, "Failed to delete medical record", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Medical record deleted successfully"))
}
