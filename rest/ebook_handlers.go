package rest

import (
	"encoding/json"
	"example/api/models"
	"example/api/services"
	"example/api/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// CreateEbook crée un nouveau livre électronique.
func CreateEbook(w http.ResponseWriter, r *http.Request) {
	var newEbook models.Ebook
	err := json.NewDecoder(r.Body).Decode(&newEbook)
	if err != nil {
		utils.Logger(err)
		ServerResponse(w, http.StatusBadRequest, "Invalid request")
		return
	}

	category := services.GetCategoryByID(newEbook.CategoryID)

	if category.ID == 0 {
		ServerResponse(w, http.StatusNotFound, "Category id not exist")
		return
	}
	// Appeler le service pour insérer le nouvel ebook
	services.InsertEbook(&newEbook)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newEbook)
}

// GetAllEbook récupère la liste de tous les livres électroniques.
func GetAllEbooks(w http.ResponseWriter, r *http.Request) {
	ebooks := services.GetAllEbooks()
	if ebooks == nil {
		ServerResponse(w, http.StatusNotFound, "No data found")
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ebooks)
}

// GetEbookByID récupère un livre électronique par son ID.
func GetEbookByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	result, err := strconv.Atoi(id)

	if err != nil {
		utils.Logger(err)
		ServerResponse(w, http.StatusBadRequest, "Invalid request")
		return
	}

	// Appeler le service pour récupérer un livre électronique par son ID
	ebook := services.GetEbookByID(result)
	if ebook.ID == 0 {
		ServerResponse(w, http.StatusNotFound, "Ebook not found")
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ebook)
}

// GetEbookByCategoryID récupère la liste de tous les livres électroniques d'une catégorie.
func GetEbookByCategoryID(w http.ResponseWriter, r *http.Request) {
	// Récupérer l'ID de la catégorie depuis les paramètres de la requête
	params := mux.Vars(r)
	id := params["id"]
	categoryID, err := strconv.Atoi(id)

	if err != nil {
		utils.Logger(err)
		ServerResponse(w, http.StatusBadRequest, "Invalid request")
		return
	}

	// Appeler le service pour récupérer les livres électroniques de la catégorie
	ebooks := services.GetEbookByCategoryID(categoryID)
	if len(ebooks) == 0 {
		ServerResponse(w, http.StatusNotFound, "Category ID not found")
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ebooks)
}

// UpdateEbook met à jour un livre électronique par son ID.
func UpdateEbook(w http.ResponseWriter, r *http.Request) {
	// Récupérer l'ID de l'ebook depuis les variables de la requête
	params := mux.Vars(r)
	id := params["id"]
	ebookID, err := strconv.Atoi(id)

	if err != nil {
		utils.Logger(err)
		ServerResponse(w, http.StatusBadRequest, "Invalid request")
		return
	}

	var updatedEbook models.Ebook
	err = json.NewDecoder(r.Body).Decode(&updatedEbook)
	if err != nil {
		ServerResponse(w, http.StatusBadRequest, "Invalid request")
		return
	}

	category := services.GetCategoryByID(updatedEbook.CategoryID)
	if category.ID == 0 {
		ServerResponse(w, http.StatusNotFound, "Category id not exist")
		return
	}

	err = services.UpdateEbook(ebookID, &updatedEbook)
	if err != nil {
		ServerResponse(w, http.StatusNotFound, "Ebook not found")
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedEbook)
}

// DeleteEbook supprime un livre électronique par son ID.
func DeleteEbook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	ebookID, err := strconv.Atoi(id)

	if err != nil {
		utils.Logger(err)
		ServerResponse(w, http.StatusBadRequest, "Invalid id")
		return
	}

	// Appeler le service pour supprimer l'ebook
	err = services.DeleteEbook(ebookID)
	if err != nil {
		ServerResponse(w, http.StatusNotFound, "Ebook not found")
		return
	}
	ServerResponse(w, http.StatusOK, "Ebook has been deleted")
}
