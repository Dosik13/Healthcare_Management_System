package controllers

import (
	"Healthcare_Management_System/app/models"
	"Healthcare_Management_System/utils"
	"fmt"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type RoleUser struct {
	Email    string
	Password string
	UserId   uint
}

type AuthController struct {
	DB *gorm.DB
}

func NewAuthController(db *gorm.DB) *AuthController {
	return &AuthController{DB: db}
}
func (ac *AuthController) LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// Parse form data
		err := r.ParseMultipartForm(10 << 20) // Parse up to 10MB of memory
		if err != nil {
			http.Error(w, "Error parsing form", http.StatusBadRequest)
			return
		}

		email := r.Form.Get("email")
		password := r.Form.Get("password")
		role := r.Form.Get("role") // Retrieve the role from the form

		fmt.Println("Email: ", email)
		fmt.Println("Password: ", password)
		fmt.Println("Role: ", role)

		var user RoleUser

		switch role {
		case "Doctor":
			var doctor models.Doctor
			if err := ac.DB.Where("email = ?", email).First(&doctor).Error; err == nil {
				user = RoleUser{Email: doctor.Email, Password: doctor.Password, UserId: doctor.UserID}
			}
		case "Nurse":
			var nurse models.Nurse
			if err := ac.DB.Where("email = ?", email).First(&nurse).Error; err == nil {
				user = RoleUser{Email: nurse.Email, Password: nurse.Password, UserId: nurse.UserID}
			}
		case "Patient":
			var patient models.Patient
			if err := ac.DB.Where("email = ?", email).First(&patient).Error; err == nil {
				user = RoleUser{Email: patient.Email, Password: patient.Password, UserId: patient.UserID}
			}
		default:
			http.Error(w, "Invalid role specified", http.StatusBadRequest)
			return
		}

		if utils.CheckPasswordHash(password, user.Password) != true {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}
		// Create session
		session, _ := utils.Store.Get(r, "session-name")
		// Depending on your session library, you might need to cast user.ID to the appropriate type
		session.Values["user"] = user.UserId // Store user ID in session
		session.Values["role"] = role        // Store role in session
		session.Save(r, w)

		// Redirect or respond to indicate success
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
	}
}

func (ac *AuthController) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if r.Method == "POST" {

		err := r.ParseMultipartForm(10 << 20) // Parse up to 10MB of memory
		if err != nil {
			http.Error(w, "Error parsing form", http.StatusBadRequest)
			return
		}

		role := r.Form.Get("role")

		fmt.Println("Role: ", role)

		switch role {
		case "Doctor":
			var doctor models.Doctor
			var doctor2 models.Doctor
			doctor.User = populateUser(r)
			doctor.Specialization = r.Form.Get("specialization")
			yearOfExperience, _ := strconv.Atoi(r.Form.Get("year_of_experience"))
			doctor.YearOfExperience = uint(yearOfExperience)

			if result := ac.DB.First(&doctor2, "email = ?", doctor.Email); result.Error != nil {
				http.Error(w, "Doctor already exists!", http.StatusConflict)
				return
			} else {
				result := ac.DB.Create(&doctor)
				if result.Error != nil {
					fmt.Fprintf(w, "Error registering doctor: %s", result.Error)
					return
				}
				w.WriteHeader(http.StatusCreated)
			}

		case "Patient":
			var patient models.Patient
			var patient2 models.Patient
			patient.User = populateUser(r)
			patient.Allergies = r.Form.Get("allergies")
			patient.MedicalHistory = r.Form.Get("medical_history")
			if result := ac.DB.First(&patient2, "email = ?", patient.Email); result.Error != nil {
				result := ac.DB.Create(&patient)
				if result.Error != nil {
					fmt.Fprintf(w, "Error registering patient: %s", result.Error)
					return
				}
				w.WriteHeader(http.StatusCreated)
			} else {
				http.Error(w, "Patient already exists!", http.StatusConflict)
				return
			}

		case "Nurse":
			var nurse models.Nurse
			var nurse2 models.Nurse
			nurse.User = populateUser(r)
			yearOfExperience, _ := strconv.Atoi(r.Form.Get("year_of_experience"))
			nurse.YearOfExperience = uint(yearOfExperience)

			if result := ac.DB.First(&nurse2, "email = ?", nurse.Email); result.Error != nil {
				http.Error(w, "Doctor already exists!", http.StatusConflict)
				return
			} else {
				result := ac.DB.Create(&nurse)
				if result.Error != nil {
					fmt.Fprintf(w, "Error registering doctor: %s", result.Error)
					return
				}
				w.WriteHeader(http.StatusCreated)
			}
		}

	} else {
		// Render the registration form

	}
}

func (ac *AuthController) LogoutHandler(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:   "session_token",
		Value:  "",
		MaxAge: -1, // This will delete the cookie
	})
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func populateUser(r *http.Request) models.User {
	return models.User{
		FirstName:   r.Form.Get("first_name"),
		MiddleName:  r.Form.Get("middle_name"),
		LastName:    r.Form.Get("last_name"),
		Email:       r.Form.Get("email"),
		Password:    utils.HashPassword(r.Form.Get("password")),
		UCN:         r.Form.Get("ucn"),
		Address:     r.Form.Get("address"),
		PhoneNumber: r.Form.Get("phone_number"),
	}
}
