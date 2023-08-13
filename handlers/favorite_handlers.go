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

func CreateFavorite(w http.ResponseWriter, r *http.Request) {
	var favorite models.Favorite
	err := json.NewDecoder(r.Body).Decode(&favorite)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = database.DB.Exec("INSERT INTO favorites (user_id, ebook_id) VALUES (?, ?)", favorite.UserID, favorite.EbookID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(favorite)
}


func GetAllFavorites(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query("SELECT id, user_id, ebook_id FROM favorites")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var favorites []models.Favorite
	for rows.Next() {
		var favorite models.Favorite
		err := rows.Scan(&favorite.ID, &favorite.UserID, &favorite.EbookID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		favorites = append(favorites, favorite)
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(favorites)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}


func GetFavoriteByID(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	favoriteID := vars["id"]

	var favorite models.Favorite

	err := database.DB.QueryRow("SELECT id, user_id, ebook_id FROM favorites WHERE id = ?", favoriteID).Scan(&favorite.ID, &favorite.UserID, &favorite.EbookID)
	if err != nil {
		if err == sql.ErrNoRows{
			http.NotFound(w,r)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(favorite)
}


func GetFavoritesByUserID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["userID"]

	rows, err := database.DB.Query("SELECT id, user_id, ebook_id FROM favorites WHERE user_id = ?", userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var favorites []models.Favorite
	for rows.Next() {
		var favorite models.Favorite
		err := rows.Scan(&favorite.ID, &favorite.UserID, &favorite.EbookID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		favorites = append(favorites, favorite)
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(favorites)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func GetFavoritesByEbookID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ebookID := vars["ebookID"]

	rows, err := database.DB.Query("SELECT id, user_id, ebook_id FROM favorites WHERE ebook_id = ?", ebookID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var favorites []models.Favorite
	for rows.Next() {
		var favorite models.Favorite
		err := rows.Scan(&favorite.ID, &favorite.UserID, &favorite.EbookID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		favorites = append(favorites, favorite)
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(favorites)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func UpdateFavorite(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	favoriteID := vars["id"]

	var updatedFavorite models.Favorite
	err := json.NewDecoder(r.Body).Decode(&updatedFavorite)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = database.DB.Exec("UPDATE favorites SET user_id = ?, ebook_id = ? WHERE id = ?", updatedFavorite.UserID, updatedFavorite.EbookID, favoriteID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(updatedFavorite)
}


func DeleteFavorite(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	favoriteID := vars["id"]
	_, err := database.DB.Exec("DELETE FROM favorites WHERE id = ?", favoriteID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Favorite with ID %s has been deleted", favoriteID)
}
