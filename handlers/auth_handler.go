package handlers

import (
	"encoding/json"
	"net/http"
	"pet-clinic/auth"
	"pet-clinic/utils"
)

// simple demo users (Option A style usernames: owner<ID>)
var demoUsers = []struct {
	Username string
	Password string
	Role     string
}{
	{Username: "staff1", Password: "staffpass", Role: "staff"},
	{Username: "owner1", Password: "ownerpass", Role: "owner"}, // owner for owner_id = 1
}

// find user helper
func findDemoUser(username, password string) (string, bool) {
	for _, u := range demoUsers {
		if u.Username == username && u.Password == password {
			return u.Role, true
		}
	}
	return "", false
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	utils.Log.Debug("Received login request")

	var username, password string

	// Basic Auth first
	u, p, ok := r.BasicAuth()
	if ok {
		username = u
		password = p
		utils.Log.WithField("username", username).Debug("Attempting login via Basic Auth")
	} else {
		var creds User
		if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
			utils.Log.WithError(err).Warn("Invalid JSON in login request")
			http.Error(w, "Invalid JSON format", http.StatusBadRequest)
			return
		}
		username = creds.Username
		password = creds.Password
		utils.Log.WithField("username", username).Debug("Attempting login via JSON body")
	}

	role, ok := findDemoUser(username, password)
	if !ok {
		utils.Log.WithField("username", username).Warn("Invalid login attempt")
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	token, err := auth.GenerateJWT(username, role)
	if err != nil {
		utils.Log.WithError(err).Error("Failed to generate JWT token")
		http.Error(w, "Could not generate token", http.StatusInternalServerError)
		return
	}

	utils.Log.WithField("username", username).Info("User logged in successfully")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
