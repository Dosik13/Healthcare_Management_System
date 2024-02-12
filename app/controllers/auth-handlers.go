package controllers

import (
	"Healthcare_Management_System/app/models"
	"Healthcare_Management_System/utils"
	"fmt"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

var users []*models.Patient

type AuthController struct {
	DB *gorm.DB
}

func NewAuthController(db *gorm.DB) *AuthController {
	return &AuthController{DB: db}
}

func (ac *AuthController) LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if r.Method == "POST" {
		r.ParseForm()
		username, password := r.Form.Get("username"), r.Form.Get("password")

		// Check if the username exists
		userExists := false
		for _, user := range users {
			if user.Email == username {
				userExists = true
				break
			}
		}

		if !userExists {
			fmt.Fprintf(w, "This username is not registered!")
			return
		}

		if utils.CheckPasswordHash(username, password) {
			http.SetCookie(w, &http.Cookie{
				Name:   "session_token",
				Value:  username,
				MaxAge: 0, // The cookie will be deleted when the user closes their browser
			})
			http.Redirect(w, r, "/welcome", http.StatusSeeOther)
		} else {
			fmt.Fprintf(w, "Invalid credentials!")
		}
	} else {
		fmt.Fprintf(w, `<form method="POST">
        Username: <input type="text" name="username">
        Password: <input type="password" name="password">
        <input type="submit" value="Login">
        </form>
        <a href="/register">Register</a>`)
	}
}

func (ac *AuthController) CheckCredentials(username, password string) bool {
	hashedPassword := utils.HashPassword(password)
	for _, user := range users {
		if user.Email == username && user.Password == hashedPassword {
			return true
		}
	}
	return false
}

func (ac *AuthController) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			return
		}
		role := r.Form.Get("role")

		switch role {
		case "Doctor":
			var doctor models.Doctor
			// Populate and validate doctor fields
			// Save doctor to database
			doctor.User = populateUser(r)
			doctor.Specialization = r.Form.Get("specialization")
			yearOfExperience, _ := strconv.Atoi(r.Form.Get("year_of_experience"))
			doctor.YearOfExperience = uint(yearOfExperience)

			result := ac.DB.Create(&doctor)
			if result.Error != nil {
				fmt.Fprintf(w, "Error registering user: %s", result.Error)
				return
			}

		case "Patient":
			var patient models.Patient
			// Populate and validate patient fields
			// Save patient to database
			// Handle other roles similarly
			patient.User = populateUser(r)
			patient.Allergies = r.Form.Get("allergies")
			patient.MedicalHistory = r.Form.Get("medical_history")

			result := ac.DB.Create(&patient)
			if result.Error != nil {
				fmt.Fprintf(w, "Error registering user: %s", result.Error)
				return
			}

		case "Nurse":
			var nurse models.Nurse
			// Populate and validate nurse fields
			// Save nurse to database
			nurse.User = populateUser(r)

			yearOfExperience, _ := strconv.Atoi(r.Form.Get("year_of_experience"))
			nurse.YearOfExperience = uint(yearOfExperience)

			result := ac.DB.Create(&nurse)
			if result.Error != nil {
				fmt.Fprintf(w, "Error registering user: %s", result.Error)
				return
			}
		}

	} else {
		// Render the registration form

	}
}

//
//func (ac *AuthController) RegisterHandler(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "text/html")
//	if r.Method == "POST" {
//		r.ParseForm()
//		email, password := r.Form.Get("email"), r.Form.Get("password")
//		firstName, lastName := r.Form.Get("first_name"), r.Form.Get("last_name")
//		ucn, address := r.Form.Get("ucn"), r.Form.Get("address")
//		phoneNumber := r.Form.Get("phone_number")
//
//		// Check if a user with the provided username already exists
//		for _, user := range users {
//			if user.Email == email {
//				fmt.Fprintf(w, "A user with this email already exists!")
//				return
//			}
//		}
//
//		// Hash the password before storing it
//		hashedPassword := utils.HashPassword(password)
//
//		// Create a new User instance
//		newUser := &models.Patient{}
//		newUser.Email = username
//		newUser.Password = hashedPassword
//
//		result := ac.DB.Create(&newUser)
//		if result.Error != nil {
//			fmt.Fprintf(w, "Error registering user: %s", result.Error)
//			return
//		}
//
//		fmt.Fprintf(w, "User registered successfully!")
//	} else {
//		fmt.Fprintf(w, `<form method="POST">
//        Username: <input type="text" name="username">
//        Password: <input type="password" name="password">
//        <input type="submit" value="Register">
//        </form>`)
//	}
//}

func (ac *AuthController) WelcomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	cookie, err := r.Cookie("session_token")
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// Check if the user associated with the session token still exists
	userExists := false
	for _, user := range users {
		if user.Email == cookie.Value { //Dani tuka go smenih s EMAIl
			userExists = true
			break
		}
	}

	if !userExists {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	//	fmt.Fprintf(w, Welcome, %s! <a href="/logout">Logout</a>, cookie.Value)
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
