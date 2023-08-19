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
	f, _ = controllers.GetFavoriteByID(id)
	return
}

// InsertFavorite
func InsertFavorite(f *models.Favorite) {
	controllers.InsertFavorite(f)
	return
}

// UpdateFavorite
func UpdateFavorite(id int, f *models.Favorite) error {
	err := controllers.UpdateFavorite(id, f)
	return err
}

// DeleteFavorite
func DeleteFavorite(id int) error {
	err := controllers.DeleteFavorite(id)
	return err
}
