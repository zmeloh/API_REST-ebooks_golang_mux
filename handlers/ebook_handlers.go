package handlers

import (
	"example/api/models"
	"example/api/database"
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"database/sql"
	"fmt"
)

func CreateEbook(w http.ResponseWriter, r *http.Request){
	var ebook models.Ebook
    err := json.NewDecoder(r.Body).Decode(&ebook)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    _, err = database.DB.Exec("INSERT INTO ebooks (title, author, category_id) VALUES (?, ?, ?)", ebook.Title, ebook.Author, ebook.CategoryID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(ebook)
}


func GetEbookByID(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	ebookID := vars["id"]

	var ebook models.Ebook

	err := database.DB.QueryRow("SELECT id, title, author, category_id FROM ebooks WHERE id = ?", ebookID).Scan(&ebook.ID, &ebook.Title, &ebook.Author, &ebook.CategoryID)
	if err != nil {
		if err == sql.ErrNoRows{
			http.NotFound(w,r)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(ebook)
}

func GetEbooksByCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	categoryID := vars["categoryID"]

	rows, err := database.DB.Query("SELECT id, title, author, category_id FROM ebooks WHERE category_id = ?", categoryID)
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

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(ebooks)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}


func GetAllEbooks(w http.ResponseWriter, r *http.Request){
	rows, err := database.DB.Query("SELECT id, title, author, category_id FROM ebooks")
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

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(ebooks)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}


func UpdateEbook(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	ebookID := vars["id"]

	var updatedEbook models.Ebook
	err := json.NewDecoder(r.Body).Decode(&updatedEbook)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = database.DB.Exec("UPDATE ebooks SET title = ?, author = ?, category_id = ? WHERE id = ?", updatedEbook.Title, updatedEbook.Author, updatedEbook.CategoryID, ebookID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(updatedEbook)
}


func DeleteEbook(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	ebookID := vars["id"]
	_, err := database.DB.Exec("DELETE FROM ebooks WHERE id = ?", ebookID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Favorite with ID %s has been deleted", ebookID)
}