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


func CreateCategory(w http.ResponseWriter, r *http.Request){
	var category models.Category

	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	_, err = database.DB.Exec("INSERT INTO categories (name) VALUES(?)", category.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(category)
}


func GetCategoryByID(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	categoryID := vars["id"]

	var category models.Category

	err := database.DB.QueryRow("SELECT id, name FROM categories WHERE id = ?", categoryID).Scan(&category.ID, &category.Name)
	if err != nil {
		if err == sql.ErrNoRows{
			http.NotFound(w,r)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(category)
}


func GetAllCategories(w http.ResponseWriter, r *http.Request){
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

func UpdateCategory(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	categoryID := vars["id"]

	var updatedCategory models.Category
	err := json.NewDecoder(r.Body).Decode(&updatedCategory)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = database.DB.Exec("UPDATE categories SET name = ?  WHERE id = ?", updatedCategory.Name, categoryID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(updatedCategory)
}

func DeleteCategory(w http.ResponseWriter, r *http.Request){
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