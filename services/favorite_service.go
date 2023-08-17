package services

import (
	"example/api/controllers"
	"example/api/models"
)

// GetAllFavorites
func GetAllFavorites() (f []models.Favorite) {
	f, _ = controllers.GetAllFavorites()
	return
}

// GetFavoriteByID
func GetFavoriteByID(id int) (f models.Favorite) {
	controllers.GetFavoriteByID(id)
	return
}

// InsertFavorite
func InsertFavorite(f *models.Favorite) {
	controllers.InsertFavorite(f)
	return
}

// UpdateFavorite
func UpdateFavorite(id int, f models.Favorite) {
	controllers.UpdateFavorite(id, f)
	return
}

// DeleteFavorite
func DeleteFavorite(id int) {
	controllers.DeleteFavorite(id)
	return
}
