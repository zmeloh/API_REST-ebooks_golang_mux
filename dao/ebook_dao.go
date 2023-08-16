package dao

import (
	"database/sql"
	"example/api/models"
	"example/api/utils"
	"fmt"
)

// Create Ebook
func InsertEbook(e models.Ebook) error {
	// Requête pour insérer un nouvel ebook dans la base de données
	result, err := DB.Exec("INSERT INTO ebooks (title, author, category_id) VALUES ($1, $2, $3)", e.Title, e.Author, e.CategoryID)
	if err != nil {
		return err
	}
	// Obtient l'ID généré lors de l'insertion
	ebookID, err := result.LastInsertId()
	if err != nil {
		return err
	}

	// Met à jour l'ID dans l'objet User
	e.ID = int(ebookID)
	utils.Logger()
	return err
}

// Select Ebooks
func SelectAllEbooks() ([]models.Ebook, error) {
	var ebooks []models.Ebook

	// Requête pour récupérer tous les livres électroniques depuis la base de données
	rows, err := DB.Query("SELECT id, title, author, category_id FROM ebooks")
	if err != nil {
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
	err := DB.QueryRow("SELECT id, title, author, category_id FROM ebooks WHERE id = ?", id).Scan(&ebook.ID, &ebook.Title, &ebook.Author, &ebook.CategoryID)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Ebook{}, fmt.Errorf("no ebook found with ID %d", id)
		}
		return models.Ebook{}, err
	}
	return ebook, nil
}

// Update Ebook
func UpdateEbook(id int, updatedEbook models.Ebook) (models.Ebook, error) {
	// Requête pour mettre à jour les informations du livre électronique dans la base de données
	_, err := DB.Exec("UPDATE ebooks SET title = ?, author = ?, category_id = ? WHERE id = ?", updatedEbook.Title, updatedEbook.Author, updatedEbook.CategoryID, id)
	if err != nil {
		return models.Ebook{}, err
	}

	// Mettre à jour l'ID dans l'objet Ebook
	updatedEbook.ID = id

	return updatedEbook, nil
}

// Delete Ebook
func DeleteEbook(id int) error {
	// Requête pour supprimer un livre électronique par ID dans la base de données
	_, err := DB.Exec("DELETE FROM ebooks WHERE id = ?", id)
	return err
}

// Select Ebooks with category_id
func SelectEbooksByCategoryID(categoryID int) ([]models.Ebook, error) {
	var ebooks []models.Ebook

	// Requête pour récupérer les livres électroniques par ID de catégorie depuis la base de données
	rows, err := DB.Query("SELECT id, title, author, category_id FROM ebooks WHERE category_id = ?", categoryID)
	if err != nil {
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
