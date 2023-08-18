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

// CreateCategory crée une nouvelle catégorie.
func CreateCategory(w http.ResponseWriter, r *http.Request) {
	var newCategory models.Category
	err := json.NewDecoder(r.Body).Decode(&newCategory)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	services.InsertCategory(&newCategory)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newCategory)
}

// GetCategoryByID récupère une catégorie par son ID.
func GetCategoryByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	result, err := strconv.Atoi(id)
	if err != nil {
		utils.Logger(err)
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	c := services.GetCategoryByID(result)
	if c.ID == 0 {
		http.Error(w, "Category not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(c)
}

// GetAllCategories récupère toutes les catégories.
func GetAllCategories(w http.ResponseWriter, r *http.Request) {
	categories := services.GetAllCategories()
	if categories == nil {
		http.Error(w, "No data found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(categories)
}

// // UpdateCategory met à jour une catégorie par son ID.
func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	categoryID, err := strconv.Atoi(id)
	if err != nil {
		utils.Logger(err)
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var updatedCategory models.Category
	err = json.NewDecoder(r.Body).Decode(&updatedCategory)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	err = services.UpdateCategory(categoryID, &updatedCategory)
	if err != nil {
		http.Error(w, "category not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedCategory)
}

// DeleteCategory supprime une catégorie par son ID.
func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	categoryID, err := strconv.Atoi(id)
	if err != nil {
		utils.Logger(err)
		http.Error(w, "Invalid id", http.StatusInternalServerError)
		return
	}
	err = services.DeleteCategory(categoryID)
	if err != nil {
		utils.Logger(err)
		http.Error(w, "category not found", http.StatusNotFound)
	}

	w.WriteHeader(http.StatusOK)
	//json.NewEncoder(w).Encode()
}
