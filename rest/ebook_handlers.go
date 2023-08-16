package rest

import (
	"encoding/json"
	"example/api/models"
	"example/api/services"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// CreateEbook crée un nouveau livre électronique.
func CreateEbook(w http.ResponseWriter, r *http.Request) {
	var newEbook models.Ebook
	if err := json.NewDecoder(r.Body).Decode(&newEbook); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Appeler le service pour insérer le nouvel ebook
	services.InsertEbook(newEbook)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newEbook)
}

// GetAllEbook récupère la liste de tous les livres électroniques.
func GetAllEbooks(w http.ResponseWriter, r *http.Request) {
	// Appeler le service pour récupérer tous les livres électroniques
	ebooks := services.GetAllEbooks()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ebooks)
}

// GetEbookByID récupère un livre électronique par son ID.
func GetEbookByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ebookID := vars["id"]
	id, err := strconv.Atoi(ebookID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Appeler le service pour récupérer un livre électronique par son ID
	ebook := services.GetEbookByID(id)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ebook)
}

// GetEbookByCategoryID récupère la liste de tous les livres électroniques d'une catégorie.
func GetEbookByCategoryID(w http.ResponseWriter, r *http.Request) {
	// Récupérer l'ID de la catégorie depuis les paramètres de la requête
	categoryIDStr := r.URL.Query().Get("category_id")
	categoryID, err := strconv.Atoi(categoryIDStr)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Appeler le service pour récupérer les livres électroniques de la catégorie
	ebooks := services.GetEbookByCategoryID(categoryID)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ebooks)
}

// UpdateEbook met à jour un livre électronique par son ID.
func UpdateEbook(w http.ResponseWriter, r *http.Request) {
	// Récupérer l'ID de l'ebook depuis les variables de la requête
	vars := mux.Vars(r)
	ebookID := vars["id"]
	id, err := strconv.Atoi(ebookID)

	// Gérer les erreurs liées à la conversion de l'ID
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Déclarer une variable pour stocker les données de l'ebook mis à jour
	var updatedEbook models.Ebook

	// Décoder le JSON du corps de la requête pour obtenir les nouvelles données de l'ebook
	if err := json.NewDecoder(r.Body).Decode(&updatedEbook); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Appeler le service pour mettre à jour l'ebook en utilisant l'ID et les nouvelles données
	services.UpdateEbook(id, updatedEbook)

	// Préparer la réponse JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedEbook)
}

// DeleteEbook supprime un livre électronique par son ID.
func DeleteEbook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ebookID := vars["id"]
	id, err := strconv.Atoi(ebookID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Appeler le service pour supprimer l'ebook
	services.DeleteEbook(id)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}
