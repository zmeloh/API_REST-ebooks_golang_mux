package handlers

import (
	"example/api/models"
	"example/api/database"
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"database/sql"
	"fmt"
)


func CreateUser(w http.ResponseWriter, r *http.Request){
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = database.DB.Exec("INSERT INTO users (username, email) VALUES (?, ?)", user.Username, user.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}


func GetUserByID(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	userID := vars["id"]

	var user models.User

	err := database.DB.QueryRow("SELECT id, username, email FROM users WHERE id = ?", userID).Scan(&user.ID, &user.Username, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows{
			http.NotFound(w,r)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(user)
}


func GetAllUsers(w http.ResponseWriter, r *http.Request) {
    rows, err := database.DB.Query("SELECT id, username, email FROM users")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var users []models.User
    for rows.Next() {
        var user models.User
        err := rows.Scan(&user.ID, &user.Username, &user.Email)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        users = append(users, user)
    }

    w.Header().Set("Content-Type", "application/json")
    err = json.NewEncoder(w).Encode(users)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}


func UpdateUser(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	userID := vars["id"]

	var updatedUser models.User
	err := json.NewDecoder(r.Body).Decode(&updatedUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = database.DB.Exec("UPDATE users SET username = ?, email = ? WHERE id = ?", updatedUser.Username, updatedUser.Email, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(updatedUser)
}


func DeleteUser(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	userID := vars["id"]
	_, err := database.DB.Exec("DELETE FROM users WHERE id = ?", userID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "User with ID %s has been deleted", userID)
}