package rest

import (
	"encoding/json"
	"strconv"

	"example/api/models"
	"example/api/services"
	"example/api/utils"
	"net/http"

	"github.com/gorilla/mux"
)

// CreateUser crée un nouvel utilisateur en utilisant les données du corps de la requête.
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser models.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		utils.Logger(err)
		ServerResponse(w, http.StatusBadRequest, "Invalid reqquest")
		return
	}

	err = services.InsertUser(&newUser)
	if err != nil {
		ServerResponse(w, http.StatusAlreadyReported, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	//ServerResponse(w, http.StatusCreated, newUser)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(newUser)
}

// GetUserByID récupère un utilisateur par son ID.
func GetUserByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	result, err := strconv.Atoi(id)
	if err != nil {
		utils.Logger(err)
		ServerResponse(w, http.StatusBadRequest, "Invalid id")
		return
	}

	user := services.GetUserByID(result)
	if user.ID == 0 {
		ServerResponse(w, http.StatusNotFound, "User not found")
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

// GetAllUsers récupère tous les utilisateurs de la base de données.
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users := services.GetAllUsers()
	if users == nil {
		ServerResponse(w, http.StatusNotFound, "No data found")
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
		ServerResponse(w, http.StatusBadRequest, "Invalid id")
		return
	}
	var updatedUser models.User
	err = json.NewDecoder(r.Body).Decode(&updatedUser)
	if err != nil {
		ServerResponse(w, http.StatusBadRequest, "Invalid reqeust")
		return
	}

	err = services.UpdateUser(userID, &updatedUser)
	if err != nil {
		ServerResponse(w, http.StatusNotFound, "Ueer not found")
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedUser)
}

// DeleteUser supprime un utilisateur par son ID.
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	userID, err := strconv.Atoi(id)
	if err != nil {
		utils.Logger(err)
		ServerResponse(w, http.StatusBadRequest, "Invalid id")
		return
	}

	err = services.DeleteUser(userID)
	if err != nil {
		ServerResponse(w, http.StatusNotFound, "User not found")
		return
	}
	ServerResponse(w, http.StatusOK, "User has been deleted")
}
