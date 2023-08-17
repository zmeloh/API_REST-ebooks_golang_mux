package dao

import (
	"database/sql"
	"example/api/models"
	"example/api/utils"
	"fmt"
)

func InsertCategory(c *models.Category) error {
	// Insère la catégorie dans la base de données
	err := DB.QueryRow("INSERT INTO categories (name) VALUES($1) RETURNING id", c.Name).Scan(&c.ID)
	if err != nil {
		utils.Logger(err)
		return err
	}

	return err
}

func SelectAllCategories() ([]models.Category, error) {
	var categories []models.Category
	// Récupère toutes les catégories depuis la base de données
	rows, err := DB.Query("SELECT id, name FROM categories")
	if err != nil {
		utils.Logger(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var category models.Category
		err := rows.Scan(&category.ID, &category.Name)
		if err != nil {
			utils.Logger(err)
			return nil, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}

func SelectCategoryByID(id int) (models.Category, error) {
	var category models.Category
	// Récupère la catégorie depuis la base de données par son ID
	err := DB.QueryRow("SELECT id, name FROM categories WHERE id = $1",id).Scan(category.ID, category.Name)
	if err != nil {
		utils.Logger(err)
		if err == sql.ErrNoRows {
			return models.Category{}, fmt.Errorf("no category found with ID %d", category.ID)
		}
		return models.Category{}, err
	}
	return category, nil
}

func UpdateCategory(id int, updatedCategory models.Category) (models.Category, error) {
	// Met à jour la catégorie dans la base de données par son ID
	_, err := DB.Exec("UPDATE categories SET name = ?  WHERE id = ?", updatedCategory.Name, id)
	if err != nil {
		utils.Logger(err)
		return models.Category{}, err
	}

	// Met à jour l'ID dans l'objet updatedCategory
	updatedCategory.ID = id
	return updatedCategory, nil
}

func DeleteCategory(id int) error {
	_, err := DB.Exec("DELETE FROM categories WHERE id = $1", id)
	if err != nil {
		utils.Logger(err)
	}
	return err
}
