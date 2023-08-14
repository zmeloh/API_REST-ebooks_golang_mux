package dao

import (
	"database/sql"
	"encoding/json"
	"example/api/models"
	//"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func insertCategory(w http.ResponseWriter, r *http.Request) {
	var category models.Category
	// Insère la catégorie dans la base de données
	result, err := DB.Exec("INSERT INTO categories (name) VALUES(?)", category.Name)
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
}

func selectAllCategories(w http.ResponseWriter, r *http.Request) {
	// Récupère toutes les catégories depuis la base de données
	rows, err := DB.Query("SELECT id, name FROM categories")
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

}

func selectCategoryByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	categoryID := vars["id"]

	var category models.Category

	// Récupère la catégorie depuis la base de données par son ID
	err := DB.QueryRow("SELECT id, name FROM categories WHERE id = ?", categoryID).Scan(&category.ID, &category.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			http.NotFound(w, r)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func updateCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	categoryID := vars["id"]

	var updatedCategory models.Category
	err := json.NewDecoder(r.Body).Decode(&updatedCategory)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Met à jour la catégorie dans la base de données par son ID
	_, err = DB.Exec("UPDATE categories SET name = ?  WHERE id = ?", updatedCategory.Name, categoryID)
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
}

func deleteCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	categoryID := vars["id"]
	_, err := DB.Exec("DELETE FROM categories WHERE id = ?", categoryID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
