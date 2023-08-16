package controllers

import (
	"example/api/dao"
	"example/api/models"
)

func InsertCategory(c models.Category) error {
	err := dao.InsertCategory(c)
	if err != nil {
		return err
	}
	return nil
}

func GetAllCategories() ([]models.Category, error) {
	categories, err := dao.SelectAllCategories()
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func GetCategoryByID(id int) (models.Category, error) {
	category, err := dao.SelectCategoryByID(id)
	if err != nil {
		return models.Category{}, err
	}
	return category, nil
}

func UpdateCategory(id int, updatedCategory models.Category) error {

	existingCategory, err := dao.SelectCategoryByID(id)
	if err != nil {
		return err
	}

	// Mettre à jour les valeurs de l'ebook existant
	existingCategory.Name = updatedCategory.Name

	_, err = dao.UpdateCategory(id, existingCategory)
	if err != nil {
		return err
	}

	return nil
}

func DeleteCategory(id int) error {
	err := dao.DeleteCategory(id)
	if err != nil {
		return err
	}
	return err
}
