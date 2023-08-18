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
		ServerResponse(w, http.StatusBadRequest, "Invalid request")
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
		ServerResponse(w, http.StatusBadRequest, "Invalid id")
		return
	}

	category := services.GetCategoryByID(result)
	if category.ID == 0 {

		ServerResponse(w, http.StatusNotFound, "Category not found")
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(category)
}

// GetAllCategories récupère toutes les catégories.
func GetAllCategories(w http.ResponseWriter, r *http.Request) {
	categories := services.GetAllCategories()
	if categories == nil {
		ServerResponse(w, http.StatusNotFound, "No data found")
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
		ServerResponse(w, http.StatusBadRequest, "Invalid id")
		return
	}

	var updatedCategory models.Category
	err = json.NewDecoder(r.Body).Decode(&updatedCategory)
	if err != nil {
		ServerResponse(w, http.StatusBadRequest, "Category not found")
		return
	}

	err = services.UpdateCategory(categoryID, &updatedCategory)
	if err != nil {
		ServerResponse(w, http.StatusNotFound, "Category not found")
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
		ServerResponse(w, http.StatusBadRequest, "Invalid id")
		return
	}
	err = services.DeleteCategory(categoryID)
	if err != nil {
		ServerResponse(w, http.StatusNotFound, "Category not found")
		return
	}

	ServerResponse(w, http.StatusOK, "Category has been deleted")
}
