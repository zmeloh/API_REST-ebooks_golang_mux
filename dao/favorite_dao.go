package dao

import (
	"database/sql"
	"example/api/models"
	"example/api/utils"
	"fmt"
)

func InsertFavorite(f models.Favorite) error {
	var favorite models.Favorite
	// Insère le favori dans la base de données
	result, err := DB.Exec("INSERT INTO favorites (user_id, ebook_id) VALUES (?, ?)", favorite.UserID, favorite.EbookID)
	if err != nil {
		return err
	}
	// Obtient l'ID généré lors de l'insertion
	favoriteID, err := result.LastInsertId()
	if err != nil {
		return err
	}

	// Met à jour l'ID dans l'objet User
	f.ID = int(favoriteID)
	utils.Logger()
	return err
}

func SelectAllFavorites() ([]models.Favorite, error) {
	var favorites []models.Favorite
	// Récupère tous les favoris depuis la base de données
	rows, err := DB.Query("SELECT id, user_id, ebook_id FROM favorites")
	if err != nil {
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
	err := DB.QueryRow("SELECT id, user_id, ebook_id FROM favorites WHERE id = ?", id).Scan(&favorite.ID, &favorite.UserID, &favorite.EbookID)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Favorite{}, fmt.Errorf("no ebook found with ID %d", id)
		}
		return models.Favorite{}, err
	}
	return favorite, nil
}

func UpdateFavorite(id int, updatedFavorite models.Favorite) (models.Favorite, error) {

	// Met à jour les données du favori dans la base de données
	_, err := DB.Exec("UPDATE favorites SET user_id = ?, ebook_id = ? WHERE id = ?", updatedFavorite.UserID, updatedFavorite.EbookID, id)
	if err != nil {
		return models.Favorite{}, err
	}

	// Mettre à jour l'ID dans l'objet Ebook
	updatedFavorite.ID = id

	return updatedFavorite, nil

}

func DeleteFavorite(id int) error {
	_, err := DB.Exec("DELETE FROM favorites WHERE id = ?", id)
	return err

}

func SelectFavoriteByUserID(userID int) ([]models.Favorite, error) {
	var favorites []models.Favorite
	// Interroge la base de données pour obtenir les favoris correspondant à l'ID d'utilisateur
	rows, err := DB.Query("SELECT id, user_id, ebook_id FROM favorites WHERE user_id = ?", userID)
	if err != nil {
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
	rows, err := DB.Query("SELECT id, user_id, ebook_id FROM favorites WHERE ebook_id = ?", ebookID)
	if err != nil {
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
