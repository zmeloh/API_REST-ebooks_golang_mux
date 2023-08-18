package rest

import (
	//"database/sql"
	"encoding/json"
	"strconv"

	//"example/api/dao"
	"example/api/models"
	"example/api/services"
	"example/api/utils"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	//"strconv"
	//"github.com/gorilla/mux"
)

// CreateUser crée un nouvel utilisateur en utilisant les données du corps de la requête.
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser models.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		utils.Logger(err)
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	services.InsertUser(&newUser)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}

// GetUserByID récupère un utilisateur par son ID.
func GetUserByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	result, err := strconv.Atoi(id)
	if err != nil {
		utils.Logger(err)
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	user := services.GetUserByID(result)
	if user.ID == 0 {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

// GetAllUsers récupère tous les utilisateurs de la base de données.
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users := services.GetAllUsers()
	if users == nil {
		http.Error(w, "No data found ", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

// UpdateUser met à jour les informations d'un utilisateur par son ID.
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	userID, err := strconv.Atoi(id)
	if err != nil {
		utils.Logger(err)
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	var updatedUser models.User
	err = json.NewDecoder(r.Body).Decode(&updatedUser)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	err = services.UpdateUser(userID, &updatedUser)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedUser)
}

// DeleteUser supprime un utilisateur par son ID.
func DeleteUser(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "User with has been deleted")
}
