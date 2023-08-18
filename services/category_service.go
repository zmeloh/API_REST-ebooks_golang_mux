package services

import (
	"example/api/controllers"
	"example/api/models"
)

// GetAllCategories
func GetAllCategories() (c []models.Category) {
	c, _ = controllers.GetAllCategories()
	return
}

// GetCategoryByID
func GetCategoryByID(id int) (c models.Category) {
	c, _ = controllers.GetCategoryByID(id)
	return
}

// InsertCategory
func InsertCategory(c *models.Category) {
	controllers.InsertCategory(c)
	return
}

// UpdateCategory
func UpdateCategory(id int, c *models.Category) error {
	err := controllers.UpdateCategory(id, c)
	return err
}

// DeleteCategory
func DeleteCategory(id int) error {
	err := controllers.DeleteCategory(id)
	return err
}
