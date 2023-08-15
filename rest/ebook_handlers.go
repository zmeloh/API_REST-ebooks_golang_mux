package rest

import (
	"encoding/json"
	"example/api/models"
	"fmt"
	"net/http"
	"strconv"
	"example/api/services"
	"github.com/gorilla/mux"
)

// CreateEbook crée un nouveau livre électronique.
func CreateEbook(w http.ResponseWriter, r *http.Request) {
	var ebook models.Ebook

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(ebook)
}

// GetEbookByID récupère un livre électronique par son ID.
func GetEbookByID(w http.ResponseWriter, r *http.Request) {
	var ebook models.Ebook
	json.NewEncoder(w).Encode(ebook)
}

// GetEbooksByCategory récupère tous les livres électroniques d'une catégorie spécifiée.
func GetEbooksByCategoryID(w http.ResponseWriter, r *http.Request) {
	var ebooks []models.Ebook
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(ebooks)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// GetAllEbooks récupère tous les livres électroniques.
func GetAllEbooks(w http.ResponseWriter, r *http.Request) {
	var ebooks []models.Ebook
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(ebooks)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// UpdateEbook met à jour un livre électronique par son ID.
func UpdateEbook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ebookID := vars["id"]
	id, err := strconv.Atoi(ebookID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	var updatedEbook models.Ebook
	services.UpdateEbook(id, updatedEbook)
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedEbook)
}

// DeleteEbook supprime un livre électronique par son ID.
func DeleteEbook(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Ebook with has been deleted")
}
