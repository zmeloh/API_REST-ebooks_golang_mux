package rest

import (
	"encoding/json"
	"strconv"

	"example/api/models"
	"example/api/services"
	"example/api/utils"
	"net/http"

	"github.com/gorilla/mux"
)

// CreateFavorite crée un nouveau favori.
func CreateFavorite(w http.ResponseWriter, r *http.Request) {
	// Décode les données JSON du corps de la requête dans une struct Favorite
	var favorite models.Favorite
	err := json.NewDecoder(r.Body).Decode(&favorite)
	if err != nil {
		utils.Logger(err)
		ServerResponse(w, http.StatusBadRequest, "Invalid request")
		return
	}
	services.InsertFavorite(&favorite)
	// Répond avec le favori créé et le code de statut 201 (Created)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(favorite)
}

// GetAllFavorites récupère tous les favoris.
func GetAllFavorites(w http.ResponseWriter, r *http.Request) {
	favorites := services.GetAllFavorites()
	if favorites == nil {
		ServerResponse(w, http.StatusNotFound, "No data found")
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(favorites)
}

// GetFavoriteByID récupère un favori par son ID.
func GetFavoriteByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	result, err := strconv.Atoi(id)
	if err != nil {
		utils.Logger(err)
		ServerResponse(w, http.StatusBadRequest, "Invalid request")
		return
	}


	favorite := services.GetFavoriteByID(result)
	if favorite.ID == 0 {
		ServerResponse(w, http.StatusNotFound, "Favorite not found")
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(favorite)
}

// GetFavoritesByUserID récupère les favoris par ID d'utilisateur.
func GetFavoritesByUserID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	favoriteID, err := strconv.Atoi(id)

	if err != nil {
		utils.Logger(err)
		ServerResponse(w, http.StatusBadRequest, "Invalid request")
		return
	}

	// Appeler le service pour récupérer les livres électroniques de la catégorie
	favorites := services.GetFavoriteByUserID(favoriteID)
	if len(favorites) == 0 {
		ServerResponse(w, http.StatusNotFound, "Favorite ID not found")
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(favorites)
}

// GetFavoritesByEbookID récupère les favoris par ID d'ebook.
func GetFavoritesByEbookID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	favoriteID, err := strconv.Atoi(id)

	if err != nil {
		utils.Logger(err)
		ServerResponse(w, http.StatusBadRequest, "Invalid request")
		return
	}

	// Appeler le service pour récupérer les livres électroniques de la catégorie
	favorites := services.GetFavoriteByEbookID(favoriteID)
	if len(favorites) == 0 {
		ServerResponse(w, http.StatusNotFound, "Favorite ID not found")
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(favorites)
}

// UpdateFavorite met à jour un favori par son ID.
func UpdateFavorite(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	favoriteID, err := strconv.Atoi(id)

	if err != nil {
		utils.Logger(err)
		ServerResponse(w, http.StatusBadRequest, "Invalid request")
		return
	}

	var updateFavorite models.Favorite
	err = json.NewDecoder(r.Body).Decode(&updateFavorite)
	if err != nil {
		ServerResponse(w, http.StatusBadRequest, "Invalid request")
		return
	}

	err = services.UpdateFavorite(favoriteID, &updateFavorite)
	if err != nil {
		ServerResponse(w, http.StatusNotFound, "Favorite not found")
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updateFavorite)
}

// DeleteFavorite supprime un favori par son ID.
func DeleteFavorite(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	favoriteID, err := strconv.Atoi(id)

	if err != nil {
		utils.Logger(err)
		ServerResponse(w, http.StatusBadRequest, "Invalid id")
		return
	}

	// Appeler le service pour supprimer l'ebook
	err = services.DeleteFavorite(favoriteID)
	if err != nil {
		ServerResponse(w, http.StatusNotFound, "Favorite not found")
		return
	}
	ServerResponse(w, http.StatusOK, "Favorite has been deleted")
}
