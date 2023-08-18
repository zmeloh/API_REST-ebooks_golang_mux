package rest

import (
	"encoding/json"
	"example/api/models"
	"example/api/services"
	"example/api/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// CreateEbook crée un nouveau livre électronique.
func CreateEbook(w http.ResponseWriter, r *http.Request) {
	var newEbook models.Ebook
	err := json.NewDecoder(r.Body).Decode(&newEbook)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Appeler le service pour insérer le nouvel ebook
	services.InsertEbook(&newEbook)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newEbook)
}

// GetAllEbook récupère la liste de tous les livres électroniques.
func GetAllEbooks(w http.ResponseWriter, r *http.Request) {
	ebooks := services.GetAllEbooks()
	if ebooks == nil {
		http.Error(w, "No data found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ebooks)
}

// GetEbookByID récupère un livre électronique par son ID.
func GetEbookByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	result, err := strconv.Atoi(id)

	if err != nil {
		utils.Logger(err)
		http.Error(w, "Invald request", http.StatusBadRequest)
		return
	}

	// Appeler le service pour récupérer un livre électronique par son ID
	ebook := services.GetEbookByID(result)
	if ebook.ID == 0 {
		http.Error(w, "Ebook not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ebook)
}

// GetEbookByCategoryID récupère la liste de tous les livres électroniques d'une catégorie.
func GetEbookByCategoryID(w http.ResponseWriter, r *http.Request) {
	// Récupérer l'ID de la catégorie depuis les paramètres de la requête
	categoryIDStr := r.URL.Query().Get("category_id")
	categoryID, err := strconv.Atoi(categoryIDStr)

	if err != nil {
		utils.Logger(err)
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
	params := mux.Vars(r)
	id := params["id"]
	ebookID, err := strconv.Atoi(id)

	if err != nil {
		utils.Logger(err)
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var updatedEbook models.Ebook
	err = json.NewDecoder(r.Body).Decode(&updatedEbook)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	err = services.UpdateEbook(ebookID, &updatedEbook)
	if err != nil {
		http.Error(w, "Ebook not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedEbook)
}

// DeleteEbook supprime un livre électronique par son ID.
func DeleteEbook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ebookID := vars["id"]
	id, err := strconv.Atoi(ebookID)

	if err != nil {
		utils.Logger(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Appeler le service pour supprimer l'ebook
	services.DeleteEbook(id)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}
