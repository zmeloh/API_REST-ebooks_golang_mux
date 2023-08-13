package handlers

import (
	"database/sql"
	"encoding/json"
	"example/api/database"
	"example/api/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// CreateFavorite crée un nouveau favori.
func CreateFavorite(w http.ResponseWriter, r *http.Request) {
	// Décode les données JSON du corps de la requête dans une struct Favorite
	var favorite models.Favorite
	err := json.NewDecoder(r.Body).Decode(&favorite)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Insère le favori dans la base de données
	result, err := database.DB.Exec("INSERT INTO favorites (user_id, ebook_id) VALUES (?, ?)", favorite.UserID, favorite.EbookID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Récupère l'ID du favori nouvellement inséré
	favoriteID, err := result.LastInsertId()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Met à jour l'ID dans l'objet favorite
	favorite.ID = int(favoriteID)

	// Répond avec le favori créé et le code de statut 201 (Created)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(favorite)
}

// GetAllFavorites récupère tous les favoris.
func GetAllFavorites(w http.ResponseWriter, r *http.Request) {
	// Récupère tous les favoris depuis la base de données
	rows, err := database.DB.Query("SELECT id, user_id, ebook_id FROM favorites")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Crée une slice pour stocker les favoris
	var favorites []models.Favorite
	for rows.Next() {
		var favorite models.Favorite
		// Scan les valeurs des colonnes dans la struct Favorite
		err := rows.Scan(&favorite.ID, &favorite.UserID, &favorite.EbookID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		favorites = append(favorites, favorite)
	}

	// Configure l'en-tête Content-Type pour la réponse JSON
	w.Header().Set("Content-Type", "application/json")
	// Encode la slice de favoris en JSON et répond avec les favoris
	err = json.NewEncoder(w).Encode(favorites)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// GetFavoriteByID récupère un favori par son ID.
func GetFavoriteByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	favoriteID := vars["id"]

	var favorite models.Favorite

	// Interroge la base de données pour obtenir le favori correspondant à l'ID
	err := database.DB.QueryRow("SELECT id, user_id, ebook_id FROM favorites WHERE id = ?", favoriteID).Scan(&favorite.ID, &favorite.UserID, &favorite.EbookID)
	if err != nil {
		if err == sql.ErrNoRows {
			// Si aucun favori n'est trouvé, répond avec le statut 404 (Not Found)
			http.NotFound(w, r)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Encode le favori en JSON et répond avec le favori
	json.NewEncoder(w).Encode(favorite)
}

// GetFavoritesByUserID récupère les favoris par ID d'utilisateur.
func GetFavoritesByUserID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["userID"]

	// Interroge la base de données pour obtenir les favoris correspondant à l'ID d'utilisateur
	rows, err := database.DB.Query("SELECT id, user_id, ebook_id FROM favorites WHERE user_id = ?", userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Crée une slice pour stocker les favoris
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

	// Configure l'en-tête Content-Type pour la réponse JSON
	w.Header().Set("Content-Type", "application/json")
	// Encode la slice de favoris en JSON et répond avec les favoris
	err = json.NewEncoder(w).Encode(favorites)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// GetFavoritesByEbookID récupère les favoris par ID d'ebook.
func GetFavoritesByEbookID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ebookID := vars["ebookID"]

	// Interroge la base de données pour obtenir les favoris correspondant à l'ID d'ebook
	rows, err := database.DB.Query("SELECT id, user_id, ebook_id FROM favorites WHERE ebook_id = ?", ebookID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Crée une slice pour stocker les favoris
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

	// Configure l'en-tête Content-Type pour la réponse JSON
	w.Header().Set("Content-Type", "application/json")
	// Encode la slice de favoris en JSON et répond avec les favoris
	err = json.NewEncoder(w).Encode(favorites)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// UpdateFavorite met à jour un favori par son ID.
func UpdateFavorite(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	favoriteID := vars["id"]

	var updatedFavorite models.Favorite
	err := json.NewDecoder(r.Body).Decode(&updatedFavorite)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Met à jour les données du favori dans la base de données
	_, err = database.DB.Exec("UPDATE favorites SET user_id = ?, ebook_id = ? WHERE id = ?", updatedFavorite.UserID, updatedFavorite.EbookID, favoriteID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Convertit l'ID du favori en int
	favoriteIDInt, err := strconv.Atoi(favoriteID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Met à jour l'ID dans l'objet updatedFavorite
	updatedFavorite.ID = favoriteIDInt

	// Configure l'en-tête Content-Type pour la réponse JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// Encode le favori mis à jour en JSON et répond avec le favori
	json.NewEncoder(w).Encode(updatedFavorite)
}

// DeleteFavorite supprime un favori par son ID.
func DeleteFavorite(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	favoriteID := vars["id"]
	_, err := database.DB.Exec("DELETE FROM favorites WHERE id = ?", favoriteID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Répond avec le statut 200 (OK) et un message indiquant que le favori a été supprimé
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Favorite with ID %s has been deleted", favoriteID)
}
