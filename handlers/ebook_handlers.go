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

// CreateEbook crée un nouveau livre électronique.
func CreateEbook(w http.ResponseWriter, r *http.Request) {
	var ebook models.Ebook
	err := json.NewDecoder(r.Body).Decode(&ebook)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Insérer le livre électronique dans la base de données
	result, err := database.DB.Exec("INSERT INTO ebooks (title, author, category_id) VALUES (?, ?, ?)", ebook.Title, ebook.Author, ebook.CategoryID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Récupérer l'ID généré pour le livre électronique
	ebookID, err := result.LastInsertId()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Mettre à jour l'ID dans l'objet Ebook
	ebook.ID = int(ebookID)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(ebook)
}

// GetEbookByID récupère un livre électronique par son ID.
func GetEbookByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ebookID := vars["id"]

	var ebook models.Ebook

	// Requête pour récupérer le livre électronique par ID depuis la base de données
	err := database.DB.QueryRow("SELECT id, title, author, category_id FROM ebooks WHERE id = ?", ebookID).Scan(&ebook.ID, &ebook.Title, &ebook.Author, &ebook.CategoryID)
	if err != nil {
		if err == sql.ErrNoRows {
			http.NotFound(w, r)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(ebook)
}

// GetEbooksByCategory récupère tous les livres électroniques d'une catégorie spécifiée.
func GetEbooksByCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	categoryID := vars["categoryID"]

	// Requête pour récupérer tous les livres électroniques d'une catégorie depuis la base de données
	rows, err := database.DB.Query("SELECT id, title, author, category_id FROM ebooks WHERE category_id = ?", categoryID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var ebooks []models.Ebook
	for rows.Next() {
		var ebook models.Ebook
		err := rows.Scan(&ebook.ID, &ebook.Title, &ebook.Author, &ebook.CategoryID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		ebooks = append(ebooks, ebook)
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(ebooks)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// GetAllEbooks récupère tous les livres électroniques.
func GetAllEbooks(w http.ResponseWriter, r *http.Request) {
	// Requête pour récupérer tous les livres électroniques depuis la base de données
	rows, err := database.DB.Query("SELECT id, title, author, category_id FROM ebooks")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var ebooks []models.Ebook
	for rows.Next() {
		var ebook models.Ebook
		err := rows.Scan(&ebook.ID, &ebook.Title, &ebook.Author, &ebook.CategoryID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		ebooks = append(ebooks, ebook)
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(ebooks)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// UpdateEbook met à jour un livre électronique par son ID.
func UpdateEbook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ebookID := vars["id"]

	var updatedEbook models.Ebook
	err := json.NewDecoder(r.Body).Decode(&updatedEbook)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Requête pour mettre à jour les informations du livre électronique dans la base de données
	_, err = database.DB.Exec("UPDATE ebooks SET title = ?, author = ?, category_id = ? WHERE id = ?", updatedEbook.Title, updatedEbook.Author, updatedEbook.CategoryID, ebookID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Convertir l'ID en entier
	ebookIDInt, err := strconv.Atoi(ebookID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Mettre à jour l'ID dans l'objet Ebook
	updatedEbook.ID = ebookIDInt

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedEbook)
}

// DeleteEbook supprime un livre électronique par son ID.
func DeleteEbook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ebookID := vars["id"]

	// Requête pour supprimer le livre électronique depuis la base de données
	_, err := database.DB.Exec("DELETE FROM ebooks WHERE id = ?", ebookID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Ebook with ID %s has been deleted", ebookID)
}
