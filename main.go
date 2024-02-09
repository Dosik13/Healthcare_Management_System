package main

import (
	"Healthcare_Management_System/app/models"
	"Healthcare_Management_System/app/routes"
	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
)

func main() {
	db, err := gorm.Open(mysql.Open("test.db"), &gorm.Config{}) // Adjust for your DB setup
	if err != nil {
		panic("failed to connect database")
	}
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		return
	} // Auto migrate your models

	r := mux.NewRouter()

	routes.RegisterUserRoutes(r, db)

	err = http.ListenAndServe(":8080", r)
	if err != nil {
		return
	}
}
