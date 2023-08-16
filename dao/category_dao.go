package dao

import (
	"database/sql"
	"example/api/models"
	"example/api/utils"
	"fmt"
)

func InsertCategory(c models.Category) error {
	var category models.Category
	// Insère la catégorie dans la base de données
	result, err := DB.Exec("INSERT INTO categories (name) VALUES(?)", category.Name)
	if err != nil {
		return err
	}
	// Obtient l'ID généré lors de l'insertion
	categoryID, err := result.LastInsertId()
	if err != nil {
		return err
	}

	// Met à jour l'ID dans l'objet User
	c.ID = int(categoryID)
	utils.Logger()
	return err
}

func SelectAllCategories() ([]models.Category, error) {
	var categories []models.Category
	// Récupère toutes les catégories depuis la base de données
	rows, err := DB.Query("SELECT id, name FROM categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var category models.Category
		err := rows.Scan(&category.ID, &category.Name)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}

func SelectCategoryByID(id int) (models.Category, error) {
	var category models.Category
	// Récupère la catégorie depuis la base de données par son ID
	err := DB.QueryRow("SELECT id, name FROM categories WHERE id = ?", id).Scan(&category.ID, &category.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Category{}, fmt.Errorf("no category found with ID %d", id)
		}
		return models.Category{}, err
	}
	return category, nil
}

func UpdateCategory(id int, updatedCategory models.Category) (models.Category, error) {
	// Met à jour la catégorie dans la base de données par son ID
	_, err := DB.Exec("UPDATE categories SET name = ?  WHERE id = ?", updatedCategory.Name, id)
	if err != nil {
		return models.Category{}, err
	}

	// Met à jour l'ID dans l'objet updatedCategory
	updatedCategory.ID = id
	return updatedCategory, nil
}

func DeleteCategory(id int) error {
	_, err := DB.Exec("DELETE FROM categories WHERE id = ?", id)
	return err
}
