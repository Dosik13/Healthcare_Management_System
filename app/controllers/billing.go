package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"Healthcare_Management_System/app/models" // Adjust the import path to where your models are located
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type BillingController struct {
	DB *gorm.DB
}

// NewBillingController creates a new controller with database connection
func NewBillingController(db *gorm.DB) *BillingController {
	return &BillingController{DB: db}
}

// CreateBilling handles POST requests to add a new billing record
func (bc *BillingController) CreateBilling(w http.ResponseWriter, r *http.Request) {
	var billing models.Billing
	if err := json.NewDecoder(r.Body).Decode(&billing); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if result := bc.DB.Create(&billing); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(billing)
}

// GetBilling handles GET requests to retrieve a billing record by ID
func (bc *BillingController) GetBilling(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	billingID, err := strconv.Atoi(vars["billingID"])
	if err != nil {
		http.Error(w, "Invalid billing ID", http.StatusBadRequest)
		return
	}
	var billing models.Billing
	if result := bc.DB.First(&billing, billingID); result.Error != nil {
		http.Error(w, "Billing record not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(billing)
}

// GetAllBillings handles GET requests to retrieve all billing records
func (bc *BillingController) GetAllBillings(w http.ResponseWriter, r *http.Request) {
	var billings []models.Billing
	if result := bc.DB.Find(&billings); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(billings)
}

// UpdateBilling handles PUT requests to update an existing billing record
func (bc *BillingController) UpdateBilling(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	billingID, err := strconv.Atoi(vars["billingID"])
	if err != nil {
		http.Error(w, "Invalid billing ID", http.StatusBadRequest)
		return
	}
	var updatedBilling models.Billing
	if err := json.NewDecoder(r.Body).Decode(&updatedBilling); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var billing models.Billing
	if result := bc.DB.First(&billing, billingID); result.Error != nil {
		http.Error(w, "Billing record not found", http.StatusNotFound)
		return
	}
	bc.DB.Model(&billing).Updates(updatedBilling)
	json.NewEncoder(w).Encode(billing)
}

// DeleteBilling handles DELETE requests to remove a billing record
func (bc *BillingController) DeleteBilling(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	billingID, err := strconv.Atoi(vars["billingID"])
	if err != nil {
		http.Error(w, "Invalid billing ID", http.StatusBadRequest)
		return
	}
	if result := bc.DB.Delete(&models.Billing{}, billingID); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
