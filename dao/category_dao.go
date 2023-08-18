package dao

import (
	"database/sql"
	"errors"
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
	rows, err := DB.Query("SELECT id, name FROM categories ORDER BY id")
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
	// Récupère la catégorie depuis la base de données par son ID
	var category models.Category
	err := DB.QueryRow("SELECT id, name FROM categories WHERE id = $1", id).Scan(&category.ID, &category.Name)

	if errors.Is(err, sql.ErrNoRows) {
		return models.Category{}, fmt.Errorf("no category found with ID %d", id)
	}
	if err != nil {
		utils.Logger(err)
		return models.Category{}, err
	}
	return category, nil
}

func UpdateCategory(updatedCategory models.Category) (models.Category, error) {
	// Met à jour la catégorie dans la base de données par son ID
	_, err := DB.Exec("UPDATE categories SET name = $1  WHERE id = $2", updatedCategory.Name, updatedCategory.ID)
	if err != nil {
		utils.Logger(err)
		return models.Category{}, err
	}

	// Met à jour l'ID dans l'objet updatedCategory
	
	return updatedCategory, nil
}

func DeleteCategory(id int) error {
	_, err := DB.Exec("DELETE FROM categories WHERE id = $1", id)
	if err != nil {
		utils.Logger(err)
	}
	return err
}
