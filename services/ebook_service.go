package services

import (
	"example/api/controllers" 
	"example/api/models"      
)

// GetAllEbooks 
func GetAllEbooks(){
	controllers.GetAllEbooks() 
	return
}

// GetEbookByID 
func GetEbookByID(id int){
	controllers.GetEbookByID(id) 
	return
}

// GetEbookByID 
func GetEbookByCategoryID(id int){
	controllers.GetEbooksByCategoryID(id) 
	return
}

// InsertEbook
func InsertEbook(e models.Ebook) {
	controllers.InsertEbook(e) 
	return
}

// UpdateEbook 
func UpdateEbook(id int, e models.Ebook) {
	controllers.UpdateEbook(id, e) 
	return
}

// DeleteEbook 
func DeleteEbook(id int) {
	controllers.DeleteEbook(id) 
	return
}
