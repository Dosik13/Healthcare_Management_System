package routes

import (
	"Healthcare_Management_System/app/controllers"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// RegisterBillingRoutes sets up the router with routes related to billing management.
func RegisterBillingRoutes(router *mux.Router, db *gorm.DB) {
	billingController := controllers.NewBillingController(db)

	router.HandleFunc("/billings", billingController.CreateBilling).Methods("POST")
	router.HandleFunc("/billings", billingController.GetAllBillings).Methods("GET")
	router.HandleFunc("/billings/{billingID}", billingController.GetBilling).Methods("GET")
	router.HandleFunc("/billings/{billingID}", billingController.UpdateBilling).Methods("PUT")
	router.HandleFunc("/billings/{billingID}", billingController.DeleteBilling).Methods("DELETE")
}
