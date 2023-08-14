package dao

import (
	"database/sql"
	"encoding/json"
	"example/api/models"
	"net/http"

	"github.com/gorilla/mux"
)

func insertEbook(w http.ResponseWriter, r *http.Request) {
	var ebook models.Ebook
	err := json.NewDecoder(r.Body).Decode(&ebook)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Insérer le livre électronique dans la base de données
	result, err := DB.Exec("INSERT INTO ebooks (title, author, category_id) VALUES (?, ?, ?)", ebook.Title, ebook.Author, ebook.CategoryID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Récupérer l'ID généré pour le livre électronique
	ebookID, err := result.LastInsertId()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Mettre à jour l'ID dans l'objet Ebook
	ebook.ID = int(ebookID)
}

func selectAllEbooks(w http.ResponseWriter, r *http.Request) {
	// Requête pour récupérer tous les livres électroniques depuis la base de données
	rows, err := DB.Query("SELECT id, title, author, category_id FROM ebooks")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var ebooks []models.Ebook
	for rows.Next() {
		var ebook models.Ebook
		err := rows.Scan(&ebook.ID, &ebook.Title, &ebook.Author, &ebook.CategoryID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		ebooks = append(ebooks, ebook)
	}

}

func SelectEbookByID(id int) models.Ebook {

	var ebook models.Ebook

	// Requête pour récupérer le livre électronique par ID depuis la base de données
	err := DB.QueryRow("SELECT id, title, author, category_id FROM ebooks WHERE id = ?", ebookID).Scan(&ebook.ID, &ebook.Title, &ebook.Author, &ebook.CategoryID)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Ebook{}
		}

	}
	return ebook
}

func updateEbook(e models.Ebook) {

	var updatedEbook models.Ebook
	err := json.NewDecoder(r.Body).Decode(&updatedEbook)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Requête pour mettre à jour les informations du livre électronique dans la base de données
	_, err = DB.Exec("UPDATE ebooks SET title = ?, author = ?, category_id = ? WHERE id = ?", updatedEbook.Title, updatedEbook.Author, updatedEbook.CategoryID, ebookID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Mettre à jour l'ID dans l'objet Ebook
	//updatedEbook.ID = ebookIDInt
}

func deleteEbook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ebookID := vars["id"]

	// Requête pour supprimer le livre électronique depuis la base de données
	_, err := DB.Exec("DELETE FROM ebooks WHERE id = ?", ebookID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func selectEbookByCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	categoryID := vars["categoryID"]

	// Requête pour récupérer tous les livres électroniques d'une catégorie depuis la base de données
	rows, err := DB.Query("SELECT id, title, author, category_id FROM ebooks WHERE category_id = ?", categoryID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var ebooks []models.Ebook
	for rows.Next() {
		var ebook models.Ebook
		err := rows.Scan(&ebook.ID, &ebook.Title, &ebook.Author, &ebook.CategoryID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		ebooks = append(ebooks, ebook)
	}
}
