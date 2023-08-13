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

// CreateCategory crée une nouvelle catégorie.
func CreateCategory(w http.ResponseWriter, r *http.Request) {
	// Décode le JSON de la requête dans une instance de Category
	var category models.Category
	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Insère la catégorie dans la base de données
	result, err := database.DB.Exec("INSERT INTO categories (name) VALUES(?)", category.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Récupère l'ID généré par la base de données
	categoryID, err := result.LastInsertId()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Met à jour l'ID dans l'objet Category
	category.ID = int(categoryID)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(category)
}

// GetCategoryByID récupère une catégorie par son ID.
func GetCategoryByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	categoryID := vars["id"]

	var category models.Category

	// Récupère la catégorie depuis la base de données par son ID
	err := database.DB.QueryRow("SELECT id, name FROM categories WHERE id = ?", categoryID).Scan(&category.ID, &category.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			http.NotFound(w, r)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(category)
}

// GetAllCategories récupère toutes les catégories.
func GetAllCategories(w http.ResponseWriter, r *http.Request) {
	// Récupère toutes les catégories depuis la base de données
	rows, err := database.DB.Query("SELECT id, name FROM categories")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var categories []models.Category
	for rows.Next() {
		var category models.Category
		err := rows.Scan(&category.ID, &category.Name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		categories = append(categories, category)
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(categories)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// UpdateCategory met à jour une catégorie par son ID.
func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	categoryID := vars["id"]

	var updatedCategory models.Category
	err := json.NewDecoder(r.Body).Decode(&updatedCategory)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Met à jour la catégorie dans la base de données par son ID
	_, err = database.DB.Exec("UPDATE categories SET name = ?  WHERE id = ?", updatedCategory.Name, categoryID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Convertit l'ID de string à int
	categoryIDInt, err := strconv.Atoi(categoryID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Met à jour l'ID dans l'objet updatedCategory
	updatedCategory.ID = categoryIDInt

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedCategory)
}

// DeleteCategory supprime une catégorie par son ID.
func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	categoryID := vars["id"]
	_, err := database.DB.Exec("DELETE FROM categories WHERE id = ?", categoryID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category with ID %s has been deleted", categoryID)
}
