package handlers

import (
	"encoding/json"
	//"example/api/dao"
	"example/api/models"
	"fmt"
	"net/http"
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


	// Répond avec le favori créé et le code de statut 201 (Created)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(favorite)
}

// GetAllFavorites récupère tous les favoris.
func GetAllFavorites(w http.ResponseWriter, r *http.Request) {
	var favorites []models.Favorite

	// Configure l'en-tête Content-Type pour la réponse JSON
	w.Header().Set("Content-Type", "application/json")
	// Encode la slice de favoris en JSON et répond avec les favoris
	err := json.NewEncoder(w).Encode(favorites)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// GetFavoriteByID récupère un favori par son ID.
func GetFavoriteByID(w http.ResponseWriter, r *http.Request) {
	var favorite models.Favorite
	// Encode le favori en JSON et répond avec le favori
	json.NewEncoder(w).Encode(favorite)
}

// GetFavoritesByUserID récupère les favoris par ID d'utilisateur.
func GetFavoritesByUserID(w http.ResponseWriter, r *http.Request) {
	var favorites []models.Favorite
	// Configure l'en-tête Content-Type pour la réponse JSON
	w.Header().Set("Content-Type", "application/json")
	// Encode la slice de favoris en JSON et répond avec les favoris
	err := json.NewEncoder(w).Encode(favorites)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// GetFavoritesByEbookID récupère les favoris par ID d'ebook.
func GetFavoritesByEbookID(w http.ResponseWriter, r *http.Request) {
	var favorites []models.Favorite
	// Configure l'en-tête Content-Type pour la réponse JSON
	w.Header().Set("Content-Type", "application/json")
	// Encode la slice de favoris en JSON et répond avec les favoris
	err := json.NewEncoder(w).Encode(favorites)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// UpdateFavorite met à jour un favori par son ID.
func UpdateFavorite(w http.ResponseWriter, r *http.Request) {
	var updatedFavorite models.Favorite
	// Configure l'en-tête Content-Type pour la réponse JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// Encode le favori mis à jour en JSON et répond avec le favori
	json.NewEncoder(w).Encode(updatedFavorite)
}

// DeleteFavorite supprime un favori par son ID.
func DeleteFavorite(w http.ResponseWriter, r *http.Request) {
	
	// Répond avec le statut 200 (OK) et un message indiquant que le favori a été supprimé
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Favorite has been deleted")
}
