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

	category := services.GetCategoryByID(result)
	if category.ID == 0 {
		http.Error(w, "Category not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(category)
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
		http.Error(w, "Category not found", http.StatusNotFound)
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
		json.NewEncoder(w).Encode(models.ErrorReponse{Message: "invalid id"})
		return
	}
	err = services.DeleteCategory(categoryID)
	if err != nil {

		http.Error(w, "category not found", http.StatusNotFound)
		return
	}

	Answer(w, http.StatusOK, "category deleted")
	//json.NewEncoder(w).Encode()
}

func Answer(w http.ResponseWriter, httpCode int, res any) {
	w.WriteHeader(httpCode)
	if e, ok := res.(error); ok {
		json.NewEncoder(w).Encode(map[string]string{"message": e.Error()})
		return
	}
	json.NewEncoder(w).Encode(map[string]interface{}{"message": res})
}
