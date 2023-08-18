package controllers

import (
	"example/api/dao"
	"example/api/models"
	"example/api/utils"
)

// Create Ebook controller
func InsertEbook(e *models.Ebook) error {
	err := dao.InsertEbook(e)
	if err != nil {
		utils.Logger(err)
		return err
	}
	return nil
}

// Get all ebooks
func GetAllEbooks() ([]models.Ebook, error) {
	ebooks, err := dao.SelectAllEbooks()
	if err != nil {
		utils.Logger(err)
		return nil, err
	}
	return ebooks, nil
}

// Get ebook with ID

func GetEbookByID(id int) (models.Ebook, error) {
	ebook, err := dao.SelectEbookByID(id)
	if err != nil {
		utils.Logger(err)
		return models.Ebook{}, err
	}
	return ebook, nil
}

// Get ebook by category ID
func GetEbooksByCategoryID(categoryID int) ([]models.Ebook, error) {
	ebooks, err := dao.SelectEbooksByCategoryID(categoryID)
	if err != nil {
		utils.Logger(err)
		return nil, err
	}
	return ebooks, nil
}

func UpdateEbook(id int, updatedEbook *models.Ebook) error {
	// Récupérer l'ebook existant par ID
	existingEbook, err := dao.SelectEbookByID(id)
	if err != nil {
		utils.Logger(err)
		return err
	}

	// Mettre à jour les valeurs de l'ebook existant
	existingEbook.Title = updatedEbook.Title
	existingEbook.Author = updatedEbook.Author
	existingEbook.CategoryID = updatedEbook.CategoryID
	updatedEbook.ID = existingEbook.ID

	// Mettre à jour l'ebook dans la base de données
	_, err = dao.UpdateEbook(existingEbook)
	if err != nil {
		utils.Logger(err)
		return err
	}

	return nil
}

func DeleteEbook(id int) error {
	err := dao.DeleteEbook(id)
	if err != nil {
		utils.Logger(err)
		return err
	}
	return nil
}
