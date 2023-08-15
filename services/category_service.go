package services

import (
	"example/api/controllers"
	"example/api/models"
)

// GetAllCategories 
func GetAllCategories(){
	controllers.GetAllCategories()
	return
}

// GetCategoryByID 
func GetCategoryByID(id int){
	controllers.GetCategoryByID(id)
	return
}

// InsertCategory 
func InsertCategory(c models.Category){
	//controllers.CreateEbook(c)
	return
}

// UpdateCategory 
func UpdateCategory(id int, c models.Category){
	controllers.UpdateCategory(id, c)
	return
}

// DeleteCategory 
func DeleteCategory(id int){
	controllers.DeleteCategory(id)
	return
}
