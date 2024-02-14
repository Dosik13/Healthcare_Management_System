package controllers

import (
	"Healthcare_Management_System/app/models"
	"Healthcare_Management_System/utils"
	"encoding/json"
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
	if utils.CheckIfLogged(w, r) == true {
		return
	}

	if r.Method == "POST" {

		email := r.FormValue("username")
		password := r.FormValue("password")
		role := r.FormValue("role")

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

		session, _ := utils.Store.Get(r, "SessionID")

		session.Values["user"] = user.UserId
		session.Values["role"] = role
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
			http.Redirect(w, r, "/patient_dashboard", http.StatusSeeOther)
		}

	} else {
		tmpl, err := template.ParseFiles("frontend/templates/login.html")
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

func (ac *AuthController) HomeHandler(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("frontend/templates/homepage.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (ac *AuthController) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	if utils.CheckIfLogged(w, r) == true {
		return
	}

	if r.Method == "POST" {

		role := r.FormValue("role")

		switch role {
		case "doctor":
			var doctor models.Doctor
			var doctor2 models.Doctor
			doctor.User = populateUser(r)
			doctor.Specialization = r.FormValue("specialization")
			yearOfExperience, _ := strconv.Atoi(r.FormValue("year_of_experience"))
			doctor.YearOfExperience = uint(yearOfExperience)
			doctor.MoreInfo = r.FormValue("about")
			doctor.User.Role = role

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
			patient.User.Role = role

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
			nurse.User.Role = role

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

	} else {
		tmpl, err := template.ParseFiles("frontend/templates/register.html")
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

func (ac *AuthController) PatientDashboardHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := utils.Store.Get(r, "SessionID")
	userId, ok := session.Values["user"].(uint)
	if !ok || userId == 0 {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	var patient models.Patient

	if result := ac.DB.Preload("Appointments").Preload("Billings").Preload("MedicalRecords").First(&patient, "user_id = ?", userId); result.Error != nil {
		http.Error(w, "Patient not found", http.StatusNotFound)
		return
	}

	tmpl, err := template.ParseFiles("frontend/templates/patient.html")
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, patient)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}

}

func (ac *AuthController) ChangeEmailHandler(w http.ResponseWriter, r *http.Request) {
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
		role, _ := session.Values["role"].(string)

		newEmail := r.FormValue("newemail")

		switch role {
		case "patient":
			var patient2 models.Patient

			if result := ac.DB.First(&patient2, "email = ?", newEmail); result.Error != nil {
				result := ac.DB.Model(&models.Patient{}).Where("user_id = ?", userId).Update("email", newEmail)
				if result.Error != nil {
					http.Error(w, "Failed to update email", http.StatusInternalServerError)
					return
				}

				http.Redirect(w, r, "/login", http.StatusSeeOther)
			} else {
				http.Error(w, "User with this email already exists!", http.StatusConflict)
				return
			}

		case "doctor":
			var doctor2 models.Doctor

			if result := ac.DB.First(&doctor2, "email = ?", newEmail); result.Error != nil {
				result := ac.DB.Model(&models.Doctor{}).Where("user_id = ?", userId).Update("email", newEmail)
				if result.Error != nil {
					http.Error(w, "Failed to update email", http.StatusInternalServerError)
					return
				}

				http.Redirect(w, r, "/login", http.StatusSeeOther)
			} else {
				http.Error(w, "User with this email already exists!", http.StatusConflict)
				return
			}
		}

	} else {
		tmpl, err := template.ParseFiles("frontend/templates/changeEmail.html")
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

func (ac *AuthController) ChangePasswordHandler(w http.ResponseWriter, r *http.Request) {
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
		role, _ := session.Values["role"].(string)

		password := r.FormValue("password")
		newPassword := r.FormValue("newpassword")

		switch role {
		case "patient":
			var patient models.Patient
			if err := ac.DB.Where("user_id = ?", userId).First(&patient).Error; err == nil {

				if utils.CheckPasswordHash(password, patient.Password) != true {
					http.Error(w, "Passwords do not match!", http.StatusConflict)
					return
				}
				err := ac.DB.Model(&patient).Update("password", utils.HashPassword(newPassword))
				if err != nil {
					http.Error(w, "Failed to update password", http.StatusInternalServerError)
					return
				}
			}

			http.Redirect(w, r, "/login", http.StatusSeeOther)

		case "doctor":
			var doctor models.Doctor
			if err := ac.DB.Where("user_id = ?", userId).First(&doctor).Error; err == nil {

				if utils.CheckPasswordHash(password, doctor.Password) != true {
					http.Error(w, "Passwords do not match!", http.StatusConflict)
					return
				}
				err := ac.DB.Model(&doctor).Update("password", utils.HashPassword(newPassword))
				if err != nil {
					http.Error(w, "Failed to update password", http.StatusInternalServerError)
					return
				}

			}

			http.Redirect(w, r, "/login", http.StatusSeeOther)
		}

	} else {
		tmpl, err := template.ParseFiles("frontend/templates/changePass.html")
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

type DocAttr struct {
	Name      string   `json:"name"`
	Specialty string   `json:"specialty"`
	Days      []string `json:"days"`
	ID        uint     `json:"id"`
}

func (ac *AuthController) DoctorSearchHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := utils.Store.Get(r, "SessionID")
	userId, ok := session.Values["user"].(uint)
	if !ok || userId == 0 {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	var doctors []models.Doctor
	result := ac.DB.Preload("Appointments").Preload("Ratings").Preload("Patients").Find(&doctors)
	if result.Error != nil {
		http.Error(w, "Error retrieving doctors", http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("frontend/templates/search.html")
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}

	DocAttrs := make([]DocAttr, len(doctors))
	for i, doctor := range doctors {
		var dates []string
		for _, appointment := range doctor.Appointments {
			dates = append(dates, appointment.Date)
		}
		DocAttrs[i] = DocAttr{Name: doctor.User.FirstName + " " + doctor.User.LastName, Specialty: doctor.Specialization, ID: doctor.User.UserID, Days: dates}
	}

	docJSON, err := json.Marshal(DocAttrs)
	if err != nil {
		http.Error(w, "Failed to serialize appointments", http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"doctorsJS": string(docJSON),
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
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
