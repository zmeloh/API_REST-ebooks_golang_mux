package rest

import (
	//"database/sql"
	"encoding/json"
	//"example/api/dao"
	"example/api/models"
	"example/api/utils"
	"fmt"
	"net/http"
	//"strconv"
	//"github.com/gorilla/mux"
)

// CreateUser crée un nouvel utilisateur en utilisant les données du corps de la requête.
func CreateUser(w http.ResponseWriter, r *http.Request) {
	// Decodage des Headers
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		utils.Logger(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// GetUserByID récupère un utilisateur par son ID.
func GetUserByID(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewEncoder(w).Encode(user)
}

// GetAllUsers récupère tous les utilisateurs de la base de données.
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(users)
	if err != nil {
		utils.Logger(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// UpdateUser met à jour les informations d'un utilisateur par son ID.
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var updatedUser models.User
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedUser)
}

// DeleteUser supprime un utilisateur par son ID.
func DeleteUser(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "User with has been deleted")
}
