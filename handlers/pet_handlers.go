package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"pet-clinic/auth"
	"pet-clinic/db"
	"pet-clinic/models"
	"pet-clinic/utils"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// helper: extract role & username from claims
func getUserFromRequest(r *http.Request) (username, role string, ok bool) {
	claims, found := auth.GetClaims(r)
	if !found {
		return "", "", false
	}
	u, _ := claims["username"].(string)
	rl, _ := claims["role"].(string)
	return u, rl, true
}

// Add Pet (any authenticated user can add; owners usually add their pets)
func AddPet(w http.ResponseWriter, r *http.Request) {
	utils.Log.Info("POST /pets called")
	var p models.Pet
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		ErrorResponse(w, "Invalid pet input", http.StatusBadRequest, err)
		return
	}

	_, err := db.DB.Exec(`INSERT INTO pets (name, species, breed, owner_id, medical_history)
        VALUES ($1, $2, $3, $4, $5)`,
		p.Name, p.Species, p.Breed, p.OwnerID, p.MedicalHistory)

	if err != nil {
		ErrorResponse(w, "Failed to create pet", http.StatusInternalServerError, err)
		return
	}

	utils.Log.WithField("name", p.Name).Info("Pet added successfully")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Pet created"))
}

// GetPets
func GetPets(w http.ResponseWriter, r *http.Request) {
	utils.Log.Info("GET /pets called")
	rows, err := db.DB.Query("SELECT id, name, species, breed, owner_id, medical_history FROM pets")
	if err != nil {
		ErrorResponse(w, "Failed to fetch pets", http.StatusInternalServerError, err)
		return
	}
	defer rows.Close()

	var pets []models.Pet
	for rows.Next() {
		var p models.Pet
		rows.Scan(&p.ID, &p.Name, &p.Species, &p.Breed, &p.OwnerID, &p.MedicalHistory)
		pets = append(pets, p)
	}

	json.NewEncoder(w).Encode(pets)
}

// UpdatePet - owner can update only their pets; staff can update any pet
func UpdatePet(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id := idStr

	username, role, ok := getUserFromRequest(r)
	if !ok {
		ErrorResponse(w, "Unauthorized", http.StatusUnauthorized, nil)
		return
	}

	// If role is owner, verify ownership
	if role == "owner" {
		// extract owner id from username like "owner1" -> 1
		if !strings.HasPrefix(username, "owner") {
			http.Error(w, "Invalid owner identity", http.StatusForbidden)
			return
		}
		ownerIDStr := strings.TrimPrefix(username, "owner")
		ownerID, err := strconv.Atoi(ownerIDStr)
		if err != nil {
			http.Error(w, "Invalid owner identity", http.StatusForbidden)
			return
		}

		var petOwnerID int
		err = db.DB.QueryRow("SELECT owner_id FROM pets WHERE id=$1", id).Scan(&petOwnerID)
		if err != nil {
			ErrorResponse(w, "Pet not found", http.StatusNotFound, err)
			return
		}
		if petOwnerID != ownerID {
			utils.Log.WithFields(map[string]interface{}{"user": username, "pet_id": id}).Warn("Owner attempted to update other user's pet")
			http.Error(w, "You can only update your own pets", http.StatusForbidden)
			return
		}
	}

	// proceed with update
	var p models.Pet
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		ErrorResponse(w, "Invalid update body", http.StatusBadRequest, err)
		return
	}

	_, err := db.DB.Exec(`UPDATE pets SET name=$1, species=$2, breed=$3, owner_id=$4, medical_history=$5 WHERE id=$6`,
		p.Name, p.Species, p.Breed, p.OwnerID, p.MedicalHistory, id)

	if err != nil {
		ErrorResponse(w, "Failed to update pet", http.StatusInternalServerError, err)
		return
	}

	utils.Log.WithField("id", id).Info("Pet updated successfully by " + username)
	w.Write([]byte("Pet updated successfully"))
}

// DeletePet - owner can delete only their pets; staff can delete any pet
func DeletePet(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id := idStr

	username, role, ok := getUserFromRequest(r)
	if !ok {
		ErrorResponse(w, "Unauthorized", http.StatusUnauthorized, nil)
		return
	}

	if role == "owner" {
		if !strings.HasPrefix(username, "owner") {
			http.Error(w, "Invalid owner identity", http.StatusForbidden)
			return
		}
		ownerIDStr := strings.TrimPrefix(username, "owner")
		ownerID, err := strconv.Atoi(ownerIDStr)
		if err != nil {
			http.Error(w, "Invalid owner identity", http.StatusForbidden)
			return
		}

		var petOwnerID int
		err = db.DB.QueryRow("SELECT owner_id FROM pets WHERE id=$1", id).Scan(&petOwnerID)
		if err != nil {
			ErrorResponse(w, "Pet not found", http.StatusNotFound, err)
			return
		}
		if petOwnerID != ownerID {
			utils.Log.WithFields(map[string]interface{}{"user": username, "pet_id": id}).Warn("Owner attempted to delete other user's pet")
			http.Error(w, "You can only delete your own pets", http.StatusForbidden)
			return
		}
	}

	_, err := db.DB.Exec("DELETE FROM pets WHERE id=$1", id)
	if err != nil {
		ErrorResponse(w, "Failed to delete pet", http.StatusInternalServerError, err)
		return
	}

	utils.Log.WithField("id", id).Warn(fmt.Sprintf("Pet deleted by %s", username))
	w.Write([]byte("Pet deleted"))
}
