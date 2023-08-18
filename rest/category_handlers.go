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
// func CreateCategory(w http.ResponseWriter, r *http.Request) {
// 	// Décode le JSON de la requête dans une instance de Category
// 	var category models.Category
// 	err := json.NewDecoder(r.Body).Decode(&category)
// 	if err != nil {
// 		utils.Logger(err)
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	services.InsertCategory(&category)

// 	w.WriteHeader(http.StatusCreated)
// 	json.NewEncoder(w).Encode(category)
// }

// GetCategoryByID récupère une catégorie par son ID.
// func GetCategoryByID(w http.ResponseWriter, r *http.Request) {
//     params := mux.Vars(r)
//     id := params["id"]
//     result, err := strconv.Atoi(id)
//     if err != nil {
//         utils.Logger(err)
//         http.Error(w, "Invalid ID", http.StatusBadRequest)
//         return
//     }

//     c := services.GetCategoryByID(result)
//     if c.ID == 0 { // Remplacez cette condition par celle qui indique que la catégorie n'a pas été trouvée
//         http.Error(w, "Category not found", http.StatusNotFound)
//         return
//     }

//     w.WriteHeader(http.StatusOK)
//     json.NewEncoder(w).Encode(c)
// }

// GetAllCategories récupère toutes les catégories.
// func GetAllCategories(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	category := services.GetAllCategories()
// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(category)

// }

// // UpdateCategory met à jour une catégorie par son ID.
// func UpdateCategory(w http.ResponseWriter, r *http.Request) {
// 	var updatedCategory models.Category
// 	w.Header().Set("Content-Type", "application/json")
// 	services.UpdateCategory(updatedCategory.ID, updatedCategory)
// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(updatedCategory)
// }

// DeleteCategory supprime une catégorie par son ID.
// func DeleteCategory(w http.ResponseWriter, r *http.Request) {
// 	var category models.Category
// 	services.DeleteCategory(category.ID)
// 	w.WriteHeader(http.StatusOK)
// 	fmt.Fprintf(w, "Category has been deleted")
// }

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	var newCategory models.Category
	err := json.NewDecoder(r.Body).Decode(&newCategory)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	services.InsertCategory(&newCategory)
	// if err != nil {
	// 	utils.Logger(err)
	// 	http.Error(w, "Error inserting category", http.StatusInternalServerError)
	// 	return
	// }

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newCategory)
}

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

func GetAllCategories(w http.ResponseWriter, r *http.Request) {
	categories := services.GetAllCategories()
	if categories == nil {
		http.Error(w, "No data found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(categories)
}


func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	categoryID, err := strconv.Atoi(id)
	// if err != nil {
	// 	utils.Logger(err)
	// 	http.Error(w, "Invalid ID", http.StatusBadRequest)
	// 	return
	// }
	services.GetCategoryByID(categoryID)

	var updatedCategory models.Category
	err = json.NewDecoder(r.Body).Decode(&updatedCategory)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	services.UpdateCategory(categoryID, updatedCategory)
	// if err != nil {
	// 	utils.Logger(err)
	// 	http.Error(w, "Error updating category", http.StatusInternalServerError)
	// 	return
	// }
	
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedCategory)
}

func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	categoryID, err := strconv.Atoi(id)
	if err != nil {
		utils.Logger(err)
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	services.DeleteCategory(categoryID)
	if err != nil {
		utils.Logger(err)
		http.Error(w, "Error deleting category", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode()
}

