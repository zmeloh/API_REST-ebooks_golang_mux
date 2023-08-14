package dao

import (
	"database/sql"
	"encoding/json"
	"example/api/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func InsertFavorite(w http.ResponseWriter, r *http.Request) {
	var favorite models.Favorite
	// Insère le favori dans la base de données
	result, err := DB.Exec("INSERT INTO favorites (user_id, ebook_id) VALUES (?, ?)", favorite.UserID, favorite.EbookID)
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
}

func selectAllFavorites(w http.ResponseWriter, r *http.Request) {
	// Récupère tous les favoris depuis la base de données
	rows, err := DB.Query("SELECT id, user_id, ebook_id FROM favorites")
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
}

func selectFavoriteByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	favoriteID := vars["id"]

	var favorite models.Favorite

	// Interroge la base de données pour obtenir le favori correspondant à l'ID
	err := DB.QueryRow("SELECT id, user_id, ebook_id FROM favorites WHERE id = ?", favoriteID).Scan(&favorite.ID, &favorite.UserID, &favorite.EbookID)
	if err != nil {
		if err == sql.ErrNoRows {
			// Si aucun favori n'est trouvé, répond avec le statut 404 (Not Found)
			http.NotFound(w, r)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func UpdateFavorite(e models.Ebook) {
	vars := mux.Vars(r)
	favoriteID := vars["id"]

	var updatedFavorite models.Favorite
	err := json.NewDecoder(r.Body).Decode(&updatedFavorite)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Met à jour les données du favori dans la base de données
	_, err = DB.Exec("UPDATE favorites SET user_id = ?, ebook_id = ? WHERE id = ?", updatedFavorite.UserID, updatedFavorite.EbookID, favoriteID)
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

}

func deleteFavorite(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	favoriteID := vars["id"]
	_, err := DB.Exec("DELETE FROM favorites WHERE id = ?", favoriteID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func selectFavoriteByUserID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["userID"]

	// Interroge la base de données pour obtenir les favoris correspondant à l'ID d'utilisateur
	rows, err := DB.Query("SELECT id, user_id, ebook_id FROM favorites WHERE user_id = ?", userID)
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
}

func selectFavoriteByEbookID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ebookID := vars["ebookID"]

	// Interroge la base de données pour obtenir les favoris correspondant à l'ID d'ebook
	rows, err := DB.Query("SELECT id, user_id, ebook_id FROM favorites WHERE ebook_id = ?", ebookID)
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
}
