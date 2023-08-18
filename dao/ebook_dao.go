package dao

import (
	"database/sql"
	"errors"
	"example/api/models"
	"example/api/utils"
	"fmt"
)

// Create Ebook
func InsertEbook(e *models.Ebook) error {
	// Requête pour insérer un nouvel ebook dans la base de données
	err := DB.QueryRow("INSERT INTO ebooks (title, author, category_id) VALUES ($1, $2, $3) RETURNING id", e.Title, e.Author, e.CategoryID).Scan(&e.ID)
	if err != nil {
		utils.Logger(err)
		return err
	}

	return err
}

// Select Ebooks
func SelectAllEbooks() ([]models.Ebook, error) {
	var ebooks []models.Ebook

	// Requête pour récupérer tous les livres électroniques depuis la base de données
	rows, err := DB.Query("SELECT id, title, author, category_id FROM ebooks ORDER BY id")
	if err != nil {
		utils.Logger(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var ebook models.Ebook
		err := rows.Scan(&ebook.ID, &ebook.Title, &ebook.Author, &ebook.CategoryID)
		if err != nil {
			return nil, err
		}
		ebooks = append(ebooks, ebook)
	}

	return ebooks, nil
}

// Select Ebooks with ID
func SelectEbookByID(id int) (models.Ebook, error) {
	var ebook models.Ebook
	// Requête pour récupérer le livre électronique par ID depuis la base de données
	err := DB.QueryRow("SELECT id, title, author, category_id FROM ebooks WHERE id = $1", id).Scan(&ebook.ID, &ebook.Title, &ebook.Author, &ebook.CategoryID)
	if errors.Is(err, sql.ErrNoRows) {
		return models.Ebook{}, fmt.Errorf("no ebook found with ID %d", id)
	}
	if err != nil {
		utils.Logger(err)
		return models.Ebook{}, err
	}
	return ebook, nil
}

// Update Ebook
func UpdateEbook(updatedEbook models.Ebook) (models.Ebook, error) {
	// Requête pour mettre à jour les informations du livre électronique dans la base de données
	_, err := DB.Exec("UPDATE ebooks SET title = $1, author = $2, category_id = $3 WHERE id = $4", updatedEbook.Title, updatedEbook.Author,updatedEbook.CategoryID, updatedEbook.ID)
	if err != nil {
		utils.Logger(err)
		return models.Ebook{}, err
	}

	return updatedEbook, nil

}

// Delete Ebook
func DeleteEbook(id int) error {
	// Requête pour supprimer un livre électronique par ID dans la base de données
	err := DB.QueryRow("DELETE FROM ebooks WHERE id = $1 RETURNING ID", id).Scan(&id)
	if err != nil {
		utils.Logger(err)
	}
	return err
}

// Select Ebooks with category_id
func SelectEbooksByCategoryID(categoryID int) ([]models.Ebook, error) {
	var ebooks []models.Ebook

	// Requête pour récupérer les livres électroniques par ID de catégorie depuis la base de données
	rows, err := DB.Query("SELECT id, title, author, category_id FROM ebooks WHERE category_id = ?", categoryID)
	if err != nil {
		utils.Logger(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var ebook models.Ebook
		err := rows.Scan(&ebook.ID, &ebook.Title, &ebook.Author, &ebook.CategoryID)
		if err != nil {
			return nil, err
		}
		ebooks = append(ebooks, ebook)
	}

	return ebooks, nil
}
