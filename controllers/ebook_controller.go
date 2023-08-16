package controllers

import (
	"example/api/dao"
	"example/api/models"
)

// Create Ebook controller
func InsertEbook(ebook models.Ebook) error {
	err := dao.InsertEbook(ebook)
	if err != nil {
		return err
	}
	return nil
}

// Get all ebooks
func GetAllEbooks() ([]models.Ebook, error) {
	ebooks, err := dao.SelectAllEbooks()
	if err != nil {
		return nil, err
	}
	return ebooks, nil
}


// Get ebook with ID 

func GetEbookByID(id int) (models.Ebook, error) {
	ebook, err := dao.SelectEbookByID(id)
	if err != nil {
		return models.Ebook{}, err
	}
	return ebook, nil
}

// Get ebook by category ID
func GetEbooksByCategoryID(categoryID int) ([]models.Ebook, error) {
	ebooks, err := dao.SelectEbooksByCategoryID(categoryID)
	if err != nil {
		return nil, err
	}
	return ebooks, nil
}


func UpdateEbook(id int, updatedEbook models.Ebook) error {
	// Récupérer l'ebook existant par ID
	existingEbook, err := dao.SelectEbookByID(id)
	if err != nil {
		return err
	}

	// Mettre à jour les valeurs de l'ebook existant
	existingEbook.Title = updatedEbook.Title
	existingEbook.Author = updatedEbook.Author
	existingEbook.CategoryID = updatedEbook.CategoryID

	// Mettre à jour l'ebook dans la base de données
	_, err = dao.UpdateEbook(id,existingEbook)
	if err != nil {
		return err
	}

	return nil
}


func DeleteEbook(id int) error {
	err := dao.DeleteEbook(id)
	if err != nil {
		return err
	}
	return nil
}

