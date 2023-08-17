package services

import (
	"example/api/controllers"
	"example/api/models"
)

// GetAllCategories
func GetAllCategories() (c []models.Category) {
	c, _ =controllers.GetAllCategories()
	return
}

// GetCategoryByID
func GetCategoryByID(id int) (c models.Category) {
	controllers.GetCategoryByID(id)
	return
}

// InsertCategory
func InsertCategory(c *models.Category) {
	controllers.InsertCategory(c)
	return
}

// UpdateCategory
func UpdateCategory(id int, c models.Category) {
	controllers.UpdateCategory(id, c)
	return
}

// DeleteCategory
func DeleteCategory(id int) {
	controllers.DeleteCategory(id)
	return
}
