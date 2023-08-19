package dao

import (
	"database/sql"
	"errors"
	"example/api/models"
	"example/api/utils"
	"fmt"
)

func InsertFavorite(f *models.Favorite) error {
	var favorite models.Favorite
	// Insère le favori dans la base de données
	err := DB.QueryRow("INSERT INTO favorites (user_id, ebook_id) VALUES ($!, $2) RETURNING id", favorite.UserID, favorite.EbookID).Scan(f.ID)
	if err != nil {
		utils.Logger(err)
		return err
	}

	return err
}

func SelectAllFavorites() ([]models.Favorite, error) {
	var favorites []models.Favorite
	// Récupère tous les favoris depuis la base de données
	rows, err := DB.Query("SELECT id, user_id, ebook_id FROM favorites ORDER BY id")
	if err != nil {
		utils.Logger(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var favorite models.Favorite
		err := rows.Scan(&favorite.ID, &favorite.UserID, &favorite.EbookID)
		if err != nil {
			return nil, err
		}
		favorites = append(favorites, favorite)
	}

	return favorites, nil
}

func SelectFavoriteByID(id int) (models.Favorite, error) {

	var favorite models.Favorite

	// Interroge la base de données pour obtenir le favori correspondant à l'ID
	err := DB.QueryRow("SELECT id, user_id, ebook_id FROM favorites WHERE id = $1", id).Scan(&favorite.UserID, &favorite.EbookID)
	if errors.Is(err, sql.ErrNoRows) {
		return models.Favorite{}, fmt.Errorf("no favorite found with ID %d", id)
	}
	if err != nil {
		utils.Logger(err)
		return models.Favorite{}, err
	}
	return favorite, nil
}

func UpdateFavorite(id int, updatedFavorite models.Favorite) (models.Favorite, error) {

	// Met à jour les données du favori dans la base de données
	_, err := DB.Exec("UPDATE favorites SET user_id = $1, ebook_id = $2 WHERE id = $3", updatedFavorite.UserID, updatedFavorite.EbookID, updatedFavorite.ID)
	if err != nil {
		utils.Logger(err)
		return models.Favorite{}, err
	}

	return updatedFavorite, nil

}

func DeleteFavorite(id int) error {
	err := DB.QueryRow("DELETE FROM favorites WHERE id = $1", id).Scan(&id)
	if err != nil {
		utils.Logger(err)
	}
	return err

}

func SelectFavoriteByUserID(userID int) ([]models.Favorite, error) {
	var favorites []models.Favorite
	// Interroge la base de données pour obtenir les favoris correspondant à l'ID d'utilisateur
	rows, err := DB.Query("SELECT id, user_id, ebook_id FROM favorites WHERE user_id = $1", userID)
	if err != nil {
		utils.Logger(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var favorite models.Favorite
		err := rows.Scan(&favorite.ID, &favorite.UserID, &favorite.EbookID)
		if err != nil {
			return nil, err
		}
		favorites = append(favorites, favorite)
	}

	return favorites, nil
}

func SelectFavoriteByEbookID(ebookID int) ([]models.Favorite, error) {
	var favorites []models.Favorite
	// Interroge la base de données pour obtenir les favoris correspondant à l'ID d'ebook
	rows, err := DB.Query("SELECT id, user_id, ebook_id FROM favorites WHERE ebook_id = $1", ebookID)
	if err != nil {
		utils.Logger(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var favorite models.Favorite
		err := rows.Scan(&favorite.ID, &favorite.UserID, &favorite.EbookID)
		if err != nil {
			utils.Logger(err)

			return nil, err
		}
		favorites = append(favorites, favorite)
	}

	return favorites, nil
}
