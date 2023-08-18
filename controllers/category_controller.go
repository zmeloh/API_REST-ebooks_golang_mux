package controllers

import (
	"example/api/dao"
	"example/api/models"
	"example/api/utils"
)

func InsertCategory(c *models.Category) error {
	err := dao.InsertCategory(c)
	if err != nil {
		utils.Logger(err)
		return err
	}
	return nil
}

func GetAllCategories() ([]models.Category, error) {
	categories, err := dao.SelectAllCategories()
	if err != nil {
		utils.Logger(err)
		return nil, err
	}
	return categories, nil
}

func GetCategoryByID(id int) (models.Category, error) {
	category, err := dao.SelectCategoryByID(id)
	if err != nil {
		utils.Logger(err)
		return models.Category{}, err
	}
	return category, nil
}

func UpdateCategory(id int, updatedCategory models.Category) error {

	existingCategory, err := dao.SelectCategoryByID(id)
	if err != nil {
		utils.Logger(err)
		return err
	}

	// Mettre à jour les valeurs de l'ebook existant
	existingCategory.Name = updatedCategory.Name

	updatedCategory, err = dao.UpdateCategory(id, existingCategory)
	if err != nil {
		utils.Logger(err)
		return err
	}

	return nil
}

func DeleteCategory(id int) error {
	err := dao.DeleteCategory(id)
	if err != nil {
		utils.Logger(err)
		return err
	}
	return err
}

/*
func InsertCategory(w http.ResponseWriter, r *http.Request) {
	var newCategory models.Category
	err := json.NewDecoder(r.Body).Decode(&newCategory)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err = services.InsertCategory(&newCategory)
	if err != nil {
		utils.Logger(err)
		http.Error(w, "Error inserting category", http.StatusInternalServerError)
		return
	}

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
	categories, err := services.GetAllCategories()
	if err != nil {
		utils.Logger(err)
		http.Error(w, "Error fetching categories", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(categories)
}

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
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err = services.UpdateCategory(categoryID, updatedCategory)
	if err != nil {
		utils.Logger(err)
		http.Error(w, "Error updating category", http.StatusInternalServerError)
		return
	}

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

	err = services.DeleteCategory(categoryID)
	if err != nil {
		utils.Logger(err)
		http.Error(w, "Error deleting category", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

type ErrorResponse struct {
    Message string `json:"message"`
}

errorResponse := ErrorResponse{
	Message: err.Error(),
}
*/
