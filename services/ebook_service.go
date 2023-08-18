package services

import (
	"example/api/controllers"
	"example/api/models"
)

// GetAllEbooks
func GetAllEbooks() (e []models.Ebook) {
	e, _ = controllers.GetAllEbooks()
	return
}

// GetEbookByID
func GetEbookByID(id int) (e models.Ebook) {
	e, _ = controllers.GetEbookByID(id)
	return
}

// GetEbookByID
func GetEbookByCategoryID(id int) (e []models.Ebook) {
	e, _ = controllers.GetEbooksByCategoryID(id)
	return
}

// InsertEbook
func InsertEbook(e *models.Ebook) {
	controllers.InsertEbook(e)
	return
}

// UpdateEbook
func UpdateEbook(id int, e *models.Ebook) error {
	err := controllers.UpdateEbook(id, e)
	return err
}

// DeleteEbook
func DeleteEbook(id int) error {
	err := controllers.DeleteEbook(id)
	return err
}
