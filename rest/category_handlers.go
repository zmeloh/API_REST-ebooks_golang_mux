package rest

import (
	"encoding/json"
	"example/api/dao"
	"example/api/models"
	"example/api/services"
	"example/api/utils"
	"fmt"
	"net/http"
)

// CreateCategory crée une nouvelle catégorie.
func CreateCategory(w http.ResponseWriter, r *http.Request) {
	// Décode le JSON de la requête dans une instance de Category
	var category models.Category
	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		utils.Logger(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	services.InsertCategory(&category)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(category)
}

// GetCategoryByID récupère une catégorie par son ID.
func GetCategoryByID(w http.ResponseWriter, r *http.Request) {
	var category models.Category
	dao.SelectCategoryByID(category.ID)
	json.NewEncoder(w).Encode(category)
}

// GetAllCategories récupère toutes les catégories.
func GetAllCategories(w http.ResponseWriter, r *http.Request) {
	var categories []models.Category
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(categories)
	if err != nil {
		//utils.Logger(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	c := services.GetAllCategories()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(c)
	
}

// UpdateCategory met à jour une catégorie par son ID.
func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	var updatedCategory models.Category
	w.Header().Set("Content-Type", "application/json")
	//dao.UpdateCategory(updatedCategory.ID)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedCategory)
}

// DeleteCategory supprime une catégorie par son ID.
func DeleteCategory(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category has been deleted")
}
