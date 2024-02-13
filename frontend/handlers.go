package main

import (
	"fmt"
	"net/http"
)

func loginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	if r.Method == http.MethodPost {
		email := r.FormValue("email")
		password := r.FormValue("password")
		role := r.FormValue("role") // Retrieve the selected role

		// Here you would add your authentication logic (e.g., check the credentials against a database)

		if email == "test" && password == "password" {
			fmt.Fprintf(w, "Welcome, %s! You are logged in as a %s.", email, role)
			return
		}

		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Render the login form
	// In loginHandler
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
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	if r.Method == http.MethodPost {
		email := r.FormValue("email")
		//password := r.FormValue("password")
		role := r.FormValue("role") // Retrieve the selected role

		// Here you would add your registration logic (e.g., store the new user's details in a database)

		fmt.Fprintf(w, "Welcome, %s! You have successfully registered as a %s.", email, role)
		return
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
