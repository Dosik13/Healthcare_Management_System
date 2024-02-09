package controllers

import (
	"Healthcare_Management_System/app/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"net/http"
)

type UserController struct {
	DB *gorm.DB
}

// NewUserController creates a new UserController with database connection
func NewUserController(db *gorm.DB) *UserController {
	return &UserController{DB: db}
}

// CreateUser handles POST requests to add a new user
func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if result := uc.DB.Create(&user); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// GetUser handles GET requests to retrieve a user by UserID
func (uc *UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["userID"]
	var user models.User
	if result := uc.DB.Where("user_id = ?", userID).First(&user); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(user)
}

// GetAllUsers handles GET requests to retrieve all users
func (uc *UserController) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	if result := uc.DB.Find(&users); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

// UpdateUser handles PUT requests to update an existing user
func (uc *UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["userID"]
	var updatedUser models.User
	if err := json.NewDecoder(r.Body).Decode(&updatedUser); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var user models.User
	if result := uc.DB.Where("user_id = ?", userID).First(&user); result.Error != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	uc.DB.Model(&user).Updates(updatedUser)
	json.NewEncoder(w).Encode(user)
}

// DeleteUser handles DELETE requests to remove a user
func (uc *UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["userID"]
	if result := uc.DB.Where("user_id = ?", userID).Delete(&models.User{}); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
