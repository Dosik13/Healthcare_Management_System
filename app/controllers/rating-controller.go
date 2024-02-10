package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"Healthcare_Management_System/app/models"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type RatingController struct {
	DB *gorm.DB
}

func NewRatingController(db *gorm.DB) *RatingController {
	return &RatingController{DB: db}
}

func (rc *RatingController) CreateRating(w http.ResponseWriter, r *http.Request) {
	var rating models.Rating
	if err := json.NewDecoder(r.Body).Decode(&rating); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	rc.DB.Create(&rating)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(rating)
}

func (rc *RatingController) GetRating(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	var rating models.Rating
	if result := rc.DB.First(&rating, id); result.Error != nil {
		http.Error(w, "Rating not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(rating)
}

func (rc *RatingController) GetAllRatings(w http.ResponseWriter, r *http.Request) {
	var ratings []models.Rating
	rc.DB.Find(&ratings)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ratings)
}

func (rc *RatingController) UpdateRating(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	var rating models.Rating
	if result := rc.DB.First(&rating, id); result.Error != nil {
		http.Error(w, "Rating not found", http.StatusNotFound)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&rating); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	rc.DB.Save(&rating)
	json.NewEncoder(w).Encode(rating)
}

func (rc *RatingController) DeleteRating(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	if result := rc.DB.Delete(&models.Rating{}, id); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
