package controllers

import (
	"example/api/dao"
	"example/api/models"
)

func InsertFavorite(f models.Favorite) error {
	err := dao.InsertFavorite(f)
	if err != nil {
		return err
	}
	return nil
}

func GetAllFavorites() ([]models.Favorite, error) {
	favorites, err := dao.SelectAllFavorites()
	if err != nil {
		return nil, err
	}
	return favorites, nil
}

func GetFavoriteByID(id int) (models.Favorite, error) {
	favorite, err := dao.SelectFavoriteByID(id)
	if err != nil {
		return models.Favorite{}, err
	}
	return favorite, nil
}

func GetFavoritesByUserID(userID int) ([]models.Favorite, error) {
	favorites, err := dao.SelectFavoriteByUserID(userID)
	if err != nil {
		return nil, err
	}
	return favorites, nil
}

func GetFavoritesByEbookID(ebookID int) ([]models.Favorite, error) {
	favorites, err := dao.SelectFavoriteByEbookID(ebookID)
	if err != nil {
		return nil, err
	}
	return favorites, nil
}

func UpdateFavorite(id int, updatedFavorite models.Favorite) error {
	existingFavorite, err := dao.SelectFavoriteByID(id)
	if err != nil {
		return err
	}
	existingFavorite.UserID = updatedFavorite.UserID
	existingFavorite.EbookID = updatedFavorite.EbookID
	_, err = dao.UpdateFavorite(id, existingFavorite)
	if err != nil {
		return err
	}
	return nil
}

func DeleteFavorite(id int) error {
	err := dao.DeleteFavorite(id)
	if err != nil {
		return err
	}
	return nil
}
