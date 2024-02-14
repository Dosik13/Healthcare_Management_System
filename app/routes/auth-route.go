package routes

import (
	"Healthcare_Management_System/app/controllers"
	"Healthcare_Management_System/utils"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"net/http"
)

func RegisterAuth(router *mux.Router, db *gorm.DB) {
	AuthHandler := controllers.NewAuthController(db)

	router.HandleFunc("/login", AuthHandler.LoginHandler).Methods("GET")
	router.HandleFunc("/login", AuthHandler.LoginHandler).Methods("POST")
	router.HandleFunc("/register", AuthHandler.RegisterHandler).Methods("GET")
	router.HandleFunc("/register", AuthHandler.RegisterHandler).Methods("POST")
	router.HandleFunc("/logout", AuthHandler.LogoutHandler).Methods("GET")
	router.HandleFunc("/home", AuthHandler.HomeHandler).Methods("GET")
	router.Handle("/doctor_dashboard", utils.AuthDoctorHandler(http.HandlerFunc(AuthHandler.DoctorDashboardHandler))).Methods("GET")
	//router.Handle("/nurse/dashboard", utils.RoleAccessMiddleware("nurse", nurseDashboardHandler)).Methods("GET")
	//router.Handle("/admin/dashboard", utils.RoleAccessMiddleware("administrator", adminDashboardHandler)).Methods("GET")

}
