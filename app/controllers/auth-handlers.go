package controllers

import (
	"Healthcare_Management_System/app/models"
	"Healthcare_Management_System/utils"
	"fmt"
	"golang.org/x/crypto/bcrypt"
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

	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse form data
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	email := r.FormValue("email")
	password := r.FormValue("password")
	role := r.FormValue("role") // Retrieve the selected role

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

	fmt.Fprint(w, `
<!DOCTYPE html>
<html>
<head>
    <title>Login</title>
    <link rel="stylesheet" type="text/css" href="/static/styles.css">
</head>
<body>
    <h1 class="fancy-title">SuperDoc</h1>
    <div>
        <p>Please log in to continue.</p>
    </div>
    <form method="POST">
        <label for="email">Email:</label><br>
        <input type="text" id="username" name="username" required><br>
        <label for="password">Password:</label><br>
        <input type="password" id="password" name="password" required><br>
        <label for="role">Role:</label><br>
        <select id="role" name="role" required>
            <option value="patient">Patient</option>
            <option value="doctor">Doctor</option>
        </select><br>
        <input type="submit" value="Submit">
    </form>
    <p>Don't have an account? <a href="/register">Register here</a></p>
</body>
</html>`)

	// Redirect or respond to indicate success
	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
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

	}
	fmt.Fprint(w, `
<!DOCTYPE html>
<html>
<head>
    <title>Register</title>
    <link rel="stylesheet" type="text/css" href="/static/styles.css">
    <script>
        function updateForm() {
            var role = document.getElementById("role").value;
            var extraFields = document.getElementById("extraFields");

            if (role === "doctor") {
                extraFields.innerHTML = '<label for="experience">Years of Experience:</label><br><input type="number" id="experience" name="experience" min="0" required><br>';
            } else {
                extraFields.innerHTML = '';
            }
        }
    </script>
</head>
<body>
    <h1 class="fancy-title">SuperDoc</h1>
    <div>
        <p>Please register to continue.</p>
    </div>
    <form method="POST">
        <label for="email">Email:</label><br>
        <input type="text" id="username" name="username" required><br>
        <label for="password">Password:</label><br>
        <input type="password" id="password" name="password" required><br>
        <label for="role">Role:</label><br>
        <select id="role" name="role" onchange="updateForm()" required>
            <option value="patient">Patient</option>
            <option value="doctor">Doctor</option>
        </select><br>
        <div id="extraFields"></div>
        <input type="submit" value="Register">
    </form>
    <p>Already have an account? <a href="/login">Login here</a></p>
</body>
</html>`)

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
