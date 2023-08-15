package dao

import (
	"database/sql"
	"encoding/json"
	"example/api/models"
	"fmt"

	//"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// CreateUser crée un nouvel utilisateur en utilisant les données du corps de la requête.
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Créez la table "users" si elle n'existe pas
	_, err = DB.Exec(`CREATE TABLE IF NOT EXISTS users (
		id INT AUTO_INCREMENT PRIMARY KEY,
		username VARCHAR(255),
		email VARCHAR(255)
	)`)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Exécute la requête pour insérer un nouvel utilisateur dans la base de données
	result, err := DB.Exec("INSERT INTO users (username, email) VALUES (?, ?)", user.Username, user.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Obtient l'ID généré lors de l'insertion
	userID, err := result.LastInsertId()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Met à jour l'ID dans l'objet User
	user.ID = int(userID)
}

func selectAllUsers(w http.ResponseWriter, r *http.Request) {

	var user models.User

	// Exécute la requête pour récupérer les informations de l'utilisateur depuis la base de données
	err := DB.QueryRow("SELECT id, username, email FROM users WHERE id = ?", userID).Scan(&user.ID, &user.Username, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(user)
}

// GetAllUsers récupère tous les utilisateurs de la base de données.
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	rows, err := DB.Query("SELECT id, username, email FROM users")
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
}

func selectUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]

	var user models.User

	// Exécute la requête pour récupérer les informations de l'utilisateur depuis la base de données
	err := DB.QueryRow("SELECT id, username, email FROM users WHERE id = ?", userID).Scan(&user.ID, &user.Username, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			http.NotFound(w, r)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]

	var updatedUser models.User
	err := json.NewDecoder(r.Body).Decode(&updatedUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Vérifie si l'utilisateur avec l'ID donné existe
	var existingUserID int
	err = DB.QueryRow("SELECT id FROM users WHERE id = ?", userID).Scan(&existingUserID)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Exécute la requête pour mettre à jour les informations de l'utilisateur dans la base de données
	_, err = DB.Exec("UPDATE users SET username = ?, email = ? WHERE id = ?", updatedUser.Username, updatedUser.Email, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Convertit l'ID en entier
	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Met à jour l'ID dans l'objet updatedUser
	updatedUser.ID = userIDInt
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]

	// Vérifie si l'utilisateur avec l'ID donné existe
	var existingUserID int
	err := DB.QueryRow("SELECT id FROM users WHERE id = ?", userID).Scan(&existingUserID)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Supprime l'utilisateur de la base de données
	_, err = DB.Exec("DELETE FROM users WHERE id = ?", userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "User with ID %s has been deleted", userID)
}
