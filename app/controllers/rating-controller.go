package controllers

import (
	"Healthcare_Management_System/utils"
	"encoding/json"
	"html/template"
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
	w.Header().Set("Content-Type", "text/html")

	vars := mux.Vars(r)
	doctorId, _ := strconv.Atoi(vars["doctorID"])

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
		score, _ := strconv.Atoi(r.FormValue("score"))

		rating := models.Rating{
			DoctorID:  uint(doctorId),
			PatientID: userId,
			Score:     score,
			Comment:   r.FormValue("comment"),
		}

		if result := rc.DB.Create(&rating); result.Error != nil {
			http.Error(w, result.Error.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/patient_dashboard", http.StatusSeeOther)

	} else {
		tmpl, err := template.ParseFiles("frontend/templates/rating.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		data := map[string]interface{}{
			"DoctorID": doctorId,
		}

		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
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
