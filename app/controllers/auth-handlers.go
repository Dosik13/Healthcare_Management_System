package controllers

import (
	"Healthcare_Management_System/app/models"
	"Healthcare_Management_System/utils"
	"fmt"
	"gorm.io/gorm"
	"html/template"
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

func (ac *AuthController) emailExists(email string) bool {
	var user models.User
	if result := ac.DB.First(&user, "email = ?", email); result.Error != nil {
		return false
	}
	return true
}

func NewAuthController(db *gorm.DB) *AuthController {
	return &AuthController{DB: db}
}

func (ac *AuthController) LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	session, err := utils.Store.Get(r, "SessionID")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	if _, ok := session.Values["user"]; ok {
		http.Redirect(w, r, "/home", http.StatusSeeOther)
		return
	}
	if r.Method == "POST" {

		email := r.FormValue("username")
		password := r.FormValue("password")
		role := r.FormValue("role") // Retrieve the selected role

		var user RoleUser

		switch role {
		case "doctor":
			var doctor models.Doctor
			if err := ac.DB.Where("email = ?", email).First(&doctor).Error; err == nil {
				user = RoleUser{Email: doctor.Email, Password: doctor.Password, UserId: doctor.UserID}
			}
		case "nurse":
			var nurse models.Nurse
			if err := ac.DB.Where("email = ?", email).First(&nurse).Error; err == nil {
				user = RoleUser{Email: nurse.Email, Password: nurse.Password, UserId: nurse.UserID}
			}
		case "patient":
			var patient models.Patient
			if err := ac.DB.Where("email = ?", email).First(&patient).Error; err == nil {
				user = RoleUser{Email: patient.Email, Password: patient.Password, UserId: patient.UserID}
			}
		default:
			http.Error(w, "Invalid role specified", http.StatusBadRequest)
			return
		}

		if utils.CheckPasswordHash(password, user.Password) != true {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		// Create session
		session, _ := utils.Store.Get(r, "SessionID")
		// Depending on your session library, you might need to cast user.ID to the appropriate type
		session.Values["user"] = user.UserId // Store user ID in session
		session.Values["role"] = role        // Store role in session
		err := session.Save(r, w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if role == "doctor" {
			http.Redirect(w, r, "/doctor_dashboard", http.StatusSeeOther)
		}
		if role == "nurse" {
			http.Redirect(w, r, "/nurse_dashboard", http.StatusSeeOther)
		}
		if role == "patient" {
			http.Redirect(w, r, "/home", http.StatusSeeOther)
		}

	} else {
		tmpl, err := template.ParseFiles("frontend/templates/login.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Execute the template and send the result to the client
		err = tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

}

func (ac *AuthController) HomeHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the HTML template
	tmpl, err := template.ParseFiles("frontend/templates/homepage.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Execute the template and send the result to the client
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (ac *AuthController) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if r.Method == "POST" {

		role := r.FormValue("role")

		//	fmt.Println("Role: ", role)

		switch role {
		case "doctor":
			var doctor models.Doctor
			var doctor2 models.Doctor
			doctor.User = populateUser(r)
			doctor.Specialization = r.FormValue("specialization")
			yearOfExperience, _ := strconv.Atoi(r.FormValue("year_of_experience"))
			doctor.YearOfExperience = uint(yearOfExperience)
			doctor.MoreInfo = r.FormValue("about")

			if result := ac.DB.First(&doctor2, "email = ?", doctor.Email); result.Error != nil {
				result := ac.DB.Create(&doctor)
				if result.Error != nil {
					fmt.Fprintf(w, "Error registering doctor: %s", result.Error)
					return
				}
				http.Redirect(w, r, "/login", http.StatusSeeOther)
			} else {
				http.Error(w, "Doctor already exists!", http.StatusConflict)
				return
			}

		case "patient":
			var patient models.Patient
			var patient2 models.Patient
			patient.User = populateUser(r)

			//fmt.Println("Email: ", patient.Email)
			//fmt.Println("FirstName: ", patient.FirstName)
			//fmt.Println("MiddleName", patient.MiddleName)
			//fmt.Println("LastName: ", patient.LastName)
			//fmt.Println("Password: ", patient.Password)
			//fmt.Println("UCN: ", patient.UCN)
			//fmt.Println("Address: ", patient.Address)
			//fmt.Println("PhoneNumber: ", patient.PhoneNumber)

			if result := ac.DB.First(&patient2, "email = ?", patient.Email); result.Error != nil {
				result := ac.DB.Create(&patient)
				if result.Error != nil {
					fmt.Fprintf(w, "Error registering patient: %s", result.Error)
					return
				}
				http.Redirect(w, r, "/login", http.StatusSeeOther)
			} else {
				http.Error(w, "Patient already exists!", http.StatusConflict)
				return
			}

		case "nurse":
			var nurse models.Nurse
			var nurse2 models.Nurse
			nurse.User = populateUser(r)
			yearOfExperience, _ := strconv.Atoi(r.FormValue("year_of_experience"))
			nurse.YearOfExperience = uint(yearOfExperience)
			nurse.MoreInfo = r.FormValue("about")

			if result := ac.DB.First(&nurse2, "email = ?", nurse.Email); result.Error != nil {
				result := ac.DB.Create(&nurse)
				if result.Error != nil {
					fmt.Fprintf(w, "Error registering nurse: %s", result.Error)
					return
				}
				http.Redirect(w, r, "/login", http.StatusSeeOther)
			} else {
				http.Error(w, "Nurse already exists!", http.StatusConflict)
				return
			}
		}

		http.Redirect(w, r, "/login", http.StatusSeeOther)

	} else {
		tmpl, err := template.ParseFiles("frontend/templates/register.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Execute the template and send the result to the client
		err = tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	}
}

func (ac *AuthController) LogoutHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := utils.Store.Get(r, "SessionID")

	session.Options.MaxAge = -1
	err := session.Save(r, w)
	if err != nil {
		http.Error(w, "Failed to logout", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func (ac *AuthController) DoctorDashboardHandler(w http.ResponseWriter, r *http.Request) {
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
	fmt.Printf("Doctor: %+v\n", doctor)
	tmpl, err := template.ParseFiles("frontend/templates/doctor.html")
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, doctor)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}

}

func populateUser(r *http.Request) models.User {
	return models.User{
		FirstName:   r.FormValue("first_name"),
		MiddleName:  r.FormValue("middle_name"),
		LastName:    r.FormValue("last_name"),
		Email:       r.FormValue("username"),
		Password:    utils.HashPassword(r.FormValue("password")),
		UCN:         r.FormValue("ucn"),
		Address:     r.FormValue("address"),
		PhoneNumber: r.FormValue("phone_number"),
		Gender:      r.FormValue("gender"),
	}
}
